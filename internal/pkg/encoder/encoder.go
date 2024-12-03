package enc

import (
	"encoding/base64"
)

func Encode(uid string) string {
	return base64.StdEncoding.EncodeToString([]byte(uid))
}

func Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(decoded), nil 
}