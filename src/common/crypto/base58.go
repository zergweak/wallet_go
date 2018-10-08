package crypto

import (
	"crypto/sha256"
	"math/big"
	"bytes"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func ReverseBytes(input []byte) {
	max := len(input)/2
	for i :=0 ; i < max; i++{
		b := input[i];
	}
}

func Base58Encode(input []byte) string {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)

	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)

		} else {
			break
		}
	}

	return string(result[:])
}

func Base58Decode(str string) []byte {
	input := []byte(str)
	result := big.NewInt(0)
	zeroBytes := 0

	for _, b := range input {
		if b != b58Alphabet[0] {
			break
		}

		zeroBytes++
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(int64(len(b58Alphabet))))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded
}

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func B58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	bcpy := append([]byte{ver}, b...)

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 Hash #1 */
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)

	/* SHA256 Hash #2 */
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)

	/* Append first four bytes of hash */
	bcpy = append(bcpy, hash2[0:4]...)

	/* Encode base58 string */
	s = Base58Encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}

	return s
}
