package main

import (
	"fmt"
	"../common/tron"
	"encoding/hex"
	"../common/crypto/secp256k1"
	"strings"
)

func main() {
	c := secp256k1.S256()
	str := "8422927AE2450CB601C42F7D56ED301D0B9E8B567818375A042A388D7A6C2BF1"
	bytes, _ := hex.DecodeString(str)
	priv, pubk := secp256k1.PrivKeyFromBytes(c, bytes)

	bytes = priv.D.Bytes()
	if strings.ToLower(str) == hex.EncodeToString(bytes) {
		fmt.Println("Private key is same!")
	} else {
		fmt.Println("Private key is not same!")
	}

	xbytes := pubk.X.Bytes()
	ybytes := pubk.Y.Bytes()
	if strings.ToLower("4D7996FBF2ECCBB9947C9B576C912079C29912EBCCE70809C7CCB404F9E9BF0E") == hex.EncodeToString(xbytes) && strings.ToLower("B1B9D9EA80790FFFE786177FC84989C99703FEFA87675589091F50C4EAFE3699") == hex.EncodeToString(ybytes) {
		fmt.Println("Public key is same!")
	} else {
		fmt.Println("Public key is not same!")
	}

	address := tron.Pub2Address(&priv.PublicKey)
	if "TXwxqnsqsuX8ohht8EjbVygimcDWivFoVs" == address {
		fmt.Println("Address is same!")
	} else {
		fmt.Println("Address is not same!")
	}
}
