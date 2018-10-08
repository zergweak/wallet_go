package crypto

import "crypto/ecdsa"
func FromECDSAPub(byteLen int, pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	ret := make([]byte, 1+2*byteLen)
	ret[0] = 4 // uncompressed point

	xBytes := pub.X.Bytes()
	copy(ret[1+byteLen-len(xBytes):], xBytes)
	yBytes := pub.Y.Bytes()
	copy(ret[1+2*byteLen-len(yBytes):], yBytes)
	return ret
}