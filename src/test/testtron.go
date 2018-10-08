package main

import (
	"fmt"
	"crypto/ecdsa"
	"crypto/elliptic"
	"../common/tron"
	"crypto/rand"
	"encoding/hex"
)

func main() {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	bytes := priv.D.Bytes()
	fmt.Println(hex.EncodeToString(bytes))
	xbytes := priv.PublicKey.X.Bytes()
	fmt.Println(hex.EncodeToString(xbytes))
	ybytes := priv.PublicKey.Y.Bytes()
	fmt.Println(hex.EncodeToString(ybytes))
	address := tron.Pub2Address(&priv.PublicKey)
	fmt.Println(address)
}
