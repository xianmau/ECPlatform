package authority

import (
	"models"
)

func Check(role models.Role, dowhat string, args ...string) (bool, string) {
	if role.Name == "root" {
		return true, ""
	}
	switch dowhat {
	case "登录":
		for _, v := range role.Authority {
			if v == dowhat {
				return true, ""
			}
		}
		return false, "请确保您有登录权限"
	case "文章管理":
		for _, v := range role.Authority {
			if v == dowhat {
				return true, ""
			}
		}
		return false, "请确保您有文章管理权限"
	case "链接管理":
		for _, v := range role.Authority {
			if v == dowhat {
				return true, ""
			}
		}
		return false, "请确保您有链接管理权限"
	case "商品管理":
		for _, v := range role.Authority {
			if v == dowhat {
				return true, ""
			}
		}
		return false, "请确保您有商品管理权限"
	case "上传文件":
		for _, v := range role.Authority {
			if v == dowhat {
				return true, ""
			}
		}
		return false, "请确保您有上传文件权限"
	default:
	}
	return false, "请确保您有相应的权限"
}
