package service

import (
	"errors"

	"github.com/masuldev/whisper-router/config"
	"github.com/masuldev/whisper-router/domain"
)

var (
	ErrNotFoundValue = errors.New("Not Found Value")
	ErrInvalidSyntax = errors.New("Invalid Syntax")
)

func GetRedis(credential string) (string, error) {
	redis := config.ConnectionRedis()

	result, err := redis.Get(credential).Result()
	if err != nil {
		return "", ErrNotFoundValue
	}

	return result, nil
}

func SetRedis(credential domain.RedisCredential) error {
	redis := config.ConnectionRedis()

	result, err := redis.Set(credential.FirstCredential, credential.SecondCredential, 0).Result()
	if err != nil && result != "OK" {
		return ErrInvalidSyntax
	}

	return nil
}
