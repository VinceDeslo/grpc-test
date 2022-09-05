package services

import (
	"context"

	"github.com/VinceDeslo/grpc-test/server/user"
	"github.com/VinceDeslo/grpc-test/server/user/pb"
)

type UserService struct {
	pb.UnimplementedUsersServer
	UserDb user.UserDb
}

func NewUserService() *UserService {
	return &UserService{
		UserDb: *user.NewUserDb(),
	}
}

func (svc *UserService) GetUserList(ctx context.Context, in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	return &pb.GetUserListResponse{
		Users: svc.UserDb.GetUsersPayload(),
	}, nil
}
