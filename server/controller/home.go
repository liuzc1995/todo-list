package controller

import (
	"net/http"

	"github.com/542213314/frame-sample/server/vm"

	"github.com/gorilla/mux"
)

type home struct{}

//设置路由
func (p home) registerRoutes() {
	r := mux.NewRouter()
	//首页路由
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/add", taskAddHandler)
	r.HandleFunc("/update", taskUpdateHandler)
	r.HandleFunc("/delete", taskDeleteHandler)
	//静态资源配置
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", r)
}

//tpPath模板路径,相对template目录文件路径
//GetVM() 取模板页面内容
//templates[tpName].Execute 执行渲染模板

//首页
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpPath := "content/index.html"
	vop := vm.IndexViewModelOp{}
	if r.Method == "GET" {
		v := vop.GetVM()
		templates[tpPath].Execute(w, &v)
	}
}
