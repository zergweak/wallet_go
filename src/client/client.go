package main

import (
	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"log"
	pb "../api"
	"os"
	"fmt"
	"bufio"
	"strings"
)

const (
	address = "localhost:50051"
)

func CreateWallet(conn *grpc.ClientConn) {
	c := pb.NewWalletApiClient(conn)
	address, err := c.CreateWallet(context.Background(), &pb.PasswordMessage{Password: "123456"})
	if err != nil {
		log.Fatal("CreateWallet faild: %v", err)
	}
	log.Printf("address: %s", address.Address)
}

func Signature(conn *grpc.ClientConn) {
	c := pb.NewWalletApiClient(conn)
	var hash = []byte{0x12, 0x34, 0x56, 0x78, 0x9A}
	r, err := c.Signature(context.Background(), &pb.AddressPasswordHashMessage{Address: "address", Password: "123456", Hash: hash})
	if err != nil {
		log.Fatal("Signature faild: %v", err)
	}
	log.Printf("Signature: %s", r.Signature)
}

func run() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn.Close()
	quit := false
	for {
		if quit {
			break
		}
		fmt.Println("Please input your command!")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		switch strings.ToLower(input.Text()) {
		case strings.ToLower("CreateWallet"):
			CreateWallet(conn)
			break
		case strings.ToLower("Signature"):
			Signature(conn)
			break
		case strings.ToLower("Exit"):
			quit = true
			break
		case strings.ToLower("Quit"):
			quit = true
			break
		default:
			fmt.Println(input.Text() + " is not support!")
			break
		}
	}
}

func main() {
	fmt.Println("wallet_go_client runing.")
	run()
}
