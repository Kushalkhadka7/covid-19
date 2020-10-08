package main

import (
	"bufio"
	"context"
	covidclient "covid_client/covidclient"
	"covid_client/grpcconn"
	"covid_client/http"
	"covid_client/util"
	"time"

	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	port := flag.Int("port", 8080, "Http port to run the client server")
	flag.Parse()

	// Creates http server.
	httpServer := http.New(8081)

	// Connect to the Grpc server on given port.
	grpcConnection, err := grpcconn.NewGrpcClient(*port)
	if err != nil {
		util.Error("Unable to establish grpc server connection ", err)
	}

	// New grpc client.
	covidClient := covidclient.New(grpcConnection)

	duration := 2 * time.Second
	go func(text string) {
		wait := duration
		for {
			time.Sleep(wait)
			covidClient.GetData(ctx, text)
		}

	}("Nepal")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()

		matched, err := regexp.MatchString("total.*", text)
		if err != nil {
			fmt.Println("total pattern not matched")
		}

		if matched {
			covidClient.GetData(ctx, text)
		} else {
			covidClient.GetData(ctx, text)
		}

	}

	_, err = httpServer.CreateHTTPServer()
	if err != nil {
		panic(err)
	}
	defer httpServer.Stop()
}
