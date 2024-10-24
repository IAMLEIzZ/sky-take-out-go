package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5DigestAsHex(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
