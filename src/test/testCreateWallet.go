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
	var hash = []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}

	for i := 0; i< len(ks.Accounts); i++  {
		account := ks.Accounts[i]
		fmt.Println(account.GetAddress())
		key,err := account.GetKey(account.GetAddress(), account.JoinPath(account.GetAddress()),"foo")
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(key.Id.String())
		fmt.Println(key.Address)
		fmt.Println(hex.EncodeToString(key.PrivateKey.D.Bytes()))

		signature, err := ks.SignHashWithPassphrase(account.GetAddress(), "foo",hash)
		fmt.Println(hex.EncodeToString(signature))
	}
}
