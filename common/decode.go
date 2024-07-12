package common

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	rsa2 "github.com/deatil/go-cryptobin/cryptobin/rsa"
	"github.com/deatil/go-cryptobin/pkcs12"
)

func ParseRsaPrivatePfx(pfxContent []byte, password string) (privateKey *rsa.PrivateKey, err error) {
	privateK, _, err := pkcs12.Decode(pfxContent, password)
	if err != nil {
		return nil, fmt.Errorf("私钥解析出错err:%s \n", err)
	}
	return privateK.(*rsa.PrivateKey), nil
}

func ParseRsaPrivatePem(pemContent []byte) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(pemContent)
	if block == nil {
		return nil, fmt.Errorf("pem.Decode(%s)：pemContent decode error", pemContent)
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		pk8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("私钥解析出错 [%s]", pemContent)
		}
		var ok bool
		privateKey, ok = pk8.(*rsa.PrivateKey)
		if !ok {
			return nil, fmt.Errorf("私钥解析出错 [%s]", pemContent)
		}
	}
	return privateKey, nil
}

func ParseRsaPublicPem(pemContent []byte) (publicKey *rsa.PublicKey, err error) {
	newRSA := rsa2.NewRSA()
	if publicKey, err = newRSA.ParsePKCS8PublicKeyFromPEM(pemContent); err == nil {
		return
	}
	if publicKey, err = newRSA.ParsePKCS1PublicKeyFromPEM(pemContent); err != nil {
		return nil, fmt.Errorf("公钥解析出错err:%s \n[%s]", err, pemContent)
	}
	return publicKey, nil
}
