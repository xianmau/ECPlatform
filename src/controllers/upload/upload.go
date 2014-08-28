package upload

import (
	"io"
	"models"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"utils/authority"
	"utils/global"
	log "utils/logger"
)

func UploadifyWebmaster(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	r.ParseMultipartForm(64 << 20) // 64M

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /updateprocess/webmaster/uploadify")
		w.Write([]byte(`{"msg":"未登录","status":"-1","url":""}`))
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /updateprocess/webmaster/uploadify")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "上传文件"); !ok {
			w.Write([]byte(`{"msg":"` + msg + `","status":"-1","url":""}`))
			return
		}

		uploadDir := "uploads/" + r.FormValue("dir") // 上传路径
		file, handler, err := r.FormFile("upload")    // 默认文件框名为upload

		if err != nil {
			log.Error(err.Error())
			w.Write([]byte(`{"msg":"获取文件失败","status":"-1","url":""}`))
			return
		}
		extension := strings.ToLower(filepath.Ext(handler.Filename))
		if !check(extension) {
			w.Write([]byte(`{"msg":"不允许的文件类型","status":"-1","url":""}`))
			return
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + extension
		fileDir, _ := filepath.Abs(uploadDir)

		// 判断目录是否存在，若不存在则生成目录
		fileStat, err := os.Stat(fileDir)
		if err != nil {
			if !os.IsExist(err) {
				err := os.MkdirAll(fileDir, os.ModePerm)
				if err != nil {
					log.Error(err.Error())
					w.Write([]byte(`{"msg":"创建目录失败","status":"-1","url":""}`))
					return
				}
			}
		} else {
			if !fileStat.IsDir() {
				err := os.MkdirAll(fileDir, os.ModePerm)
				if err != nil {
					log.Error(err.Error())
					w.Write([]byte(`{"msg":"创建目录失败","status":"-1","url":""}`))
					return
				}
			}
		}

		f, _ := os.OpenFile(path.Join(fileDir, fileName), os.O_CREATE|os.O_WRONLY, 0666)
		defer f.Close()
		_, err = io.Copy(f, file)
		log.Info("save file: " + path.Join(fileDir, fileName))
		if err != nil {
			log.Error(err.Error())
			w.Write([]byte(`{"msg":"上传失败","status":"-1","url":""}`))
			return
		}

		w.Write([]byte(`{"msg":"上传成功","status":"0","url":"` + uploadDir + fileName + `"}`))
	}
}

func CKEditorWebmaster(w http.ResponseWriter, r *http.Request) {
	// prepare session
	session := global.Sessions.Prepare(w, r)
	// get client ip
	client_ip := string([]byte(r.RemoteAddr)[0:strings.LastIndex(r.RemoteAddr, ":")])

	r.ParseMultipartForm(64 << 20) // 64M

	var admin models.Admin
	if session.Get("admin") != nil {
		admin = (session.Get("admin")).(models.Admin)
	} else {
		log.Info(client_ip + " access /updateprocess/webmaster/ckeditor")
		w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(1, \"\", \"未登录\");</script></body></html>"))
		return
	}

	if r.Method == "POST" {
		// deal with post method
		log.Info(client_ip + " post /updateprocess/webmaster/ckeditor")

		// check authorities
		if ok, msg := authority.Check(admin.Role, "上传文件"); !ok {
			w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(1, \"\", \"" + msg + "\");</script></body></html>"))
			return
		}

		ckeditorFuncNum := r.FormValue("CKEditorFuncNum") // 客户端脚本编号，默认为1
		uploadDir := "uploads/" + r.FormValue("dir")     // 上传路径
		file, handler, err := r.FormFile("upload")        // 默认文件框名为upload

		if err != nil {
			log.Error(err.Error())
			w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"\", \"获取文件失败\");</script></body></html>"))
			return
		}
		extension := strings.ToLower(filepath.Ext(handler.Filename))
		if !check(extension) {
			w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"\", \"不允许的文件类型\");</script></body></html>"))
			return
		}
		fileName := strconv.FormatInt(time.Now().Unix(), 10) + extension
		fileDir, _ := filepath.Abs(uploadDir)

		// 判断目录是否存在，若不存在则生成目录
		fileStat, err := os.Stat(fileDir)
		if err != nil {
			if !os.IsExist(err) {
				err := os.MkdirAll(fileDir, os.ModePerm)
				if err != nil {
					log.Error(err.Error())
					w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"\", \"创建目录失败\");</script></body></html>"))
					return
				}
			}
		} else {
			if !fileStat.IsDir() {
				err := os.MkdirAll(fileDir, os.ModePerm)
				if err != nil {
					log.Error(err.Error())
					w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"\", \"创建目录失败\");</script></body></html>"))
					return
				}
			}
		}

		f, _ := os.OpenFile(path.Join(fileDir, fileName), os.O_CREATE|os.O_WRONLY, 0666)
		defer f.Close()
		_, err = io.Copy(f, file)
		log.Info("save file: " + path.Join(fileDir, fileName))
		if err != nil {
			log.Error(err.Error())
			w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"\", \"上传失败\");</script></body></html>"))
			return
		}

		w.Write([]byte("<html><body><script>window.parent.CKEDITOR.tools.callFunction(" + ckeditorFuncNum + ", \"" + uploadDir + fileName + "\", \"\");</script></body></html>"))
	}
}

func check(extension string) bool {
	availableExtension := []string{".jpg", ".jpeg", ".png", ".gif", ".doc", ".docx", ".xsl", ".xslx", ".ppt", ".pptx", ".txt", ".pdf"}
	for _, v := range availableExtension {
		if v == extension {
			return true
		}
	}
	return false
}
