package handlers

import (
	"net/http"
	"path"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("./static/*")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}
