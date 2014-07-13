package global

import (
	"strconv"
	"time"
	"utils/config"
	log "utils/logger"
	"utils/session"
)

var (
	Config      *config.Config
	Sessions    *session.Sessions
	GoodsStatus map[int]string
)

func init() {
	log.Trace("global init")

	// 初始化配置信息
	Config = config.NewConfig()

	// 初始化Session
	session_expires, err := strconv.Atoi(Config.Get("session_expires"))
	if err != nil {
		session_expires = 60
	}
	Sessions = session.NewSessions("sessionid", time.Duration(session_expires)*time.Second)
	go Sessions.GC()

	// 初始化商品状态
	GoodsStatus = map[int]string{
		0: "待售",
		1: "上架",
		2: "缺货",
		3: "下架",
		4: "禁售",
	}
}
