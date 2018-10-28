package main

import (
	"bufio"
	"common/crypto"
	"encoding/hex"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	pb "protocol/api"
	"protocol/core"
	"strings"
)

const (
	wallet_address   = "localhost:50051"
	fullnode_address = "54.236.37.243:50051"
)

func inputpassword() string {
	for {
		fmt.Println("Please input your password!")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		password := input.Text()

		fmt.Println("Please input your password again!")
		input = bufio.NewScanner(os.Stdin)
		input.Scan()
		password_1 := input.Text()

		if password == password_1 {
			return password
		}
		fmt.Println("Twice password diffrent, try again!")
	}
}

func CreateWallet(conn *grpc.ClientConn) {
	pwd := inputpassword()
	c := pb.NewWalletApiClient(conn)
	address, err := c.CreateWallet(context.Background(), &pb.PasswordMessage{Password: pwd})
	if err != nil {
		log.Printf("CreateWallet faild: %v", err)
	} else {
		log.Printf("address: %s", address.Address)
	}
}

func Signature(conn *grpc.ClientConn, parameters []string) {
	c := pb.NewWalletApiClient(conn)
	if len(parameters) != 2 {
		log.Println("Signature need 2 parameters : address and passphrase.")
		return
	}
	var hash = []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	r, err := c.Signature(context.Background(), &pb.AddressPasswordHashMessage{Address: parameters[0], Password: parameters[1], Hash: hash})
	if err != nil {
		log.Printf("Signature faild: %v", err)
	} else {
		log.Printf("Signature: %s", hex.EncodeToString(r.Signature))
	}
}

func GetNowBlock(conn *grpc.ClientConn) {
	c := pb.NewWalletClient(conn)
	r, err := c.GetNowBlock2(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		log.Printf("GetNowBlock faild: %v", err)
	} else {
		log.Printf("Current block number : %d", r.BlockHeader.GetRawData().GetNumber())
	}
}

func GetAccount(conn *grpc.ClientConn, parameters []string) {
	c := pb.NewWalletClient(conn)
	if len(parameters) != 1 {
		log.Println("GetAccount need 1 parameters : address.")
		return
	}
	fmt.Println(parameters[0])
	address := crypto.B58checkdecode(parameters[0])
	r, err := c.GetAccount(context.Background(), &core.Account{Address:address})
	if err != nil {
		log.Printf("GetAccount faild: %v", err)
	} else {
		log.Printf("Account of address %s is balance is %d", crypto.B58checkencode(r.GetAddress()), r.GetBalance())
	}
}

func run() {
	conn_0, err := grpc.Dial(wallet_address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn_0.Close()
	conn_1, err := grpc.Dial(fullnode_address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	defer conn_1.Close()
	quit := false
	for {
		if quit {
			break
		}
		fmt.Println("Please input your command!")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		cmdArray := strings.Split(input.Text(), " ")
		cmd := cmdArray[0]
		parameters := cmdArray[1:]
		switch strings.ToLower(cmd) {
		case strings.ToLower("CreateWallet"):
			CreateWallet(conn_0)
			break
		case strings.ToLower("Signature"):
			Signature(conn_0,parameters)
			break
		case strings.ToLower("GetNowBlock"):
			GetNowBlock(conn_1)
			break
		case strings.ToLower("GetAccount"):
			GetAccount(conn_1,parameters)
			break
		case strings.ToLower("Exit"):
			quit = true
			break
		case strings.ToLower("Quit"):
			quit = true
			break
		default:
			fmt.Println(cmd + " is not support!")
			break
		}
	}
}

func main() {
	fmt.Println("wallet_go_client runing.")
	run()
}
