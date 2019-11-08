package resolver

import (
	"crypto/md5"
	"encoding/hex"
)

func HASH(message string) string {
	digest := md5.New()
	digest.Write([]byte(message))
	return hex.EncodeToString(digest.Sum(nil))
}
