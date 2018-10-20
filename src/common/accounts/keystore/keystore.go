package keystore

import (
	"common/accounts"
	"errors"
	"path/filepath"
	crand "crypto/rand"
	"os"
	"io/ioutil"
	"strings"
	"github.com/sasaxie/go-client-api/common/crypto"
)

var (
	ErrLocked  = accounts.NewAuthNeededError("password or unlock")
	ErrNoMatch = errors.New("no key for given address or file")
	ErrDecrypt = errors.New("could not decrypt key with given passphrase")
)

// KeyStore manages a key storage directory on disk.
type KeyStore struct {
	Accounts    []keyStore // Storage backend, might be cleartext or encrypted
	keysDirPath string
	scryptN     int
	scryptP     int
}

func (ks *KeyStore) walkfunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}
	keyjson, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	temp := strings.Split(path, "/")
	account := keyStorePassphrase{ks.keysDirPath, temp[len(temp)-1], ks.scryptN, ks.scryptP, false, keyjson}
	ks.Accounts = append(ks.Accounts, account)
	return nil
}

func (ks *KeyStore) init() {

	//if keystore exist loadit else create it
	kinfo, err := os.Stat(ks.keysDirPath)
	if err != nil || kinfo == nil || !kinfo.IsDir() {
		_ = os.Mkdir(ks.keysDirPath, 0700)
	}
	filepath.Walk(ks.keysDirPath, ks.walkfunc)
}

func (ks *KeyStore) find(address string) keyStore {
	for i := 0; i < len(ks.Accounts); i++ {
		account := ks.Accounts[i]
		if address == account.GetAddress(){
			return account
		}
	}
	return nil
}

// SignHash calculates a ECDSA signature for the given hash. The produced
// signature is in the [R || S || V] format where V is 0 or 1.
func (ks *KeyStore) SignHash(a accounts.Account, hash []byte) ([]byte, error) {

	//unlockedKey, found := ks.unlocked[a.Address]
	//if !found {
	//	return nil, ErrLocked
	//}
	//// Sign the hash using plain ECDSA operations
	//return crypto.Sign(hash, unlockedKey.PrivateKey)
	return nil, nil
}

// SignHashWithPassphrase signs hash if the private key matching the given address
// can be decrypted with the given passphrase. The produced signature is in the
// [R || S || V] format where V is 0 or 1.
func (ks *KeyStore) SignHashWithPassphrase(address, passphrase string, hash []byte) (signature []byte, err error) {
	account := ks.find(address)
	if account == nil {
		 return  nil, ErrNoMatch
	}
	key, err := account.GetKey(address, "", passphrase)
	if err != nil {
		return nil, err
	}
	defer zeroKey(key.PrivateKey)
	return crypto.Sign(hash, key.PrivateKey)
}

// NewKeyStore creates a keystore for the given directory.
func NewKeyStore(keydir string, scryptN, scryptP int) *KeyStore {
	keydir, _ = filepath.Abs(keydir)
	ks := &KeyStore{keysDirPath: keydir, scryptN: scryptN, scryptP: scryptP}
	ks.init()
	return ks
}

// NewAccount generates a new key and stores it into the key directory,
// encrypting it with the passphrase.
func (ks *KeyStore) NewAccount(passphrase string) (string, error) {
	account := keyStorePassphrase{ks.keysDirPath, "", ks.scryptN, ks.scryptP, false, nil}
	_, address, err := storeNewKey(account, crand.Reader, passphrase)
	if err != nil {
		return "", err
	}
	account.address = address
	ks.Accounts = append(ks.Accounts, account)
	// Add the account to the cache immediately rather
	// than waiting for file system notifications to pick it up.
	return address, nil
}
