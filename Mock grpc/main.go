package Mock_grpc

import (
	"context"
	"google.golang.org/grpc"
)

func (c *controllers) SayHelloWorld() (response string) {
	res, err := GreeterClient.SayHello(nil, nil)
	if res.word == "Hi" {
		response = "Hi there! how are you doing?"
	} else if res.word == "Bye" {
		response = "Bye, see you again soon."
	}
	return response
}

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error)
}
type greeterClient struct {
	cc *grpc.ClientConn
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	// ...
	// gRPC specific code here
	// ...
}

type MockGreeterClient struct {
	pb.MockGreeterClient
}

func (c *MockGreeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	// return fake response
}

type grpcServices struct {
	HelloServices HelloServices
}
type HelloServices struct {
	GreeterClient pb.GreeterClient
}

var GRPCServices = &grpcServices{}

func init() {
	GRPCServices.HelloServices.GreeterClient = pb.NewGreeterClient(&conn)
}
