package dsafe

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

func unPadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

func EncryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}

func DecryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unPadding(src)
	return src
}

func Base64AESEncrypt(src []byte, key []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(EncryptAES(src, key)))
}

func Base64AESDecrypt(src []byte, key []byte) ([]byte, error) {
	d, err := base64.StdEncoding.DecodeString(string(src))
	if err != nil {
		return nil, err
	}
	return DecryptAES(d, key), nil
}
