package main

import (
	"common/crypto"
	"common/log"
	"common/util"
	"context"
	"fmt"
	"protocol/core"
)

func getBalance(addres string) (b int64, fb int64, err error) {
	f := util.GetFullCli()
	from := crypto.B58checkdecode(addres)
	acount := core.Account{Address: from}
	acount1, err := f.GetAccount(context.Background(), &acount)
	if err != nil {
		return 0, 0, err
	}
	b = acount1.GetBalance()/1000000
	if len(acount1.GetFrozen()) > 0 {
		fb = acount1.GetFrozen()[0].GetFrozenBalance()/1000000
	}
	asset := "XXXXXX"
	if acount1.AssetIssuedName != nil && len(acount1.AssetIssuedName) > 0 {
		asset = string(acount1.AssetIssuedName)
	}
	fmt.Printf("address : %s, asset : %s, balance : %d, frozen %d\n", addres, asset, b, fb)
	if len(acount1.GetVotes()) > 0 {
		fmt.Printf("vote to %s : %d\n", crypto.B58checkencode(acount1.GetVotes()[0].VoteAddress), acount1.GetVotes()[0].VoteCount)
	}
	return b, fb, nil
}
func main() {
	fmt.Println("getbalance is runing.")
	util.InitClient()
	log.Init()
	var total, totalfrez, balance, frez int64
	balance, frez, _ = getBalance("TJCnKsPa7y5okkXvQAidZBzqx3QyQ6sxMW")
	total = balance
	totalfrez = frez
	balance, frez, _ = getBalance("TWvzaTM8UrpSeXufseXJqtkfc5GcRieprF")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TCTb1zM4zQ6Ai5XmU9Lb19NEmL6HToxs86")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TL5mpGbtr5L2Gi7CtotBQzjN8pK7SmbyFz")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TRonANwbu14QM4Pnu3PdUPsEMpUw2bcssg")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TJ66qRHiBbrCx1LmgBMm7mJp1Vc8MsmLfq")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TTbioAsbefqWxtoGk3MsKkUw7jgtFKPH9E")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TV9Qi1uauZAt8KZWopHV1LbQd5RrLBrS3V")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TRXnA3LdY5LqFatpLPpyYFYmKyJJCB3ZzR")
	total += balance
	totalfrez += frez
	balance, frez, _ = getBalance("TKkJnDF8mRH24YZ7VKBfxhBUQrp1xn8c2S")
	total += balance
	totalfrez += frez
	fmt.Printf("balance :%d, frez : %d, total : %d", total, totalfrez, total+totalfrez)
	log.Printf("balance :%d, frez : %d, total : %d", total, totalfrez, total+totalfrez)
}
