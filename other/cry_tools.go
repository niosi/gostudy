/*
package: package tools
author: Wuhaiqiang
date: 2018-06-27 上午 11:32
*/

package main

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/rsa"
	"flag"
	"fmt"
	"encoding/pem"
	"errors"
	"encoding/base64"
	"crypto/sha1"
	"encoding/hex"
)

func main() {
	k1, k2 := GenRsaKey()
	fmt.Printf("%v\n%v", string(k1[:]), string(k2[:]))
	crtData := []byte("18030021264")
	b, err := RsaEncrypt(crtData, k2)
	if err != nil {
		fmt.Println(err.Error())
	}
	cryBase64 := base64.StdEncoding.EncodeToString(b)
	fmt.Println(cryBase64)
	deValue, err := base64.StdEncoding.DecodeString(cryBase64)
	b, err = RsaDecrypt(deValue, k1)

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b[:]))

	fmt.Println(Sha1Encrypt([]byte("HelloWorld"))[:])
}

//RSA公私钥生成
func GenRsaKey() ([]byte, []byte) {
	var bits int
	flag.IntVar(&bits, "b", 2048, "密钥长度，默认为1024位")
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Print(err.Error())
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	pub := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvStream := pem.EncodeToMemory(pub)

	// 生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Print(err.Error())
	}
	prv := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubStream := pem.EncodeToMemory(prv)
	return prvStream, pubStream
}

// 加密
func RsaEncrypt(origData []byte, pubkey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte, prvkey []byte) ([]byte, error) {
	block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

//sha1加密
func Sha1Encrypt(text []byte) string {
	hash := sha1.New()
	hash.Write(text)
	crydata := hash.Sum(nil)
	return hex.EncodeToString(crydata)
}
