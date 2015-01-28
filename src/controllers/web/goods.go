package web

import (
	"models"
	"net/http"
	log "utils/glog"
	"controllers/common"
)

func GoodsCat(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("gcat.html", []string{"views/web/gcat.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
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

		// current category
		goodsCategoryList, err := models.GetGoodsCategoryList()
		if err != nil {
			log.Error(err.Error())
			return
		}
		// left category list
		goodsCategories := []*models.SliceWithName{}
		for _, v := range goodsCategoryList{
			if v.Parent == ""{
				subs, err := models.GetSubGoodsCategoryList(v.Name)
				if err != nil{
					defaultHandler.LogError(err)
					return
				}
				if len(subs) > 0{
					item := new(models.SliceWithName)
					item.Name = v.Name
					item.Slice = []interface{}{}
					for _, v2 := range subs{
						item.Slice = append(item.Slice, v2.Name)
					}
					goodsCategories = append(goodsCategories, item)
				}
			}
		}

		data["goodsCategories"] = goodsCategories

		// current category
		r.ParseForm()
		get_c := r.Form.Get("c")
		for _, v := range goodsCategoryList {
			if v.Name == get_c {
				data["currentCategory"] = v
				break
			}
		}

		// get all goods of the category
		goodsList, err := models.GetGoodsListByCategory(get_c)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		data["goodsList"] = goodsList
		cur_c, err := models.GetGoodsCategory(get_c)
		goodsCategoryInfo, err := models.GetGoodsCategoryInfo(cur_c.Name)
		if err != nil {
			defaultHandler.LogError(err)
			return
		}
		if goodsCategoryInfo == nil {
			goodsCategoryInfo, err = models.GetGoodsCategoryInfo(cur_c.Parent)
			if err != nil {
				defaultHandler.LogError(err)
				return
			}
		}
		if goodsCategoryInfo == nil {
			goodsCategoryInfo = new(models.GoodsCategoryInfo)
		}
		data["goodsCategoryInfo"] = goodsCategoryInfo

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}

func GoodsDetail(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	session := defaultHandler.Prepare()

	if r.Method == "GET" {
		// render template
		err := defaultHandler.RenderTemplate("goods.html", []string{"views/web/goods.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html"})
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

		// current goods
		r.ParseForm()
		get_id := r.Form.Get("id")
		goods, err := models.GetGoods(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if goods != nil {
			data["goods"] = goods
			// current category
			r.ParseForm()
			currentCategory, err := models.GetGoodsCategory(goods.Category)
			if err != nil {
				log.Error(err.Error())
				return
			}
			data["currentCategory"] = currentCategory
		} else {
			data["NotFound"] = "找不到该商品的信息"
		}
		comments, err := models.GetCommentListByGoods(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["comments"] = comments

		recommends := []models.Goods{}
		data["recommends"] = recommends

		// add goods click
		goods_statistic, err := models.GetGoodsStatistic(get_id)
		if err != nil {
			log.Error(err.Error())
			return
		}
		if goods_statistic == nil {
			models.CreateGoodsStatistic(get_id)
		}
		models.IncreaseGoodsStatisticViewTimes(get_id)

		// execute template
		defaultHandler.ExecuteTemplate(data)
	}
}
