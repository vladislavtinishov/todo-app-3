package utils

import (
	"crypto/sha1"
	"fmt"
	"os"
)

func GenerateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("USER_PASSWORD_SALT"))))
}
