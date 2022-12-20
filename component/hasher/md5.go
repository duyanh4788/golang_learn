package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct{}

func NewMd5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hashes := md5.New()
	hashes.Write([]byte(data))
	return hex.EncodeToString(hashes.Sum(nil))
}
