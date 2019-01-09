package utils

import "os"

// DefaultGetEnv : get environment variables with default value
func DefaultGetEnv(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultVal
	}
	return val
}

// RequireGetEnv : get environment variable with requirement
func RequireGetEnv(key string) (val string, exist bool) {
	val, exist = os.Getenv(key), true
	if val == "" {
		exist = false
	}
	return
}

// DefaultStringEmpty : get default string if value is empty
func DefaultStringEmpty(value string, defaultVal string) string {
	if value == "" {
		value = defaultVal
	}
	return value
}
