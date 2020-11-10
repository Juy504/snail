package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func PasswordEncrypt(password string) (string, string){
	/*
	密码加盐加密
	*/
	salt := time.Now().Unix()
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st), string(salt)
}

func VerifyPassword(password, hashcode, salt string) bool {
	/*
	密码校验
	*/
	m5 := md5.New()
	m5.Write([]byte(password))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	if code := hex.EncodeToString(st); hashcode == code{
		return true
	}
	return false
}
