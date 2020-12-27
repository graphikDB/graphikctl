package helpers

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(val []byte) string {
	h := sha1.New()
	h.Write(val)
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
