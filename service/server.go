package service

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/Huangscar/cloudgo-io/routes"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.HandleFunc("/login", routes.LoginHandler(formatter))
	mx.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("templates/home.html")
		log.Println(t.Execute(w, nil))
	})
	mx.HandleFunc("/api/info", routes.HomeHandler(formatter)).Methods("GET")

	mx.HandleFunc("/unknown", routes.NotFoundHandler())
	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

}
