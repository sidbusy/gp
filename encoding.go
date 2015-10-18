package gp

import (
	"encoding/base64"
)

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)
	return string(b), err
}
