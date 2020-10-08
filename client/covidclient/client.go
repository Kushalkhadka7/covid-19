package covidclient

import (
	"context"
	"covid_client/grpcconn"
	pb "covid_client/pb/proto"
	"fmt"
	"log"
)

var responseData *pb.Response

// CovidClient intializes grpc connection.
type CovidClient struct {
	pb.CovidServiceClient
}

// Servicer grpc service interface.
type Servicer interface {
	GetData(ctx context.Context, text string) error
	GetTotalData(text string) error
}

// New create new covid-19 client.
func New(conn *grpcconn.GrpcClientConn) Servicer {
	return &CovidClient{
		pb.NewCovidServiceClient(conn),
	}
}

// GetData get covid related data.
func (cc *CovidClient) GetData(ctx context.Context, text string) error {
	fmt.Println("hello world")
	req := &pb.CovidCasesRequest{
		SearchString: text,
	}

	res, err := cc.GetCurrentCovidInfo(ctx, req)
	if err != nil {
		panic(err)
	}

	responseData = res

	fmt.Println("hello world")
	log.Print(responseData)

	return nil
}

// GetTotalData total data.
func (cc *CovidClient) GetTotalData(text string) error {
	log.Println("text")
	log.Println(text)
	return nil

}
