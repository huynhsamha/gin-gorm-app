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
