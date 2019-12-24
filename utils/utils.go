package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(text string) string {
	data := []byte(text)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func GenerateToken(userID int, phone, email string) string {
	return ""
}
