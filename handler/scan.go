package handler

import (
	"github.com/tjgq/sane"
	"net/http"
)

func init() {
	err := sane.Init()
	if err != nil {
		panic(err)
	}
}

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	templates["scan"].ExecuteTemplate(w, "layout", "")
}

func scanGET(w http.ResponseWriter, r *http.Request) {
	devices, err := sane.Devices()

	if err != nil {
		panic(err)
	}

	model := struct {
		Devices []sane.Device
	}{
		devices,
	}

	templates["scan"].ExecuteTemplate(w, "layout", model)
}
