package main

import (
	_ "common/crypto/sha3"
	"common/accounts/keystore"
	"fmt"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
	root             = "keystore"
)

func main() {

	ks := keystore.NewKeyStore(root, veryLightScryptN, veryLightScryptP)
	address, err := ks.NewAccount("foo")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(address)
	for i := 0; i< len(ks.Accounts); i++  {
		account := ks.Accounts[i]
		fmt.Println(account.GetAddress())
	}
}
