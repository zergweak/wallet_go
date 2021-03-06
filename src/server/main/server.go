package main

// server.go

import (
	"common/accounts/keystore"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "protocol/api"
	"protocol/core"
)

const (
	port             = ":50051"
	veryLightScryptN = 2
	veryLightScryptP = 1
	root             = "keystore"
)

type server struct{}

func (s *server) Signature(ctx context.Context, in *pb.AddressPasswordHashMessage) (*pb.SignatureMessage, error) {
	ks := keystore.NewKeyStore(root, veryLightScryptN, veryLightScryptP)
	signature, err := ks.SignHashWithPassphrase(in.Address, in.Password, in.Hash)
	if err != nil {
		return nil, err
	}
	return &pb.SignatureMessage{Signature: signature}, err
}

func (s *server) CreateWallet(ctx context.Context, in *pb.PasswordMessage) (*pb.AddressMessage, error) {
	ks := keystore.NewKeyStore(root, veryLightScryptN, veryLightScryptP)
	address, err := ks.NewAccount(in.Password)
	return &pb.AddressMessage{Address: address}, err
}

func (s *server) SignatureTx(ctx context.Context, in *pb.SignatureTxMessage) (*core.Transaction, error) {
	return nil,nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterWalletApiServer(s, &server{})
	s.Serve(lis)
}
