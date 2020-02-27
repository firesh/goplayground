package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {

	msg := []byte("")

	// SHA1
	// SHA("") = 0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	hash1 := sha1.Sum(msg)
	fmt.Printf("SHA-1: 0x%v\n", hex.EncodeToString(hash1[:]))

	// SHA-2 SHA-256
	// SHA("") = 0xe3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	hash2 := sha256.Sum256(msg)
	fmt.Printf("SHA-2 SHA-256: 0x%v\n", hex.EncodeToString(hash2[:]))

	// SHA-3 SHA3-256
	// SHA("") = 0xa7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a
	hash3 := sha3.Sum256(msg)
	fmt.Printf("SHA-3 SHA3-256: 0x%v\n", hex.EncodeToString(hash3[:]))

	// SHA-3 Keccak256
	// SHA("") = 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
	hash4 := sha3.NewLegacyKeccak256().Sum(msg)
	fmt.Printf("SHA-3 Keccak256: 0x%v\n", hex.EncodeToString(hash4))
}
