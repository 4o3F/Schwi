package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func PasswordCrypto(pwd string) string {
	tmp := md5.Sum([]byte(pwd))
	nextpassword := fmt.Sprintf("%x", tmp)
	res := md5.Sum([]byte(nextpassword + "%233@schwi.me+whiteroomelite666?"))
	newpassword := fmt.Sprintf("%x", res)
	return newpassword
}

func DoMD5(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
