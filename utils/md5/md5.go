package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5EnCode(string string) string {
	h := md5.New()
	h.Write([]byte(string)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}