package svc

import (
	"github.com/kouame-florent/axone-cx/api/grpc/gen"
	"github.com/kouame-florent/axone-cx/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func grpcClient(auth BasicAuth) (gen.AxoneClient, *grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(config.ServerCertFile, "")
	if err != nil {
		return nil, &grpc.ClientConn{}, err
		//log.Fatalf("failed to load credentials: %v", err)
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, &grpc.ClientConn{}, err
		//log.Fatalf("did not connect: %v", err)
	}

	return gen.NewAxoneClient(conn), conn, nil

}

func Dial(username, password string) (gen.AxoneClient, *grpc.ClientConn, error) {
	return grpcClient(BasicAuth{
		Username: username,
		Password: password,
	})
}
