package main

import (
	"bufio"
	"context"
	covidclient "covid_client/covidclient"
	"covid_client/grpcconn"
	"covid_client/http"
	pb "covid_client/pb/proto"
	"covid_client/util"
	"time"

	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type newData struct {
	country string
}

// newDataValue is a NewDataValue.
func newDataValue() *newData {
	return &newData{}
}

var dataValue = newDataValue()

var searchText string
var shouldFollowSearch string

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	port := flag.Int("port", 8080, "Http port to run the client server")
	flag.Parse()

	// Creates http server.
	httpServer := http.New(8081)

	_, err := httpServer.CreateHTTPServer()
	if err != nil {
		panic(err)
	}
	defer httpServer.Stop()

	// Connect to the Grpc server on given port.
	grpcConnection, err := grpcconn.NewGrpcClient(*port)
	if err != nil {
		util.Error("Unable to establish grpc server connection ", err)
	}

	covidServiceClient := pb.NewCovidServiceClient(grpcConnection)
	// New grpc client.
	covidClient := covidclient.New(covidServiceClient)

	var count int
	chang := make(chan string)
	newChannel := make(chan string)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the country you want to know info: ")
	for scanner.Scan() {
		searchText = scanner.Text()

		fmt.Print("Do you want to follow the update? ")
		go func() {
			shouldFollowSearch = scanner.Text()

			chang <- shouldFollowSearch
		}()

		readValue := <-chang

		if readValue == "y" {
			duration := 2 * time.Second
			go func(text string) {
				wait := duration
				for {
					time.Sleep(wait)
					_, err := covidClient.GetData(ctx, text)
					count = count + 1
					if err != nil {
						panic(err)
					}

					newChannel <- fmt.Sprintf("%v", count)
				}

			}("Nepal")

			dataValue.country = <-newChannel
			data := [][]string{
				[]string{
					dataValue.country,
					// fmt.Sprintf("%d", res.TotalCases),
					// fmt.Sprintf("%d", res.TotalDeaths),
					// fmt.Sprintf("%d", res.TotalRecovered),
					// fmt.Sprintf("%v", res.DeathsPerOneMillion),
				},
			}
			// "TotalCases", "Total Deaths", "Total Recovered", "Death per million"}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Name"})

			for _, v := range data {
				table.Append(v)
			}
			table.Render()
		}

		if readValue == "n" {
			fmt.Println(shouldFollowSearch)

		}

		if readValue != "n" || readValue != "y" {
			fmt.Println("You need to put y or n.")

		}

	}

	// for scanner.Scan() {
	// 	searchText = scanner.Text()

	// 	matched, err := regexp.MatchString("total.*", searchText)
	// 	if err != nil {
	// 		fmt.Println("total pattern not matched")
	// 	}

	// 	if matched {
	// 		covidClient.GetData(ctx, searchText)
	// 	} else {
	// 		covidClient.GetData(ctx, searchText)
	// 	}

	// 	fmt.Print("Enter the country you want to know info: ")
	// }

}
