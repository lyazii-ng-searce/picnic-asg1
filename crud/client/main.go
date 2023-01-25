package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lyazii22/picnic-asg1/crud/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCrudClient(conn)

	// doGetUser(client, "1")
	// doCreateUser(client, "Lya", "me")
	// doUpdateUser(client, "1", "Peter", "Parker")
	doDeleteUser(client, "2")
}

func doCreateUser(client pb.CrudClient, firstName, lastName string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.CreateUserRequest{Firstname: firstName, Lastname: lastName}
	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateUser(_) = _, %v", client, err)
	}
	if res.Id != "" {
		log.Printf("CreateUser: %v", res)
	} else {
		log.Printf("CreateUser Failed")
	}
}

func doGetUser(client pb.CrudClient, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Id: userId}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUser(_) = _, %v", client, err)
	}
	log.Printf("UserInfo: %v", res)
}

func doUpdateUser(client pb.CrudClient, userId string, firstname string, lastname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UserInfo{Id: userId,
		Firstname: firstname, Lastname: lastname}
	res, err := client.UpdateUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateUser(_) = _, %v", client, err)
	}
	if int(res.Status) == 1 {
		log.Printf("UpdateUser Success")
	} else {
		log.Printf("UpdateUser Failed")
	}
}

func doDeleteUser(client pb.CrudClient, userId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Id: userId}
	res, err := client.DeleteUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteUser(_) = _, %v", client, err)
	}
	if int(res.Status) == 1 {
		log.Printf("DeleteUser Success")
	} else {
		log.Printf("DeleteUser Failed")
	}
}
