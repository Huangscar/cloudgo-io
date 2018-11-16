package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

func LoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("method:", r.Method) //获取请求的方法
		if r.Method == "GET" {
			t, _ := template.ParseFiles("templates/login.html")
			log.Println(t.Execute(w, nil))
		} else {
			r.ParseForm()
			//请求的是登录数据，那么执行登录的逻辑判断

			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			//fmt.Fprintf(w, "username:"+r.Form["username"][0]+"\n")
			//fmt.Fprintf(w, "password:"+r.Form["password"][0])
			formatter := render.New(render.Options{
				Directory:  "templates",
				Extensions: []string{".html"},
				IndentJSON: true,
			})
			formatter.HTML(w, http.StatusOK, "index", struct {
				USER     string `json:"user"`
				PASSWORD string `json:"password"`
			}{USER: r.Form["username"][0], PASSWORD: r.Form["password"][0]})

		}
	}
}
