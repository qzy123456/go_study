package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	Key = "111111"
	Iv  = "2222"
)

//func main() {
//
//
//
//	str := `{
//	"success": true,
//		"data": {
//		"code": 1,
//			"ice_endAt": 0,
//			"limit_endAt": 0,
//			"notifications": [],
//			"activityMessage": {"isOpen": false,"lastId": ""}
//}
//}
//`
//   //aes加密,并且base64之后的数据
//	hashData, _ := Hash(str, []byte(key))
//	fmt.Println(hashData)
//    //验证解密,判断加密后的数据
//	unHashData, _ := unHash(hashData, []byte(key))
//	fmt.Println(string(unHashData))
//
//}
//验证hash
func Hash(encodeStr string, key []byte) (string, error)  {
	hashData, _ := AesEncrypt(encodeStr, []byte(key))
	//拼接密码 sha256
	sha256Data := getHmacCode(hashData,[]byte(key)) + hashData
	return sha256Data,nil
}

//验证unhash
func unHash(decodeStr string, key []byte) ([]byte, error) {
	hash :=decodeStr[0:64]
	reData := decodeStr[64:]
	fmt.Println(reData)
	//加密后的数据，和加密前的数据作比较，看看是否正常，不正常就返回错误
	newHash := getHmacCode(reData,[]byte(key))
	if hash != newHash{
		fmt.Println("验证出错")
	}
	unHashData, _ := AesDecrypt(reData, []byte(key))
	return unHashData,nil
}
//AES 128 256加密
func AesEncrypt(encodeStr string, key []byte) (string, error) {
	encodeBytes := []byte(encodeStr)
	//根据key 生成密文
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encodeBytes = PKCS5Padding(encodeBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(Iv))
	crypted := make([]byte, len(encodeBytes))
	blockMode.CryptBlocks(crypted, encodeBytes)

	return base64.StdEncoding.EncodeToString(crypted), nil
}


//AES 128 256解密
func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	//先解密base64
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(Iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
//填充，解填充
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
//填充，解填充
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//填充
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}
//sha256加密跟PHP一样
func getHmacCode(message string, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
