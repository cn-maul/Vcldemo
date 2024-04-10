package base

import "encoding/base64"

func Decode(input string) (error, string) {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return err, ""
	}
	return nil, string(decoded)
}

func Encode(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	return encoded
}
