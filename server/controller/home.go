package controller

import (
	"encoding/json"
	"net/http"

	"github.com/542213314/frame-sample/server/vm"
	"go.mongodb.org/mongo-driver/bson"
	"golanger.com/log"

	"github.com/gorilla/mux"
)

type home struct{}

//设置路由
func (p home) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", p.indexHandler)
	http.Handle("/", r)

	staticHandler()
}

//静态资源路径配置
func staticHandler() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

//tpPath模板路径,相对template目录文件路径
//GetVM() 取模板页面内容
//templates[tpName].Execute 执行渲染模板

//首页
func (p home) indexHandler(w http.ResponseWriter, r *http.Request) {
	tpPath := "content/index.html"
	vop := vm.IndexViewModelOp{}
	if r.Method == "GET" {
		v := vop.GetVM()
		templates[tpPath].Execute(w, &v)
	}
	if r.Method == "POST" {
		r.ParseForm()
		m := bson.M{
			"status":  0,
			"message": "",
		}
		taskName := r.Form.Get("name")
		if err := vm.AddTask(taskName); err != nil {
			m["message"] = "Error in add task"
			log.Debug("add task error:", err)
		} else {
			m["message"] = "添加成功"
			m["status"] = 1
		}
		ret, _ := json.Marshal(m)
		w.Write(ret)
		return
	}
}
