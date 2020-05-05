package common

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
)

//Encode encodes b as a hex string with 0x prefix.
func encode(b []byte) string {
	enc := make([]byte, len(b)*2+2)
	copy(enc, "0x")
	hex.Encode(enc[2:], b)
	return string(enc)
}

// Decode decodes a hex string with 0x prefix.
func decode(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("empty hex string")
	}
	if !has0xPrefix(input) {
		return nil, errors.New("hex string without 0x prefix")
	}
	b, err := hex.DecodeString(input[2:])
	return b, err
}

//
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

// 随机生成16进制私钥字符串
func RandPk() string {
	priKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	return fmt.Sprintf("0x%x", priKey.D)
}

// 由16进制私钥字符串导出16进制公钥字符串
func Pk2Pub(pk string) string {
	if has0xPrefix(pk) {
		pk = pk[2:]
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return ""
	}
	publicKey := elliptic.Marshal(privateKey.Curve, privateKey.X, privateKey.Y)
	return encode(publicKey)
}

// 由16进制私钥字符串导出16进制地址字符串
func Pk2Addr(pk string) string {
	if has0xPrefix(pk) {
		pk = pk[2:]
	}
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return ""
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	paddedAddress := common.LeftPadBytes(fromAddress.Bytes(), 20)
	return encode(paddedAddress)
}

// 由16进制公钥字符串导出16进制地址字符串
func Pub2Addr(pub string) string {
	if !has0xPrefix(pub) {
		pub = "0x" + pub
	}
	pubBytes, err := decode(pub)
	if err != nil {
		return ""
	}
	publicKey, err := crypto.UnmarshalPubkey(pubBytes)
	if err != nil {
		return ""
	}
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	return encode(fromAddress.Bytes())
}

// 由16进制公钥字符串生成公钥
func PublicKey(pub string) (*ecdsa.PublicKey, error) {
	if !has0xPrefix(pub) {
		pub = "0x" + pub
	}
	pubBytes, err := decode(pub)
	if err != nil {
		return nil, err
	}
	publicKey, err := crypto.UnmarshalPubkey(pubBytes)
	return publicKey, nil
}

//
func PrivateKey(pk string) (*ecdsa.PrivateKey, error) {
	pk = strings.ReplaceAll(pk, "0X", "")
	return crypto.HexToECDSA(pk)
}

// 使用16进制私钥字符串对data数据进行加密数字签名
func Sign(data []byte, privateKey string) (string, error) {
	privateKey = strings.ReplaceAll(privateKey, "0x", "")
	privateKey = strings.ReplaceAll(privateKey, "0X", "")
	_privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	r, s, err := ecdsa.Sign(rand.Reader, _privateKey, data)
	if err != nil {
		return "", err
	}
	publicKey := elliptic.Marshal(_privateKey.Curve, _privateKey.X, _privateKey.Y)
	sign := fmt.Sprintf("%x:%x:%x", r, s, publicKey)
	return sign, nil
}

// 对data数据的加密数字签名进行验证
func Verify(data []byte, sign string) bool {
	array := strings.Split(sign, ":")
	if len(array) != 3 {
		return false
	}
	r, ok := new(big.Int).SetString(array[0], 16)
	if !ok {
		return false
	}
	s, ok := new(big.Int).SetString(array[1], 16)
	if !ok {
		return false
	}
	publicKey, err := PublicKey(array[2])
	if err != nil {
		return false
	}
	return ecdsa.Verify(publicKey, data, r, s)
}
