package main

import (
	_ "common/crypto/sha3"
	"common/accounts/keystore"
	"fmt"
	cryp "common/crypto"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

func main()  {

	ks := keystore.NewKeyStore("keystore", veryLightScryptN, veryLightScryptP)
	a, err := ks.NewAccount("foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	address := cryp.B58checkencode(a.Address.Bytes())
	fmt.Println(address)
}
