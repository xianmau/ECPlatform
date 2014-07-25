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
	ArticleStatus map[int]string
	LinkStatus map[int]string
	UserStatus map[int]string
	UserLevel map[int]string

	AuthorityList []string
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

	// 初始化权限列表
	AuthorityList = []string{
		"登录",
		"文章管理",
		"链接管理",
		"用户管理",
		"商店管理",
		"商品管理",
		"上传文件",
		"网站设置",
		"系统管理",
		"角色管理",
		"管理员管理",
		"货源管理",
		"帐号管理",
	}

	// 初始化商品状态
	GoodsStatus = map[int]string{
		0: "待售",
		1: "上架",
		2: "缺货",
		3: "下架",
		4: "禁售",
	}

	// 初始化文章状态
	ArticleStatus = map[int]string{
		0: "未公开",
		1: "公开",
	}

	// 初始化链接状态
	LinkStatus = map[int]string{
		0: "未公开",
		1: "公开",
	}

	// 初始化用户状态
	UserStatus = map[int]string{
		0: "冻结",
		1: "未认证",
		2: "已认证",
	}

	// 初始化用户等级
	UserLevel = map[int]string{
		0: "注册用户",
		1: "黄金用户",
		2: "白金用户",
		3: "钻石用户",
	}
}
