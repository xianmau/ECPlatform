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

		F1 := []*models.Goods{}
		F2 := []*models.Goods{}
		F3 := []*models.Goods{}
		F4 := []*models.Goods{}

		var F1GoodsId []string
		err = json.Unmarshal([]byte(webInfoes[0].Floor1), &F1GoodsId)
		var F2GoodsId []string
		err = json.Unmarshal([]byte(webInfoes[0].Floor2), &F2GoodsId)
		var F3GoodsId []string
		err = json.Unmarshal([]byte(webInfoes[0].Floor3), &F3GoodsId)
		var F4GoodsId []string
		err = json.Unmarshal([]byte(webInfoes[0].Floor4), &F4GoodsId)

		for _, gid := range F1GoodsId {
			goods, err := models.GetGoods(gid)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
			if goods != nil {
				F1 = append(F1, goods)
			} else {
				goods = &models.Goods{
					Id:     0,
					Title:  "--",
					Price:  0.00,
					Images: `{"small":"/statics/img/web/notfound_small.png","large":"/statics/img/web/notfound_large.png"}`,
				}
				F1 = append(F1, goods)
			}
		}

		for _, gid := range F2GoodsId {
			goods, err := models.GetGoods(gid)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
			if goods != nil {
				F2 = append(F2, goods)
			} else {
				goods = &models.Goods{
					Id:     0,
					Title:  "--",
					Price:  0.00,
					Images: `{"small":"/statics/img/web/notfound_small.png","large":"/statics/img/web/notfound_large.png"}`,
				}
				F2 = append(F2, goods)
			}
		}

		for _, gid := range F3GoodsId {
			goods, err := models.GetGoods(gid)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
			if goods != nil {
				F3 = append(F3, goods)
			} else {
				goods = &models.Goods{
					Id:     0,
					Title:  "--",
					Price:  0.00,
					Images: `{"small":"/statics/img/web/notfound_small.png","large":"/statics/img/web/notfound_large.png"}`,
				}
				F3 = append(F3, goods)
			}
		}

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

		data["F1_UP"] = F1[0]
		data["F1_UPs"] = F1[1:4]
		data["F1_DOWN"] = F1[4]
		data["F1_DOWNs"] = F1[5:8]

		data["F2_UP"] = F2[0]
		data["F2_UPs"] = F2[1:4]
		data["F2_DOWN"] = F2[4]
		data["F2_DOWNs"] = F2[5:8]

		data["F3_UP"] = F3[0]
		data["F3_UPs"] = F3[1:4]
		data["F3_DOWN"] = F3[4]
		data["F3_DOWNs"] = F3[5:8]

		data["F4_UP"] = F4[0]
		data["F4_UPs"] = F4[1:4]
		data["F4_DOWN"] = F4[4]
		data["F4_DOWNs"] = F4[5:8]

		F1_SubCol, _ := models.GetSubGoodsCategoryList("优质粮油")
		F2_SubCol, _ := models.GetSubGoodsCategoryList("农家干货")
		F3_SubCol, _ := models.GetSubGoodsCategoryList("新鲜果蔬")
		F4_SubCol, _ := models.GetSubGoodsCategoryList("华农出品")
		if len(F1_SubCol) > 6 {
			F1_SubCol = F1_SubCol[0:6]
		}
		if len(F2_SubCol) > 6 {
			F2_SubCol = F2_SubCol[0:6]
		}
		if len(F3_SubCol) > 6 {
			F3_SubCol = F3_SubCol[0:6]
		}
		if len(F4_SubCol) > 6 {
			F4_SubCol = F4_SubCol[0:6]
		}
		data["F1_SubCol"] = F1_SubCol
		data["F2_SubCol"] = F2_SubCol
		data["F3_SubCol"] = F3_SubCol
		data["F4_SubCol"] = F4_SubCol

		F1_HotGoods, _ := models.GetHotGoodsListByCategory("优质粮油")
		F2_HotGoods, _ := models.GetHotGoodsListByCategory("农家干货")
		F3_HotGoods, _ := models.GetHotGoodsListByCategory("新鲜果蔬")
		F4_HotGoods, _ := models.GetHotGoodsListByCategory("华农出品")
		if len(F1_HotGoods) > 7 {
			F1_HotGoods = F1_HotGoods[0:7]
		}
		if len(F2_HotGoods) > 7 {
			F2_HotGoods = F2_HotGoods[0:7]
		}
		if len(F3_HotGoods) > 7 {
			F3_HotGoods = F3_HotGoods[0:7]
		}
		if len(F4_HotGoods) > 7 {
			F4_HotGoods = F4_HotGoods[0:7]
		}
		data["F1_HotGoods"] = F1_HotGoods
		data["F2_HotGoods"] = F2_HotGoods
		data["F3_HotGoods"] = F3_HotGoods
		data["F4_HotGoods"] = F4_HotGoods

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}
