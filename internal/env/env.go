package env

import (
	"errors"
	"os"
	"strconv"
)

func GetString(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(errors.New("env value not found"))
	}

	return value
}

func GetInt(key string) int {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(errors.New("env value not found"))
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func GetBool(key string) bool {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(errors.New("env value not found"))
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}
