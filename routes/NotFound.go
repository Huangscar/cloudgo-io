package routes

import (
	"net/http"
)

func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "501 NOT Implemented", 501)

	}
}
