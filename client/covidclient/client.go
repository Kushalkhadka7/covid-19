package covidclient

import (
	"context"
	pb "covid_client/pb/proto"
	"fmt"
	"log"
)

var responseData *pb.CovidCasesResponse

// CovidClient intializes grpc connection.
type CovidClient struct {
	client pb.CovidServiceClient
}

// Servicer grpc service interface.
type Servicer interface {
	GetData(ctx context.Context, text string) (*pb.CovidCasesResponse, error)
	GetTotalData(text string) error
}

// New create new covid-19 client.
func New(client pb.CovidServiceClient) Servicer {
	return &CovidClient{client}
}

// GetData get covid related data.
func (cc *CovidClient) GetData(ctx context.Context, text string) (*pb.CovidCasesResponse, error) {
	fmt.Println("hello world")
	req := &pb.CovidCasesRequest{
		SearchString: text,
	}

	res, err := cc.client.GetCurrentCovidInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetTotalData total data.
func (cc *CovidClient) GetTotalData(text string) error {
	log.Println("text")
	log.Println(text)
	return nil

}
