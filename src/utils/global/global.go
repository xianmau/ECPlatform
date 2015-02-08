package global

import (
	"strconv"
	"time"
	log "utils/glog"
	"utils/session"
	"os"
	"flag"
	"encoding/json"
)

var (
	configFilePath = flag.String("config_file_path", "conf/conf.json", "If non-empty, read this config file")
)

var (
	Config        map[string]string
	Sessions      *session.Sessions
	GoodsStatus   map[int]string
	ArticleStatus map[int]string
	LinkStatus    map[int]string
	UserStatus    map[int]string
	UserLevel     map[int]string
	ShopStatus    map[int]string
	ShopKind      map[int]string

	AuthorityList []string
)

func Init() {

	// 加载配置文件
	loadConfigFile()

	// 初始化Session
	session_expires, err := strconv.Atoi(Config["session_expires"])
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

	// 初始化商店状态
	ShopStatus = map[int]string{
		0: "停止营业",
		1: "正常营业",
	}

	// 初始化商店类型
	ShopKind = map[int]string{
		0: "个体户",
		1: "合作社",
	}
}

func loadConfigFile() {
	configFile, err := os.Open(*configFilePath)
	defer configFile.Close()
	if err != nil {
		log.Fatalf("Loading config file `%s` failed: %s\n", *configFilePath, err.Error())
		os.Exit(1)
	}
	buf := make([]byte, 1024)
	n, _ := configFile.Read(buf)
	buf = buf[:n]
	kv := make(map[string]string)
	err = json.Unmarshal(buf, &kv)
	if err != nil {
		log.Fatalf("Loading config file `%s` failed: %s\n", *configFilePath, err.Error())
		os.Exit(1)
	}
	// Notice that Config is read only
	Config = kv
	log.Errorf("Config file `%s` loaded, config info: %v\n", *configFilePath, Config);
	log.Flush()
}
