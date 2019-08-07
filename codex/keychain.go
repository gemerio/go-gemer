package codex

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type Keychain struct {
	publickeys []crypto.PublicKey
}

// Generate new keychain

func NewKeychain() *Keychain {
	keychain := &Keychain{
		publickeys: nil,
	}
	return keychain
}

// Generate unique device keys

func GenerateKeys() *ecdsa.PrivateKey {
	pkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return pkey
}

// Write the public key to the keychain

func WriteToKeychain(keychain Keychain, pubkey crypto.PublicKey) {
	keychain.publickeys = append(keychain.publickeys, pubkey)
}
