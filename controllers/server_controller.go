package controllers

import (
	"net/http"

	"github.com/dsukesato/layerArc/controllers/handler"
)

func Serve() {
	mux := http.NewServeMux() //マルチプレクサを用いてhandlerを管理

	mux.HandleFunc("/index", handler.BookHandler{}.BookIndex)

	http.ListenAndServe(":8080", mux)
}