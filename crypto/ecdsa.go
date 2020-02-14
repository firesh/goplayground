package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

func main() {
	// generate keys
	privateKey, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: (0x%x, 0x%x, 0x%x)\n", privateKey.D, privateKey.X, privateKey.Y)
	msg := "hello world"
	hash := sha256.Sum256([]byte(msg))

	// sign message with private key
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: (0x%x, 0x%x)\n", r, s)

	// validate message with public key
	valid := ecdsa.Verify(&privateKey.PublicKey, hash[:], r, s)
	fmt.Println("signature verifyied:", valid)
}
