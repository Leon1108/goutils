package goutils

import (
	"bytes"
	"errors"
	"crypto/aes"
	"crypto/cipher"
)

// 5和7貌似是一样的
// AES/CBC/PKCS5Padding
// AES/CBC/PKCS7Padding

func EncryptAES(plaintext, key, iv []byte) (ciphertext []byte, err error) {
	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	plaintext = PKCS5Padding(plaintext, aes.BlockSize)
	ciphertext = make([]byte, len(plaintext))
	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(ciphertext, plaintext)
	return
}

func DecryptAES(ciphertext, key, iv []byte) (plaintext []byte, err error) {
	defer func() {
		if obj := recover(); obj != nil {
			err = errors.New(obj.(string))
		}
	}()

	var block cipher.Block
	if block, err = aes.NewCipher(key); err != nil {
		return
	}
	plaintext = make([]byte, len(ciphertext))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(plaintext, ciphertext)
	plaintext = PKCS5UnPadding(plaintext)
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
