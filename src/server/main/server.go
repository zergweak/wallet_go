package main

// server.go

import (
	"net"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "api"
	"common/accounts/keystore"
)

const (
	port             = ":50051"
	veryLightScryptN = 2
	veryLightScryptP = 1
)

type server struct{}

func (s *server) Signature(ctx context.Context, in *pb.AddressPasswordHashMessage) (*pb.SignatureMessage, error) {
	ss := []byte(in.Password)
	ss = append(ss, in.Hash[:]...)
	ss = append(ss, []byte(in.Address)[:]...)
	return &pb.SignatureMessage{Signature: ss}, nil
}

func (s *server) CreateWallet(ctx context.Context, in *pb.PasswordMessage) (*pb.AddressMessage, error) {
	ks := keystore.NewKeyStore("keystore", veryLightScryptN, veryLightScryptP)
	address, err := ks.NewAccount("foo")
	if nil != err {
		return nil, err
	}
	return &pb.AddressMessage{Address: address}, nil
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
