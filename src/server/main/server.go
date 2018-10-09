package main

// server.go

import (
	"net"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../../api"
	"crypto/ecdsa"
	"../../common/tron"
	"../../common/crypto/secp256k1"
	"crypto/rand"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Signature(ctx context.Context, in *pb.AddressPasswordHashMessage) (*pb.SignatureMessage, error) {
	ss := []byte(in.Password)
	ss = append(ss, in.Hash[:]...)
	ss = append(ss, []byte(in.Address)[:]...)
	return &pb.SignatureMessage{Signature: ss}, nil
}

func (s *server) CreateWallet(ctx context.Context, in *pb.PasswordMessage) (*pb.AddressMessage, error) {
	curve := secp256k1.S256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	address := tron.Pub2Address(&priv.PublicKey)

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
