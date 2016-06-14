package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func AesEncrypt(origData, key []byte) (crypted []byte, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New("AesEncrypt error")
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	origData = PKCS5Padding(origData, blockSize)
	// 必须为block.Blocksize的倍数
	if len(origData)%blockSize != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted = make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) (origData []byte, err error) {
	defer func() {
		if recoverErr := recover(); recoverErr != nil {
			err = errors.New("AesDecrypt error")
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData, err = PKCS5UnPadding(origData)
	if err != nil {
		return origData, err
	}
	return origData, nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	if length-unpadding < 0 {
		return origData, errors.New("aes解密错误")
	}
	return origData[:(length - unpadding)], nil
}
