package grpcconn

import (
	"covid_client/util"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

//GrpcClientConn connection.
type GrpcClientConn struct {
	*grpc.ClientConn
}

// NewGrpcClient grpc client to connect to the server in given port number.
func NewGrpcClient(port int) (*GrpcClientConn, error) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		return nil, util.Error("Cannot connect to grpc server: ", err)

	}

	log.Printf("Grpc Server listening on :%d\n", port)
	return &GrpcClientConn{conn}, nil
}
