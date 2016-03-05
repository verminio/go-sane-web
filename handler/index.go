package handler

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		indexGET(w, r)
		return
	default:
		http.Error(w, "", http.StatusNotImplemented)
		return
	}
}

func indexGET(w http.ResponseWriter, r *http.Request) {
	model := struct {
		Uptime string
	}{
		"uptime",
	}

	templates["index"].ExecuteTemplate(w, "layout", model)
}
