package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"water-tank-api/controllers"
	"water-tank-api/core/entity/logs"
	database_mock "water-tank-api/infra/database/mock"
	"water-tank-api/infra/stdout"
	"water-tank-api/infra/web"
	"water-tank-api/infra/web/routes"

	kingpin "github.com/alecthomas/kingpin/v2"
	iris "github.com/kataras/iris/v12"
	"golang.org/x/sync/errgroup"
)

var (
	port = kingpin.Flag("port", "Server's port").Short('p').Default("8080").Envar("SERVER_PORT").Int()
)

func main() {
	kingpin.Parse()

	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	logs.SetLogger(stdout.NewSTDOutLogger())

	app := iris.New()
	internalRouter := routes.InternalRouter{}

	web.SetControllers(controllers.NewController(database_mock.NewWaterTankMockData()))
	internalRouter.Route(app)

	go func() {
		if err := app.Run(iris.Addr(fmt.Sprintf(":%d", *port))); err != nil {
			logs.Gateway().Fatal(fmt.Sprintf("Error on starting http listener: %s", err.Error()))
		}
	}()

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		<-gCtx.Done()

		app.Shutdown(context.Background())

		return nil
	})

	if err := g.Wait(); err != nil {
		logs.Gateway().Fatal(fmt.Sprintf("exit reason: %s \n", err))
	}
}
