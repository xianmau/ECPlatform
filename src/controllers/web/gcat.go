package web

import (
	"html/template"
	"models"
	"net/http"
	"strings"
	"utils/global"
	log "utils/logger"
	"regexp"
)

func GoodsCat(w http.ResponseWriter, r *http.Request) {
	// prepare session
	_ = global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	if r.Method == "GET" {
		// deal with get method
		log.Info(client_ip + " get /web/gcat")

		// render template
		t := template.New("gcat.html")
		t.Funcs(template.FuncMap{"GetJsonData": GetJsonData})
		t, err := t.ParseFiles("views/web/gcat.html", "views/web/styles.html", "views/web/scripts.html", "views/web/headerpart.html", "views/web/footerpart.html")
		if err != nil {
			log.Error(err.Error())
			return
		}

		// bind data
		data := make(map[string]interface{})

		// 分类列表
		goodsCategoryList, err := models.GetGoodsCategoryList()
		if err != nil {
			log.Error(err.Error())
			return
		}

		// 左侧分类列表
		goodsCategories := []*models.SliceWithName{}
	LOOP_1:
		for _, v := range goodsCategoryList {
			if v.Parent != "" {

				for _, v2 := range goodsCategories {
					if v2.Name == v.Parent {
						v2.Slice = append(v2.Slice, v.Name)
						continue LOOP_1
					}
				}
				item := new(models.SliceWithName)
				item.Name = v.Parent
				item.Slice = []interface{}{v.Name}
				goodsCategories = append(goodsCategories, item)
			}
		}
		data["goodsCategories"] = goodsCategories

		// 当前分类
		r.ParseForm()
		get_c := r.Form.Get("c")
		for _, v := range goodsCategoryList {
			if v.Name == get_c {
				data["currentCategory"] = v
				break
			}
		}

		// 获取当前分类下的所有商品
		goodsList, err := models.GetGoodsListByCategory(get_c)
		if err != nil {
			log.Error(err.Error())
			return
		}
		data["goodsList"] = goodsList
		// execute template
		err = t.Execute(w, data)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
}

func GetJsonData(jsonStr string, key string) string {
	log.Error(jsonStr)
	log.Error(key)
	re, err := regexp.Compile(`"` + key + `":"(.*?)"`)
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	value := re.FindStringSubmatch(jsonStr)
	if len(value) > 1{
		return value[1]
	}
	return ""
}
