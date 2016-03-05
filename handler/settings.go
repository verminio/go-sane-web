package handler

import (
	"net/http"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	templates["settings"].ExecuteTemplate(w, "layout", "")
}
