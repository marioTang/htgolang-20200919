package utils

import (
	"crypto/md5"
	"fmt"
)

// 将传递的string转换为MD5加密后的十六进制返回
func Md5sum(txt string) (md5Value string) {
	md5Value = fmt.Sprintf("%X", md5.Sum([]byte(txt)))
	return
}