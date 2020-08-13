package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"hash"
)

const (
	// SHA1 is the name of sha1 hash alg
	SHA1 = "sha1"
	// SHA256 is the name of sha256 hash alg
	SHA256 = "sha256"
)

// HashAlg used to get correct alg for hash
var HashAlg = map[string]func() hash.Hash{
	SHA1:   sha1.New,
	SHA256: sha256.New,
}

// Encrypt encrypts the content with salt
func Encrypt(content string, salt string, encrptAlg string) string {
	return fmt.Sprintf("%x", pbkdf2.Key([]byte(content), []byte(salt), 4096, 16, HashAlg[encrptAlg]))
}