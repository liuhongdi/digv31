package md5

import (
	"crypto/md5"
	"encoding/hex"
)

//返回md5结果
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
