package grpcconn

import (
	"covid_client/util"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

// NewGrpcClient grpc client to connect to the server in given port number.
func NewGrpcClient(port int) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		return nil, util.Error("Cannot connect to grpc server: ", err)

	}

	log.Printf("Grpc Server listening on :%d\n", port)
	return conn, nil
}
