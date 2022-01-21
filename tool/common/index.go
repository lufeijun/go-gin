package common

import (
	"crypto/md5"
	"encoding/hex"
)

func StrToMd5(str string) string {
	d := []byte(str)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
