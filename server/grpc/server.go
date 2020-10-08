package grpcservice

import (
	"covid_server/covidservice"
	pb "covid_server/pb/proto"
	"covid_server/util"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GrpcServer initializes new covid grpc server.
type GrpcServer struct {
	listener net.Listener
}

// NewGrpcServer initilizes listner with grpc server.
func NewGrpcServer(listener net.Listener) *GrpcServer {
	return &GrpcServer{listener}
}

// CreateCovidServer creates new covid service server
func (cs *GrpcServer) CreateCovidServer() error {
	grpcServer := grpc.NewServer()
	covidServiceServer := covidservice.NewCovidServiceServer()

	pb.RegisterCovidServiceServer(grpcServer, covidServiceServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(cs.listener); err != nil {
		return util.Error("Unable to establish grpc server: ", err)
	}

	return nil
}
