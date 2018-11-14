package util

import (
	"google.golang.org/grpc"
	"log"
	pb "protocol/api"
)

const (
	wallet_address   = "localhost:50051"
	fullnode_address = "54.236.37.243:50051"
)

var walletCli pb.WalletApiClient

var fullCli pb.WalletClient

func InitClient() {
	connWallet, err := grpc.Dial(wallet_address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
//	defer connWallet.Close()
	connFull, err := grpc.Dial(fullnode_address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect: %v", err)
	}
	//defer connFull.Close()

	fullCli = pb.NewWalletClient(connFull)
	walletCli = pb.NewWalletApiClient(connWallet)
}

func GetWalletCli() pb.WalletApiClient {
	return walletCli
}

func GetFullCli() pb.WalletClient {
	return fullCli
}
