package tools

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/freefishgo/freefishgo"
)

func FormaterWrite(w io.Writer, s interface{}) error {
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	return e.Encode(s)
}

//生成Guid字串
func GetGuid() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

/**
 * 对一个字符串进行MD5加密,不可解密
 */
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s)) //使用zhifeiya名字做散列值，设定后不要变
	return hex.EncodeToString(h.Sum(nil))
}

// 写文件
func WriteFile(path string, b []byte) error {
	err := os.MkdirAll(filepath.Dir(path), 0644)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.Write(b)
	defer f.Close()
	return err
}

// 写文件
func WriteFileFromReader(path string, r io.Reader) error {
	err := os.MkdirAll(filepath.Dir(path), 0644)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	io.Copy(f, r)
	defer f.Close()
	return err
}

// 判断用户是否登录
func IsLogin(ir freefishgo.IResponse, r *freefishgo.Request) bool {
	userinfo, _ := ir.GetSession("userinfo")
	userip, _ := ir.GetSession("userip")
	if userinfo != nil && userip != nil {
		if userip != r.Request.Host {
			return false
		}
		return true
	} else {
		return false
	}
}
