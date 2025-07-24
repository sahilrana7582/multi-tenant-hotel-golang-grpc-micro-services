package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sahilrana7582/multi-tenant-hotel/grpc/user/user"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/models"
	"github.com/sahilrana7582/multi-tenant-hotel/user-service/service"
	ggrpc "google.golang.org/grpc"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	service service.UserService
}

func NewUserGrpcServer(service service.UserService) *UserServer {
	return &UserServer{
		service: service,
	}
}

func StartGRPCServer(userService service.UserService, port string) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := ggrpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewUserGrpcServer(userService))

	fmt.Printf("🚀 gRPC server started on :%s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *UserServer) GenerateNewUser(ctx context.Context, grpcNewUser *pb.CreateNewUser) (*pb.CreateNewUserResponse, error) {
	var newUser models.NewUser

	newUser.Name = grpcNewUser.Name
	newUser.Email = grpcNewUser.Email
	newUser.Password = "1234"
	newUser.TenantID = grpcNewUser.TenantId

	user, err := s.service.CreateUser(ctx, &newUser)

	if err != nil {
		return nil, fmt.Errorf("error in grpc call: %w", err)
	}

	return &pb.CreateNewUserResponse{
		Id: user.ID,
	}, nil
}
