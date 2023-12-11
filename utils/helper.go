package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func GenerateImageFilename(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "_")
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
