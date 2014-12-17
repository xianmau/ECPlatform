package web

import (
	"encoding/json"
	"models"
	"net/http"
	"strings"
	"controllers/common"
)

func Home(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	// special check for home page
	checkUrl := strings.ToLower(r.URL.Path)
	if checkUrl != "/" && checkUrl != "/index.html" && checkUrl != "/web" && checkUrl != "/web/home" {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("home.html", []string{"views/web/home.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
		if err != nil {
			return
		}
		data := make(map[string]interface {})
		// check login
		isUserLogin := false
		var user models.User
		if session.Get("user") != nil {
			user = (session.Get("user")).(models.User)
			isUserLogin = true
		} else {
			isUserLogin = false
		}
		// set login info
		data["User"] = user
		data["IsUserLogin"] = isUserLogin

		webInfoes, err := models.GetWebHomeInfoList()
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if len(webInfoes) == 0 {
			defaultHandler.LogError("floor info can't find")
			return
		}

		F4 := []*models.Goods{}

		var F4GoodsId []string
		err = json.Unmarshal([]byte(webInfoes[0].Floor4), &F4GoodsId)

		for _, gid := range F4GoodsId {
			goods, err := models.GetGoods(gid)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
			if goods != nil {
				F4 = append(F4, goods)
			} else {
				goods = &models.Goods{
					Id:     0,
					Title:  "--",
					Price:  0.00,
					Images: `{"small":"/statics/img/web/notfound_small.png","large":"/statics/img/web/notfound_large.png"}`,
				}
				F4 = append(F4, goods)
			}
		}

		data["F4_UP"] = F4[0]
		data["F4_UPs"] = F4[1:4]
		data["F4_DOWN"] = F4[4]
		data["F4_DOWNs"] = F4[5:8]

		F4_SubCol, _ := models.GetSubGoodsCategoryList("华农出品")
		if len(F4_SubCol) > 6 {
			F4_SubCol = F4_SubCol[0:6]
		}
		data["F4_SubCol"] = F4_SubCol

		F4_HotGoods, _ := models.GetHotGoodsListByCategory("华农出品")
		if len(F4_HotGoods) > 7 {
			F4_HotGoods = F4_HotGoods[0:7]
		}
		data["F4_HotGoods"] = F4_HotGoods


		// 故事
		AL1, err := models.GetArticleListByCategory("一品故事")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if len(AL1) > 12{
			AL1 = AL1[0:12]
		}
		data["AL1"] = AL1

		// 菜谱
		AL2, err := models.GetArticleListByCategory("一品菜谱")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if len(AL2) > 5{
			AL2 = AL2[0:5]
		}
		data["AL2"] = AL2

		// 养生
		AL3, err := models.GetArticleListByCategory("一品养生")
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if len(AL3) > 5{
			AL3 = AL3[0:5]
		}
		data["AL3"] = AL3

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}
