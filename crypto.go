package gp

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

func MD5(str string) string {
	s := md5.Sum([]byte(str))
	return hex.EncodeToString(s[:])
}

func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	return hex.EncodeToString(s[:])
}
