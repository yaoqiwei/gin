package cryp

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesECBEncrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}
	ecb := NewECBEncryptEr(block)
	// 加PKCS7填充
	content := PKCS7Padding(data, block.BlockSize())
	encryptData := make([]byte, len(content))
	// 生成加密数据
	ecb.CryptBlocks(encryptData, content)
	return encryptData, nil
}

type ecb struct {
	b         cipher.Block
	blockSize int
}
type ecbEncryptEr ecb

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecbEncryptEr) BlockSize() int { return x.blockSize }

func (x *ecbEncryptEr) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

func NewECBEncryptEr(b cipher.Block) cipher.BlockMode {
	return (*ecbEncryptEr)(newECB(b))
}

func AesEncryptCBC(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	str, _ := AesEncryptByte(origData, key)
	return base64.StdEncoding.EncodeToString(str)
}

func AesDecryptCBC(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	str, _ := AesDecryptByte(crytedByte, key)
	return string(str)
}

// AES 加解密
// 这个工具类采用的是CBC分组模式

func AesEncryptByte(origData []byte, key string) ([]byte, error) {
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return []byte(""), err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}
func AesEncrypt(plaintext string, key string) (string, error) {
	encryptByte, err := AesEncryptByte([]byte(plaintext), key)
	if err != nil {
		return "", err
	}
	return string(encryptByte), nil
}
func AesDecryptByte(crytedByte []byte, key string) ([]byte, error) {
	// 转成字节数组
	k := []byte(key)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return []byte(""), err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return orig, nil
}
func AesDecrypt(cryted string, key string) (string, error) {
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return "", err
	}
	decryptByte, err := AesDecryptByte(crytedByte, key)
	if err != nil {
		return "", err
	}
	return string(decryptByte), nil
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PPVODAesEncryptByte(origData []byte, rawKey string) ([]byte, error) {

	genKeyData := ""

	for len(genKeyData) < 32 {
		genKeyData += string(MD5SUM(genKeyData + rawKey))
	}

	generatedKey := genKeyData[0:16]
	generatedIV := genKeyData[16:32]

	// 分组秘钥
	block, err := aes.NewCipher([]byte(generatedKey))
	if err != nil {
		return []byte(""), err
	}
	// 补全码
	origData = PKCS7Padding(origData, 16)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, []byte(generatedIV))
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return cryted, nil
}
