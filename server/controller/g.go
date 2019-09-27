package controller

import (
	"html/template"
)

var (
	homeController home
	taskController task
	templates      map[string]*template.Template
)

//初始化模板
func init() {
	templates = PopulateTemplates()
}

//开启注册路由
func Startup() {
	homeController.registerRoutes()
	taskController.registerRoutes()
}
