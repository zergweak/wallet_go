package tron

import (
	cryp "../crypto"
	"crypto"
	"crypto/ecdsa"
	_ "../crypto/sha3"
	"fmt"
	"encoding/hex"
)

func sha3omit12(intput []byte) []byte {
	Sha3256H := crypto.SHA3_256.New()
	Sha3256H.Reset()
	Sha3256H.Write(intput)
	hash := Sha3256H.Sum(nil)
	address := hash[10:31]
	address[0] = 0x41
	return address
}

func Pub2Address(p *ecdsa.PublicKey) string {
	pubKey := cryp.FromECDSAPub(32, p)
	fmt.Println(hex.EncodeToString(pubKey))
	address := sha3omit12(pubKey)
	base58 := cryp.B58checkencode(address)
	return base58
}
