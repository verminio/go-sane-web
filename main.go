package main

import (
	"github.com/verminio/go-sane-web/handler"
	"net/http"
)

func init() {
}

func main() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/scan", handler.ScanHandler)
	http.HandleFunc("/settings", handler.SettingsHandler)
	http.ListenAndServe(":8080", nil)
}
