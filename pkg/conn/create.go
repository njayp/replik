package conn

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = 9090

// TODO replace pkg with env
func NewListener() net.Listener {
	address := fmt.Sprintf(":%v", port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	return lis
}

func NewConn() *grpc.ClientConn {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(fmt.Sprintf(":%v", port), opts...)
	if err != nil {
		panic(err)
	}
	return conn
}
