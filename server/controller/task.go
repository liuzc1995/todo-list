package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type task struct{}

//设置路由
func (p *task) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", p.IndexHandler)
}

func (p *task) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		m := bson.M{
			"status":  0,
			"message": "",
		}
		ret, _ := json.Marshal(m)
		w.Write(ret)
		return
	}
}
