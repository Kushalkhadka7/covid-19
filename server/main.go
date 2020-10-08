package main

import (
	grpcservice "covid_server/grpc"
	"covid_server/http"
	"log"
)

func main() {
	httpServer := http.New(8080)
	listener, err := httpServer.CreateHTTPServer()
	if err != nil {
		log.Fatal(err)
	}

	grpcserver := grpcservice.NewGrpcServer(listener)
	if err := grpcserver.CreateCovidServer(); err != nil {
		log.Fatal(err)
	}
}
