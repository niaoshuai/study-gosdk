package aes

import (
	"crypto/aes"
	"encoding/base64"
	"log"
	"testing"
)

//https://studygolang.com/articles/15642

func TestAesECBEncrypt(t *testing.T) {
	// 原始密码
	passwordByte := []byte("123456")
	// 秘钥 (16,24,32 位数)
	key := []byte("0123456789012345")
	// generateKey
	generateKey := make([]byte, 16)
	copy(generateKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			generateKey[j] ^= key[i]
		}
	}
	cipher, _ := aes.NewCipher(generateKey)
	length := (len(passwordByte) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, passwordByte)
	pad := byte(len(plain) - len(passwordByte))
	for i := len(passwordByte); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(passwordByte); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	log.Printf("最终加密密码为: %s", base64.StdEncoding.EncodeToString(encrypted))
}

func TestAesECBDecrypt(t *testing.T) {
	encryptedByte, _ := base64.StdEncoding.DecodeString("hZN/2kM7eU+OyaPjKQfIxg==")
	key := []byte("0123456789012345")
	// generateKey
	generateKey := make([]byte, 16)
	copy(generateKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			generateKey[j] ^= key[i]
		}
	}
	cipher, _ := aes.NewCipher(generateKey)
	decrypted := make([]byte, len(encryptedByte))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encryptedByte); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encryptedByte[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	t.Logf("password: %s", decrypted[:trim])
}
