package tron

import (
	cryp "../crypto"
	"crypto"
	"crypto/ecdsa"
)

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

func sha3omit12(intput []byte) []byte {
	SHA3_256_h := crypto.SHA3_256.New()
	SHA3_256_h.Reset()
	SHA3_256_h.Write(intput)
	hash := SHA3_256_h.Sum(nil)
	address := hash[10:21]
	address[0] = 0x21
	return address
}

func Pub2Address(p *ecdsa.PublicKey) string {
	pubKey := FromECDSAPub(p)
	address := sha3omit12(pubKey)
	base58 := cryp.B58checkencode(0x00, address)
	return base58
}