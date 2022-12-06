package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"log"
	"mitte-task/routes"
	"mitte-task/services"
	"net/http"
	"os"
	"time"
)

var ginLambda *ginadapter.GinLambda

func init() {
	services.LoadConfig()
	services.InitMongoDB()

	routes.InitGin()
	router := routes.New()

	if services.Config.Mode == "release" {
		ginLambda = ginadapter.New(router)
	} else {
		server := &http.Server{
			Addr:         services.Config.ServerAddr + ":" + services.Config.ServerPort,
			WriteTimeout: time.Second * 30,
			ReadTimeout:  time.Second * 30,
			IdleTimeout:  time.Second * 30,
			Handler:      router,
		}

		go func() {
			if err := server.ListenAndServe(); err != nil {
				log.Printf("listen: %s\n", err)
			}
		}()

		// Wait for interrupt signal to gracefully shut down the server with
		// a timeout of 15 seconds.
		quit := make(chan os.Signal, 1)
		//signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}

		log.Println("Server exiting")
	}

}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

func main() {
	lambda.Start(Handler)
}
