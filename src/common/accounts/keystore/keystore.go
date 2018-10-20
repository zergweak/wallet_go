package keystore

import (
	"common"
	"common/accounts"
	"errors"
	"path/filepath"
	"github.com/sasaxie/go-client-api/common/crypto"
	crand "crypto/rand"
)

var (
	ErrLocked  = accounts.NewAuthNeededError("password or unlock")
	ErrNoMatch = errors.New("no key for given address or file")
	ErrDecrypt = errors.New("could not decrypt key with given passphrase")
)

// KeyStore manages a key storage directory on disk.
type KeyStore struct {
	storage  keyStore                     // Storage backend, might be cleartext or encrypted
	changes  chan struct{}                // Channel receiving change notifications from the cache
	unlocked map[common.Address]*unlocked // Currently unlocked account (decrypted private keys)

	wallets []accounts.Wallet // Wallet wrappers around the individual key files
	//	updateFeed  event.Feed              // Event feed to notify wallet additions/removals
	//	updateScope event.SubscriptionScope // Subscription scope tracking current live listeners
	updating bool // Whether the event notification loop is running

}

type unlocked struct {
	*Key
	abort chan struct{}
}

func newAccountCache(keydir string) (*accountCache, chan struct{}) {
	ac := &accountCache{
		keydir: keydir,
		byAddr: make(map[common.Address][]accounts.Account),
		notify: make(chan struct{}, 1),
		//	fileC:  fileCache{all: mapset.NewThreadUnsafeSet()},
	}
	return ac, ac.notify
}

func (ks *KeyStore) init(keydir string) {

	// Create the initial list of wallets from the cache
}

// SignHash calculates a ECDSA signature for the given hash. The produced
// signature is in the [R || S || V] format where V is 0 or 1.
func (ks *KeyStore) SignHash(a accounts.Account, hash []byte) ([]byte, error) {


	unlockedKey, found := ks.unlocked[a.Address]
	if !found {
		return nil, ErrLocked
	}
	// Sign the hash using plain ECDSA operations
	return crypto.Sign(hash, unlockedKey.PrivateKey)
}

// SignHashWithPassphrase signs hash if the private key matching the given address
// can be decrypted with the given passphrase. The produced signature is in the
// [R || S || V] format where V is 0 or 1.
func (ks *KeyStore) SignHashWithPassphrase(a accounts.Account, passphrase string, hash []byte) (signature []byte, err error) {
	//_, key, err := ks.getDecryptedKey(a, passphrase)
	//if err != nil {
	//	return nil, err
	//}
	//defer zeroKey(key.PrivateKey)
	//return crypto.Sign(hash, key.PrivateKey)
	return  nil,nil
}

// NewKeyStore creates a keystore for the given directory.
func NewKeyStore(keydir string, scryptN, scryptP int) *KeyStore {
	keydir, _ = filepath.Abs(keydir)
	//if keystore exist
	//loadit
	//else create it
	ks := &KeyStore{storage: &keyStorePassphrase{keydir, scryptN, scryptP, false}}
	ks.init(keydir)
	return ks
}

// NewAccount generates a new key and stores it into the key directory,
// encrypting it with the passphrase.
func (ks *KeyStore) NewAccount(passphrase string) (string, error) {
	_, address, err := storeNewKey(ks.storage, crand.Reader, passphrase)
	if err != nil {
		return "", err
	}
	// Add the account to the cache immediately rather
	// than waiting for file system notifications to pick it up.
	return address, nil
}
