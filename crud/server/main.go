package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/lyazii22/picnic-asg1/crud/proto"
	"github.com/lyazii22/picnic-asg1/crud/store"
	"google.golang.org/grpc"
)

const (
	port  = ":50051"
	dbURI = "projects/your-project-id/instances/test-instance/databases/crud"
)

type crudServer struct {
	pb.UnimplementedCrudServer
}

func main() {
	store.SetUpSpanner(dbURI)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterCrudServer(s, &crudServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *crudServer) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)
	newUser := store.UserInfo{
		FirstName: in.Firstname,
		LastName:  in.Lastname,
		Id:        strconv.Itoa(rand.Intn(100000000)),
	}

	user, err := store.CreateUser(dbURI, newUser)
	if err != nil {
		return &pb.UserInfo{}, err
	}
	res := &pb.UserInfo{
		Id:        user.Id,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
	}

	return res, nil
}

func (s *crudServer) GetUser(ctx context.Context,
	in *pb.Id) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)

	user, err := store.GetUser(in.Id, dbURI)
	if err != nil {
		return &pb.UserInfo{}, err
	}

	res := &pb.UserInfo{
		Id:        user.Id,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
	}

	return res, nil
}

func (s *crudServer) GetUsers(context.Context, *pb.Empty) (*pb.Users, error) {
	dbUsers, err := store.GetUsers(dbURI)
	if err != nil {
		return &pb.Users{}, err
	}
	users := []*pb.UserInfo{}
	for _, dbUser := range dbUsers {
		user := &pb.UserInfo{
			Id:        dbUser.Id,
			Firstname: dbUser.FirstName,
			Lastname:  dbUser.LastName,
		}
		users = append(users, user)
	}
	return &pb.Users{Users: users}, nil
}

func (s *crudServer) UpdateUser(ctx context.Context,
	in *pb.UserInfo) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	user := store.UserInfo{
		Id:        in.Id,
		FirstName: in.Firstname,
		LastName:  in.Lastname,
	}

	err := store.UpdateUser(dbURI, user)
	if err != nil {
		return &pb.Status{Status: int32(-1)}, err
	} else {
		return &pb.Status{Status: int32(1)}, nil
	}
}

func (s *crudServer) DeleteUser(ctx context.Context,
	in *pb.Id) (*pb.Status, error) {
	log.Printf("Received: %v", in)

	err := store.DeleteUser(dbURI, in.Id)
	if err != nil {
		return &pb.Status{Status: int32(-1)}, err
	} else {
		return &pb.Status{Status: int32(1)}, nil
	}
}
