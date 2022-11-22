package ggos

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

// Env get environment variable from key
func Env(key string, out *string) {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		*out = val
	}
}

// MustEnv get environment variable from key, if both value and *out is empty, panic
func MustEnv(key string, out *string) {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		*out = val
	} else {
		if *out == "" {
			panic(errors.New("missing environment variable '" + key + "'"))
		}
	}
}

// BoolEnv get bool environment variable from key
func BoolEnv(key string, out *bool) {
	if val := strings.TrimSpace(os.Getenv(key)); val != "" {
		*out, _ = strconv.ParseBool(val)
	}
}
