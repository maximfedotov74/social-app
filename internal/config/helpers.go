package config

import (
	"fmt"
	"os"
	"strconv"
)

func createPanicMessage(key string) string {
	return fmt.Sprintf("Key \"%s\" does not exist inside the env file", key)
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		err := createPanicMessage(key)
		panic(err)
	}

	return value
}

func getEnvAsInt(name string) int {
	valueStr := getEnv(name)
	value, err := strconv.Atoi(valueStr)

	if err != nil {
		panic("Error when coverting env variable to integer")
	}

	return value
}

func getEnvAsBool(name string) bool {
	valueStr := getEnv(name)
	value, err := strconv.ParseBool(valueStr)

	if err != nil {
		panic("Error when coverting env variable to boolean")
	}

	return value
}
