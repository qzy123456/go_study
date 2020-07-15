package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/hmac"
	"crypto/sha256"
	//"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"tests"
)

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

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	src = PKCS5Padding(src, bs)
	if len(src) % bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src) % bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

//CBC加密
func EncryptDES_CBC(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	data = PKCS5Padding(data, block.BlockSize())
	//获取CBC加密模式
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCEncrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)
	return fmt.Sprintf("%X", out)
}

//CBC解密
func DecryptDES_CBC(src, key string) string {
	keyByte := []byte(key)
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	iv := keyByte //用密钥作为向量(不建议这样使用)
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext)
}

//ECB加密
func EncryptDES_ECB(src, key string) string {
	data := []byte(src)
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	//对明文数据进行补码
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//对明文按照blocksize进行分块加密
		//必要时可以使用go关键字进行并行加密
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return fmt.Sprintf("%X", out)
}

//ECB解密
func DecryptDES_ECB(src, key string) string {
	data, err := hex.DecodeString(src)
	if err != nil {
		panic(err)
	}
	keyByte := []byte(key)
	block, err := des.NewCipher(keyByte)
	if err != nil {
		panic(err)
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = PKCS5UnPadding(out)
	return string(out)
}
func main() {
	//var aeskey = []byte("e4ad2122c3d2b2bc")
	//pass := []byte("57542fc19f98607c")
	//xpass, err := AesEncrypt(pass, aeskey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//pass64 := base64.StdEncoding.EncodeToString(xpass)
	//fmt.Printf("加密后:%v\n",pass64)
	//
	//bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//tpass, err := AesDecrypt(bytesPass, aeskey)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("解密后:%s\n", tpass,"\n")
	//c := getSha256Code("4E531CF2564B4514E2BE78F38854E7F2F6E4B55A6C68C8977CA8A8227CC54F1E",aeskey)
	//fmt.Printf("sha256解密前:%s\n", c,"\n")
	//
	//c = getHmacCode("4E531CF2564B4514E2BE78F38854E7F2F6E4B55A6C68C8977CA8A8227CC54F1E",aeskey)
	//fmt.Printf("sha256解密后:%s\n", c,"\n")
	//
	//strEncrypted, err := DesEncrypt(pass, aeskey)
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println("Encrypted:", strEncrypted)
	//strDecrypted, err := DesDecrypt(strEncrypted, aeskey)
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//fmt.Println("Decrypted:", strDecrypted,"\n")
	//
	////key的长度必须都是8位
	//var key = "12345678"
	//var info = "110 119 120 122"
	//
	//Enc_str := EncryptDES_CBC(info, key)
	//fmt.Println(Enc_str)
	//Dec_str := DecryptDES_CBC(Enc_str, key)
	//fmt.Println(Dec_str)
	//
	//Enc_str = EncryptDES_ECB(info, key)
	//fmt.Println(Enc_str)
	//Dec_str = DecryptDES_ECB(Enc_str, key)
	//fmt.Println(Dec_str,"\n")
	//var str = "123456"
	//strbytes := []byte(str)
	//encoded := base64.StdEncoding.EncodeToString(strbytes)
	//fmt.Println(encoded)
	//decoded, err := base64.StdEncoding.DecodeString("sNpdut6NqS+vdOF/mi8H4g==")
	//decodestr := string(decoded)
	//fmt.Println(decodestr)
  hash := tests.Hash("hash",tests.Key)
  fmt.Println(hash)

}

func getHmacCode(message string, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	//return base64.StdEncoding.EncodeToString([]byte(sha))
	return sha
}


func getSha256Code(s string,secret []byte) string {
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

