package tools

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"
	"regexp"
	"html/template"
)

func GetDateYYYYMMDD() string {
	return time.Now().Format("20060102")
}

func GetDateMMDD() string {
	return time.Now().Format("0102")
}

// 获取程序执行路径
func GetExecutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}

// 生成32位UUID字串
func GetUUID() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	h := md5.New()
	h.Write([]byte(base64.URLEncoding.EncodeToString(b)))
	return hex.EncodeToString(h.Sum(nil))
}

// URL编码
func UrlEncode(s string) string {
	u, _ := url.Parse(s)
	ret := u.Path + "?" + u.Query().Encode()
	return ret
}

// 获取json数据
func GetJsonData(jsonStr string, key string) string {
	re, err := regexp.Compile(`"` + key + `":"(.*?)"`)
	if err != nil {
		return ""
	}
	value := re.FindStringSubmatch(jsonStr)
	if len(value) > 1{
		return value[1]
	}
	return ""
}

// 将被转义的html代码还原
func ConvertToHtml (x string) interface{}{
	return template.HTML(x)
}
