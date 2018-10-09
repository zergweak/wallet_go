package main

import (
	"crypto"
	"encoding/hex"
	"fmt"
	_ "../common/crypto/sha3"
)

func main()  {
	input, _ := hex.DecodeString("8F2680A4A96A6009F198F376FC57AB733E76184F88F49F1C45898DFD0EF854D3DCCF5141B40007B924CAB56A86A864E6FB5CCA39625BB5A3F659E2ACD353C304")
	Sha3256H := crypto.SHA3_256.New()
	Sha3256H.Reset()
	Sha3256H.Write(input)
	hash := Sha3256H.Sum(nil)
	fmt.Println(hex.EncodeToString(hash))
}
