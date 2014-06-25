package tools

import (
	"os"
	"path/filepath"
	"time"
)

func GetDateYYYYMMDD() string {
	t := time.Now()
	return t.Format("20060102")
}

func GetDateMMDD() string {
	t := time.Now()
	return t.Format("0102")
}

func GetExecutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return dir
}
