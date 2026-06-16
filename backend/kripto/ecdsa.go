package kripto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"strings"
	"math/big"
)

// Nanti di sini akan dibuat keypair P-256 untuk kurir.
func GenerateKeyPair() (privateKey string, publicKey string) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil{
		return "",""
	}

	privBytes, err := x509.MarshalECPrivateKey(privKey)
	if err != nil{
		return "",""
	}

	privateKeyHex := hex.EncodeToString(privBytes)

	pubBytes, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil{
		return "",""
	}

	publicKeyHex := hex.EncodeToString(pubBytes)

	return privateKeyHex, publicKeyHex
}

// SignData masih berupa skeleton; implementasi ECDSA signature akan diisi nanti.
func SignData(data, privateKey string) string {
	privBytes, err := hex.DecodeString(privateKey)
	if err != nil{
		return ""
	}

	privKey, err := x509.ParseECPrivateKey(privBytes)
	if err != nil{
		return ""
	}

	hashHex := CalculateHash(data)
	hashedData, err := hex.DecodeString(hashHex)
	if err != nil{
		return ""
	}

	r, s, err := ecdsa.Sign(rand.Reader,privKey, hashedData[:])
	if err != nil{
		return ""
	}
	
	signature := r.Text(16) + "." + s.Text(16)

	return signature
}

// VerifySignature masih berupa skeleton; validasi signature akan diisi nanti.
func VerifySignature(data, signature, publicKey string) bool {
	parts := strings.Split(signature, ".")

	if len(parts) != 2 {
		return false
	}

	r, check := new(big.Int).SetString(parts[0], 16)
	if !check {
		return false
	}

	s, check := new(big.Int).SetString(parts[1], 16)
	if !check{
		return false
	}
	pubBytes, err := hex.DecodeString(publicKey)
	if err != nil{
		return false
	}
	pubParsed, err := x509.ParsePKIXPublicKey(pubBytes)
	if err != nil{
		return false
	}

	pubKey, check := pubParsed.(*ecdsa.PublicKey)
	if !check{
		return false
	}

	hashHex := CalculateHash(data)
	hashedData, err := hex.DecodeString(hashHex)
	if err != nil{
		return false
	}
	

	return ecdsa.Verify(pubKey,hashedData[:],r,s)
}
