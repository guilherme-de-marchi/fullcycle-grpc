package services

import (
	"context"
	"fmt"

	"github.com/Guilherme-De-Marchi/fullcycle-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.GetName())

	return &pb.User{
		Id:    "13579",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
