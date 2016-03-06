package handler

import (
	"github.com/tjgq/sane"
	"net/http"
    "log"
)

func init() {
    log.Println("Initializing SANE bindings")
	err := sane.Init()
	if err != nil {
		panic(err)
	}
    log.Println("Finished initializing SANE bindings")
}

func ScanHandler(w http.ResponseWriter, r *http.Request) {
  switch {
    case r.Method == "GET":
        scanGET(w, r)
        return
    default:
        http.Error(w, "", http.StatusNotImplemented)
        return
    }
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
