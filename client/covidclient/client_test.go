package covidclient_test

import (
	"context"
	"covid_client/covidclient"
	pb "covid_client/pb/proto"
	"fmt"
	"testing"

	"google.golang.org/grpc"
)

type mockConnection struct{}

func (c *mockConnection) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
	return nil
}
func (c *mockConnection) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, nil
}

type mockcovidServiceClient struct{}

var newMockcovidServiceClient = pb.NewCovidServiceClient(&mockConnection{})

func (mcsc *mockcovidServiceClient) GetCurrentCovidInfo(ctx context.Context, in *pb.CovidCasesRequest, opts ...grpc.CallOption) (*pb.CovidCasesResponse, error) {
	fmt.Println("hello world")
	return nil, nil
}

var mockCovidClient = covidclient.New(newMockcovidServiceClient)

func TestGetData(t *testing.T) {
	testCases := []struct {
		name string
	}{
		{
			name: "kushal",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := mockCovidClient.GetData(context.Background(), "helllo")
			if err != nil {
				panic(err)
			}
		})
	}
}

// func (c *mockGrpcConn) Invoke(ctx context.Context, method string, args interface{}, reply interface{}, opts ...grpc.CallOption) error {
// return nil
// }
// func (c *mockGrpcConn) NewStream(ctx context.Context, desc *grpc.StreamDesc, method string, opts ...grpc.CallOption) (grpc.ClientStream, error) {
// return nil, nil
// }
//
// func (c *mockGrpcConn) GetCurrentCovidInfo(ctx context.Context, in *pb.CovidCasesRequest, opts ...grpc.CallOption) (*pb.CovidCasesResponse, error) {
// return nil, nil
// }
//
