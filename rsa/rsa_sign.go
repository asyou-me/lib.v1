package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func Sign(src []byte, privateKey string, hash crypto.Hash) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("Failed to parse RSA PKCS8 private key")
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("Failed to parse RSA PKCS8 private key: " + err.Error())
	}
	rsaPrivate, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("Failed to parse RSA PKCS8 private key")
	}
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.SignPKCS1v15(rand.Reader, rsaPrivate, hash, hashed)
}

func Verify(src []byte, publicKey string, sign []byte, hash crypto.Hash) error {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("Failed to parse RSA public key ")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return errors.New("Failed to parse RSA public key: " + err.Error())
	}
	rsaPub, _ := key.(*rsa.PublicKey)

	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)

	return rsa.VerifyPKCS1v15(rsaPub, hash, hashed, sign)
}
