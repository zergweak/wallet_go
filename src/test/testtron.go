package main

import (
	"fmt"
	"../common/tron"
	"encoding/hex"
	"../common/crypto/secp256k1"
)

func main() {
	c := secp256k1.S256()
	priv, err := secp256k1.NewPrivateKey(c)
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
