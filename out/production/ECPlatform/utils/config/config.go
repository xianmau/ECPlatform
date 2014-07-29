package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"utils/tools"
)

type Config struct {
	kv map[string]string
}

func NewConfig() *Config {
	conffilepath := path.Join(tools.GetExecutePath(), "conf", "conf.json")
	conffile, err := os.Open(conffilepath)
	defer conffile.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	buf := make([]byte, 1024)
	n, _ := conffile.Read(buf)
	buf = buf[:n]
	kv := make(map[string]string)
	err = json.Unmarshal(buf, &kv)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return &Config{
		kv: kv,
	}
}

func (this *Config) Get(k string) string {
	if _, ok := this.kv[k]; ok {
		return this.kv[k]
	}
	return ""
}
