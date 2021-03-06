package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// RsaEncrypt 私钥加密数据
func RsaEncrypt(data, key []byte) ([]byte, error) {
	privateKeyBlock, _ := pem.Decode(key)
	if privateKeyBlock == nil {
		return nil, errors.New("Private key format is wrong")
	}
	priv, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return priKeyByte(priv.(*rsa.PrivateKey), data, true)
}

// RsaPubEncrypt 公钥加密数据
func RsaPubEncrypt(data, key []byte) ([]byte, error) {
	publicKeyBlock, _ := pem.Decode(key)
	if publicKeyBlock == nil {
		return nil, errors.New("Public key format is wrong")
	}
	publ, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKeyByte(publ.(*rsa.PublicKey), data, true)
}

// RsaDecrypt 私钥解密数据
func RsaDecrypt(data, key []byte) ([]byte, error) {
	privateKeyBlock, _ := pem.Decode(key)
	if privateKeyBlock == nil {
		return nil, errors.New("Private key format is wrong")
	}
	priv, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return priKeyByte(priv.(*rsa.PrivateKey), data, false)
}

// RsaPubDecrypt 公钥解密数据
func RsaPubDecrypt(data, key []byte) ([]byte, error) {
	publicKeyBlock, _ := pem.Decode(key)
	if publicKeyBlock == nil {
		return nil, errors.New("Public key format is wrong")
	}
	publ, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKeyByte(publ.(*rsa.PublicKey), data, false)
}
