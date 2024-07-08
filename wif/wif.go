package wif

import (
	"crypto/sha256"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/ripemd160"
)

// Base58 alphabet
var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// GeneratePublicKey generates a compressed public key from a private key.
func GeneratePublicKey(privKeyBytes []byte) []byte {
	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)
	pubKey := privKey.PubKey()
	return pubKey.SerializeCompressed()
}

// PublicKeyToAddress converts a public key to a Bitcoin address.
func PublicKeyToAddress(pubKey []byte) string {
	pubKeyHash := hash160(pubKey)
	versionedPayload := append([]byte{0x00}, pubKeyHash...)
	checksum := doubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58Encode(fullPayload)
}

// PrivateKeyToWIF converts a private key to Wallet Import Format.
func PrivateKeyToWIF(privKey *big.Int) string {
	privKeyBytes := privKey.FillBytes(make([]byte, 32)) // Fill to 32 bytes
	payload := append([]byte{0x80}, privKeyBytes...)
	payload = append(payload, 0x01) // Compressed key indicator
	checksum := doubleSha256(payload)[:4]
	fullPayload := append(payload, checksum...)
	return base58Encode(fullPayload)
}

// doubleSha256 computes SHA-256 twice.
func doubleSha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	secondHash := sha256.Sum256(hash[:])
	return secondHash[:]
}

// hash160 computes RIPEMD-160 after SHA-256.
func hash160(data []byte) []byte {
	sha256Hash := sha256.Sum256(data)
	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(sha256Hash[:])
	return ripemd160Hash.Sum(nil)
}

// base58Encode encodes input bytes to a Base58 string.
func base58Encode(input []byte) string {
	x := new(big.Int).SetBytes(input)
	base := big.NewInt(int64(len(base58Alphabet)))
	zero := big.NewInt(0)
	mod := new(big.Int)

	var result []byte
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append(result, base58Alphabet[mod.Int64()])
	}

	// Leading zero bytes
	for _, b := range input {
		if b == 0x00 {
			result = append(result, base58Alphabet[0])
		} else {
			break
		}
	}

	// Reverse the result
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
