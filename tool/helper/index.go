package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// 一些公用的函数

func StrToMd5(str string) string {
	d := []byte(str)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
