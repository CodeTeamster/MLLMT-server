package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func MD5Encode(args ...any) string {
	if len(args) == 0 {
		sum := md5.Sum([]byte{})
		return hex.EncodeToString(sum[:])
	}

	if b, err := json.Marshal(args); err == nil {
		sum := md5.Sum(b)
		return hex.EncodeToString(sum[:])
	}

	var s strings.Builder
	for i, a := range args {
		if i > 0 {
			s.WriteString("|")
		}
		fmt.Fprintf(&s, "%v", a)
	}
	sum := md5.Sum([]byte(s.String()))
	return hex.EncodeToString(sum[:])
}
