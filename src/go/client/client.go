package main

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	pb "../../gen"
)

const (
	address     = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWalletApiClient(conn)
	var hash = []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	r, err := c.Signature0(context.Background(), &pb.SignAddPassHashMessage{Address:"address", Password:"123456", Hash:hash})
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	log.Printf("Signature: %s", r.Signature)
}
