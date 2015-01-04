package web

import (
	"models"
	"net/http"
	"controllers/common"
)


func ProOrderCreate(w http.ResponseWriter, r *http.Request) {
	defaultHandler := common.NewDefaultHandler(w, r)
	_ = defaultHandler.Prepare()

	if r.Method == "POST" {
		form_gid := r.PostFormValue("gid")
		form_name := r.PostFormValue("name")
		form_tel := r.PostFormValue("tel")
		form_addr := r.PostFormValue("addr")
		form_num := r.PostFormValue("num")

		err := models.CreatePreorder(form_gid, "0", form_name, form_tel, form_addr, form_num, "0")
		if err != nil {
			defaultHandler.LogError(err)
			w.Write([]byte(""))
			return
		}

		// return data
		w.Write([]byte("success"))
	}
}
