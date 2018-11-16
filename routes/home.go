package routes

import (
	"net/http"

	"github.com/unrolled/render"
)

func HomeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			STATICP string `json:"staticPJ"`
			JSP     string `json:"jsPJ"`
			TABLEP  string `json:"tablePJ"`
		}{STATICP: "/", JSP: "/home", TABLEP: "/login"})
	}

}
