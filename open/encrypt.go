package open

import (
	"bytes"
	"crypto/des"
	"errors"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("PKCS5UnPadding error : the input is empty!")
	}
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)], nil
}

func Sm4EncryptECB(key, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

func Sm4DecryptECB(key, data []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out, err = PKCS5UnPadding(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//func main() {
//	key := []byte("1234567812345678")
//	data := []byte("Hello, world!")
//	fmt.Printf("原文: %s\n", data)
//
//	encryptedData, err := Sm4EncryptECB(key, data)
//	if err != nil {
//		fmt.Println("加密失败:", err)
//		return
//	}
//	fmt.Printf("加密后: %s\n", hex.EncodeToString(encryptedData))
//
//	decryptedData, err := Sm4DecryptECB(key, encryptedData)
//	if err != nil {
//		fmt.Println("解密失败:", err)
//		return
//	}
//	fmt.Printf("解密后: %s\n", decryptedData)
//}
