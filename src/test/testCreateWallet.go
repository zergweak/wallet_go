package main

import (
	_ "common/crypto/sha3"
	"common/accounts/keystore"
	"fmt"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func main()  {

	ks := keystore.NewKeyStore("keystore", veryLightScryptN, veryLightScryptP)
	address, err := ks.NewAccount("foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(address)
}
