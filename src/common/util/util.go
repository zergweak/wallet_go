package util

import (
	"bytes"
	"common/crypto/secp256k1"
	"crypto/sha256"
	"encoding/binary"
	"protocol/core"
	"errors"
	p "github.com/golang/protobuf/proto"
)

func IntToByte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		return nil
	}
	return buffer.Bytes()
}

func ByteToInt(num []byte) int64 {
	var result int64
	result  = 0
	for i :=0; i < len(num); i++ {
		result *= 256
		result += (int64)(num[i])
	}
	return result
}

func SignTx(trx *core.Transaction, key []byte) error {
	h := sha256.New()
	data, err := p.Marshal(trx.GetRawData())
	if err != nil {
		return errors.New("signTx faild when Marshal cause by :: " + err.Error())
	}
	h.Write(data)
	hash := h.Sum(nil)
	c := secp256k1.S256()
	prikey, _ := secp256k1.PrivKeyFromBytes(c, key)
	signature, err := secp256k1.SignCompact(secp256k1.S256(), (*secp256k1.PrivateKey)(prikey), hash, false)
	if err != nil {
		return errors.New("signTx faild when call SignCompact cause by :: " + err.Error())
	}
	trx.Signature = append(trx.Signature, signature)
	return nil
}