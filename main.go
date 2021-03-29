package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/masuldev/example-golang/controller"
)

func main() {
	godotenv.Load("./.env")

	http.HandleFunc("/set", controller.SetRedis)
	http.HandleFunc("/get", controller.GetRedis)

	http.ListenAndServe(":8000", nil)
}
