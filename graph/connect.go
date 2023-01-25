package graph

import (
	"log"

	// pb "github.com/lyazii22/picnic-asg1/crud/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func connectToGRPC() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	// return pb.NewCrudClient(conn)
}
