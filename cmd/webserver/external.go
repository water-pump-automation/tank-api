package webserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"tank-api/app/entity/logs"
	"tank-api/app/usecases/get_group"
	"tank-api/app/usecases/get_tank"
	web "tank-api/cmd/webserver/routing"
	mongodb "tank-api/infra/database/mongoDB"
	"tank-api/infra/logs/kafka"

	"github.com/IBM/sarama"
)

func External() {
	mainCtx := context.Background()

	version, err := sarama.ParseKafkaVersion("0.0.1")
	if err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("Error parsing kafka version: %s", err.Error()))
	}

	logs.SetLogger(kafka.NewKafkaLogger(version, kafkaTopic, kafkaBrokers...))

	mongoClient, err := mongodb.InitClient(mainCtx, databaseURI)
	if err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("Error on starting mongo DB client: %s", err.Error()))
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}
	collection := mongodb.NewCollection(mainCtx, mongoClient, databaseName, tankCollection, stateCollection)

	externalAPI := web.NewExternalAPI(
		get_tank.NewGetTank(collection),
		get_group.NewGetGroupTank(collection),
	)

	externalAPI.Route(mux)

	go func() {
		logs.Gateway().Info("Started internal server on port:" + serverPort)
		if err := http.ListenAndServe(":"+serverPort, mux); err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error on starting http listener: %s", err.Error()))
		}
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctx, shutdownRelease := context.WithTimeout(mainCtx, 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(ctx); err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("Shutdown error: %s", err.Error()))
	}
	logs.Gateway().Info("Graceful shutdown complete")
}
