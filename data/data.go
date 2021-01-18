package data

import (
	"crypto/sha1"
	"fmt"
)

func Encrypt(password string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	return cryptext
}
