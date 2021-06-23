package svc

import (
	"os"

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
	}
	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(auth),
		grpc.WithTransportCredentials(creds),
	}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, &grpc.ClientConn{}, err
	}

	return gen.NewAxoneClient(conn), conn, nil

}

func Dial(login, password string) (gen.AxoneClient, *grpc.ClientConn, error) {
	return grpcClient(BasicAuth{
		Login:    login,
		Password: password,
	})
}

func DialWithEnvVariables() (gen.AxoneClient, *grpc.ClientConn, error) {
	return Dial(os.Getenv("AXONE_LOGIN"), os.Getenv("AXONE_PASSWORD"))
}
