package codex

import (
	"os"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/gemerio/go-gemer/consensus"
)

type Keychain struct {
	publickeys []PublicKey
}

// Generate new keychain

func NewKeychain() *Keychain {
	keychain := &Keychain {
		publickeys:	nil
	}
	return keychain
}

// Generate unique device keys

func GenerateKeys() *PrivateKey {
	pkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return pkey
}

// Write the public key to the keychain

func WriteToKeychain(keychain *Keychain, pubkey *PublicKey) {
	&keychain.publickeys  := append(&keychain.publickeys, &pubkey)
}