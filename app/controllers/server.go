package controllers

import (
	"app/config"
	"net/http"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files)) // StripPrefix=要求されたURLから’/static/’を外してパスを探す

	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
