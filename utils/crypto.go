package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// Crypto : Utils Crypto
type Crypto struct{}

// SHA256 : SHA256 Hash string -> hexacode
func (Crypto) SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}

// MD5 : MD5 Hash string -> hexacode
func (Crypto) MD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	byteSlice := hash.Sum(nil)
	hex := fmt.Sprintf("%x", byteSlice)
	return hex
}
