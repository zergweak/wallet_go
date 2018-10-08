package main

// server.go

import (
	"net"
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "../../../gen/api"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto"
	"crypto/sha256"
)

const (
	port = ":50051"
)

type server struct{}
type tronWallet struct{}

func sha3omit12(intput []byte) []byte {
	SHA3_256_h := crypto.SHA3_256.New()
	SHA3_256_h.Reset()
	SHA3_256_h.Write(intput)
	hash := SHA3_256_h.Sum(nil)
	address := hash[10:21]
	address[0] = 0x21
	return address
}

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) string {
	return ""
}

func Base58Decode(string2 string) []byte {
	return nil
}

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	bcpy := append([]byte{ver}, b...)

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 Hash #1 */
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)

	/* SHA256 Hash #2 */
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)

	/* Append first four bytes of hash */
	bcpy = append(bcpy, hash2[0:4]...)

	/* Encode base58 string */
	s = Base58Encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}

	return s
}

func FromECDSAPub(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	byteLen := 32;
	ret := make([]byte, 1+2*byteLen)
	ret[0] = 4 // uncompressed point

	xBytes := pub.X.Bytes()
	copy(ret[1+byteLen-len(xBytes):], xBytes)
	yBytes := pub.Y.Bytes()
	copy(ret[1+2*byteLen-len(yBytes):], yBytes)
	return ret
}

func (t *tronWallet) pub2Address(p *ecdsa.PublicKey) string {
	pubKey := FromECDSAPub(p)
	address := sha3omit12(pubKey)
	base58 := b58checkencode(0x00, address)
	return base58
}

func (s *server) Signature(ctx context.Context, in *pb.AddressPasswordHashMessage) (*pb.SignatureMessage, error) {
	ss := []byte(in.Password)
	ss = append(ss, in.Hash[:]...)
	ss = append(ss, []byte(in.Address)[:]...)
	return &pb.SignatureMessage{Signature: ss}, nil
}

func (s *server) CreateWallet(ctx context.Context, in *pb.PasswordMessage) (*pb.AddressMessage, error) {
	curve := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curve, nil)
	if err != nil {
		return nil, err
	}
	t := new(tronWallet)
	address := t.pub2Address(&privKey.PublicKey)

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
