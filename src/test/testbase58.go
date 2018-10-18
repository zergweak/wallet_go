package main

import (
	"common/crypto"
	"encoding/hex"
	"fmt"
)

func main() {
	base58 := "THCSayiq44pGAAoJeeBEdGd36Ta4XRFxwQ"
	bytes := crypto.Base58Decode(base58)
	hexString := hex.EncodeToString(bytes)
	if hexString == "414f49a0c68845a27eb9b94952de1f23171c7efc6dfd46d0bf" {
		fmt.Println("Base58Decode is right!")
	} else {
		fmt.Println("Base58Decode is wrong!")
	}
	base58_1 := crypto.Base58Encode(bytes)

	if base58_1 == base58 {
		fmt.Println("Base58Encode is right!")
	} else {
		fmt.Println("Base58Encode is wrong!")
	}

	hexString = "414F49A0C68845A27EB9B94952DE1F23171C7EFC6D"
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	base58_2 := crypto.B58checkencode(bytes)
	if base58_2 ==base58 {
		fmt.Println("B58checkencode is right!")
	} else {
		fmt.Println("B58checkencode is wrong!")
	}
}
