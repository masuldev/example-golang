package main

import (
	"net/http"

	"github.com/masuldev/whisper-router/controller"
)

func main() {
	http.HandleFunc("/set", controller.SetRedis)
	http.HandleFunc("/get", controller.GetRedis)

	http.ListenAndServe(":8000", nil)
}
