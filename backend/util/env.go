package util

import (
	"fmt"
	"os"
)

// Getenv returns an environment variable or panics if it's not found.
func Getenv(key string) (result string) {
	var found bool
	if result, found = os.LookupEnv(key); !found {
		panic(fmt.Sprintf("environment variable %s not found", key))
	}
	return
}
