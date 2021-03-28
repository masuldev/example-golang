package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/masuldev/whisper-router/config"
	"github.com/masuldev/whisper-router/domain"
)

var (
	ErrRedisConnection  = errors.New("Connection to redis is unable")
	ErrInvalidParameter = errors.New("Request parameter is invalid")
)

func GetRedis(w http.ResponseWriter, r *http.Request) {
	var credential domain.RedisCredential

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Requset: ", ErrInvalidParameter)
		return
	}

	redis := config.ConnectionRedis()

	result, err := redis.Get(credential.FirstCredential).Result()
	if err != nil {
		fmt.Fprint(w, ErrRedisConnection)
		return
	}
	fmt.Fprint(w, result)
}

func SetRedis(w http.ResponseWriter, r *http.Request) {
	var credential domain.RedisCredential

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Requset: ", ErrInvalidParameter)
		return
	}

	redis := config.ConnectionRedis()

	result, err := redis.Set(credential.FirstCredential, credential.SecondCredential, 0).Result()
	if err != nil {
		fmt.Fprint(w, "Server: ", ErrRedisConnection)
		return
	}
	fmt.Println(result)

	w.Header().Add("content-type", "application/json")

	data, _ := json.Marshal(credential)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}
