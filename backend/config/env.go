package config

import "os"

//GetEnvWithDefault reads env by specified name or returns default value if env is not found.
func GetEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
