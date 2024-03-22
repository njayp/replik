package conn

import (
	"fmt"
	"net"

	"github.com/njayp/replik/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewListener() net.Listener {
	port := 9090
	address := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	return lis
}

func NewClient() api.ReplikClient {
	port := 9090
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), opts...)
	if err != nil {
		panic(err)
	}
	return api.NewReplikClient(conn)
}
