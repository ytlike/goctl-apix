package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// 生成密码
func GenPwd(pwd string) (string, error) {
	// 加密处理
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

// 比对密码
func ComparePwd(pwd1 string, pwd2 string) bool {
	// pwd1：数据库中的密码
	// pwd2：用户输入的密码
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
