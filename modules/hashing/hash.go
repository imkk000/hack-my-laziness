package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

func hashMD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func hashSHA1(data string) string {
	hash := sha1.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

func hashSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func hashSHA512(data string) string {
	hash := sha512.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

func hashSHA3_256(data string) string {
	hash := sha3.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func hashBLAKE2b(data string) string {
	hash := blake2b.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
