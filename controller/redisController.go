package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/masuldev/whisper-router/domain"
	"github.com/masuldev/whisper-router/dto"
	"github.com/masuldev/whisper-router/service"
)

var (
	ErrRedisConnection  = errors.New("Connection to redis is unable")
	ErrInvalidParameter = errors.New("Request parameter is invalid")
)

func GetRedis(w http.ResponseWriter, r *http.Request) {
	var credential dto.RedisDto

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Requset: ", ErrInvalidParameter)
		return
	}

	result, err := service.GetRedis(credential.Parameter)

	if err != nil {
		fmt.Fprint(w, ErrRedisConnection)
		return
	}
	fmt.Fprint(w, result)
}

func SetRedis(w http.ResponseWriter, r *http.Request) {
	var credential domain.RedisDomain

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Requset: ", ErrInvalidParameter)
		return
	}

	err = service.SetRedis(credential)
	if err != nil {
		fmt.Fprint(w, "Server: ", ErrRedisConnection)
		return
	}

	w.Header().Add("content-type", "application/json")

	data, _ := json.Marshal(credential)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}
