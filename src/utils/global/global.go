package global

import (
	"strconv"
	"time"
	"utils/config"
	"utils/session"
)

var (
	Config   *config.Config
	Sessions *session.Sessions
)

func init() {
	Config = config.NewConfig()
	session_expires, err := strconv.Atoi(Config.Get("session_expires"))
	if err != nil {
		session_expires = 60
	}
	Sessions = session.NewSessions("sessionid", time.Duration(session_expires)*time.Second)
	go Sessions.GC()
}
