package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/VinceDeslo/grpc-test/server/services"
	"github.com/VinceDeslo/grpc-test/server/user/pb"
)

const (
	protocol = "tcp"
	port     = ":9090"
)

// type service struct {
// 	pb.UnimplementedUsersServer
// 	userDb user.UserDb
// }

// func NewUserService() *service {
// 	return &service{
// 		userDb: *user.NewUserDb(),
// 	}
// }

// func (svc *service) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
// 	return &pb.GetUserListResponse{
// 		Users: svc.userDb.GetUsersPayload(),
// 	}, nil
// }

func main() {
	// Create listener
	listener, err := net.Listen(protocol, port)
	if err != nil {
		log.Fatalf("Failed to listen on %s%s\n", protocol, port)
	}

	// Generate servers and services
	grpcServer := grpc.NewServer()
	svc := services.NewUserService()

	// Seed the database
	svc.UserDb.LoadUsersData("./server/db/users.json")

	// Register services
	pb.RegisterUsersServer(grpcServer, svc)
	reflection.Register(grpcServer)

	log.Printf("Server listening on %s%s\n", protocol, port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v\n", err)
	}
}
