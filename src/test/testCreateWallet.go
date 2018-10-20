package main

import (
	_ "common/crypto/sha3"
	"common/accounts/keystore"
	"fmt"
	"encoding/hex"
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
		key,err := account.GetKey(account.GetAddress(), account.JoinPath(account.GetAddress()),"faa")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(key.Id.String())
		fmt.Println(key.Address)
		fmt.Println(hex.EncodeToString(key.PrivateKey.D.Bytes()))
	}
}
