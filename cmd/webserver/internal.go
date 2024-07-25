package webserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"water-tank-api/app/entity/logs"
	"water-tank-api/app/usecases/create_tank"
	"water-tank-api/app/usecases/get_group"
	"water-tank-api/app/usecases/get_tank"
	"water-tank-api/app/usecases/update_tank_state"
	mongodb "water-tank-api/infra/database/mongoDB"
	"water-tank-api/infra/logs/stdout"
	web "water-tank-api/infra/web/http"
)

func Internal() {
	mainCtx := context.Background()

	logs.SetLogger(stdout.NewSTDOutLogger())

	mongoClient, err := mongodb.InitClient(mainCtx, databaseURI)
	if err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("Error on starting mongo DB client: %s", err.Error()))
	}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    serverPort,
		Handler: mux,
	}

	collection := mongodb.NewCollection(mainCtx, mongoClient, databaseName, databaseCollection)

	getTankUsecase := get_tank.NewGetWaterTank(collection)

	internalAPI := web.NewInternalAPI(
		getTankUsecase,
		get_group.NewGetGroupWaterTank(collection),
		create_tank.NewWaterTank(collection, getTankUsecase),
		update_tank_state.NewWaterTankUpdate(collection, getTankUsecase),
	)

	internalAPI.Route(mux)

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
		return
	}
	logs.Gateway().Info("Graceful shutdown complete")
}
