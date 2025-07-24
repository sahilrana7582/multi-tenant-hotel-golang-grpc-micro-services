package grpcclient

import (
	"context"
	"fmt"
	"time"

	pb "github.com/sahilrana7582/multi-tenant-hotel/grpc/user/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGrpcClient struct {
	userClient pb.UserServiceClient
}

func NewUserGrpcClient(addr string) (*UserGrpcClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("could not connect to user-service: %w", err)
	}

	client := pb.NewUserServiceClient(conn)

	return &UserGrpcClient{userClient: client}, nil
}

func (u *UserGrpcClient) GenerateNewUser(ctx context.Context, email, name, tenantID string) (*pb.CreateNewUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	newUser := pb.CreateNewUser{
		TenantId: tenantID,
		Name:     name,
		Email:    email,
	}

	return u.userClient.GenerateNewUser(ctx, &newUser)
}
