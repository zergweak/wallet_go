package crypto

import (
	sha256 "crypto/sha256"
)
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func Base58Encode(input []byte) string {
	return ""
}

func Base58Decode(string2 string) []byte {
	return nil
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
