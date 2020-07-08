package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"testing"
)

//https://studygolang.com/articles/15642

func TestAesCBCEncrypt(t *testing.T) {
	// 原始密码
	passwordByte := []byte("123456")
	// 秘钥 (16,24,32 位数)
	key := []byte("0123456789012345")
	// 秘钥处理 分组
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("加密秘钥块为: %d", block)
	// 秘钥块大小
	blockSize := block.BlockSize()
	log.Printf("加密秘钥块大小为: %d", blockSize)
	// 补码 PKCS7
	padding := blockSize - len(passwordByte)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	padPassword := append(passwordByte, padText...)
	log.Printf("加密补码为: %s", padPassword)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	log.Printf("加密模式为: %s", blockMode)
	// 加密
	cryptPassword := make([]byte, len(padPassword))
	blockMode.CryptBlocks(cryptPassword, padPassword)
	log.Printf("加密密码为: %s", cryptPassword)
	cryptResultPassword := base64.StdEncoding.EncodeToString(cryptPassword)
	log.Printf("最终加密密码为: %s", cryptResultPassword)
}

func TestAesCBCDecrypt(t *testing.T) {
	cryptPasswordByte, err := base64.StdEncoding.DecodeString("MfRbwVwos9/7eIbYauATbQ==")
	if err != nil {
		t.Fatal(err)
	}
	// 秘钥 (16,24,32 位数)
	key := []byte("0123456789012345")
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal(err)
	}
	blockSize := block.BlockSize()
	deBlockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	// 创建数组
	passwordByte := make([]byte, len(cryptPasswordByte))
	deBlockMode.CryptBlocks(passwordByte, cryptPasswordByte)
	t.Logf("passwordByte: %s", passwordByte)
	// 去补全码
	length := len(passwordByte)
	unPadding := int(passwordByte[length-1])
	passwordData := passwordByte[:(length - unPadding)]

	t.Logf("password: %s", passwordData)
}
