package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"water-tank-api/app/controllers"
	"water-tank-api/app/core/entity/logs"
	mongodb "water-tank-api/app/infra/database/mongoDB"
	"water-tank-api/app/infra/logs/stdout"
	"water-tank-api/app/infra/web/routes"

	iris "github.com/kataras/iris/v12"
	"golang.org/x/sync/errgroup"
)

var (
	port               = os.Getenv("SERVER_PORT")
	databaseURI        = os.Getenv("DATABASE_URI")
	databaseName       = os.Getenv("DATABASE_NAME")
	databaseCollection = os.Getenv("DATABASE_COLLECTION")
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logs.SetLogger(stdout.NewSTDOutLogger())

	mongoClient, err := mongodb.InitClient(mainCtx, databaseURI)
	if err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("Error on starting mongo DB client: %s", err.Error()))
	}
	collection := mongodb.NewCollection(mainCtx, mongoClient, databaseName, databaseCollection)

	app := iris.New()
	externalRouter := routes.ExternalRouter{}

	externalRouter.Route(app, controllers.NewController(collection))

	go func() {
		if err := app.Run(iris.Addr(fmt.Sprintf(":%s", port))); err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error on starting http listener: %s", err.Error()))
		}
	}()

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		<-gCtx.Done()

		app.Shutdown(context.Background())

		if err := mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("exit reason: %s \n", err))
	}
}
