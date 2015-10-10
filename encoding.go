package gp

import (
	"encoding/base64"
)

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}
