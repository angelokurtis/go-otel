package _test

import "os"

// EnvironmentVariables is designed for managing environment variables in tests and ensuring a clean testing environment.
type EnvironmentVariables struct{ keys []string }

// SetEnvironmentVariables creates a new EnvironmentVariables struct and sets the environment variables in the specified map.
func SetEnvironmentVariables(kv map[string]string) *EnvironmentVariables {
	envVars := new(EnvironmentVariables)
	for k, v := range kv {
		envVars.Set(k, v)
	}

	return envVars
}

// Set sets the environment variable with the specified key and value.
func (e *EnvironmentVariables) Set(k, v string) {
	_ = os.Setenv(k, v)
	e.keys = dedupe(e.keys, k)
}

// Unset unsets all the environment variables.
func (e *EnvironmentVariables) Unset() {
	for _, v := range e.keys {
		_ = os.Unsetenv(v)
	}
}

func dedupe(slice []string, s string) []string {
	for _, existing := range slice {
		if existing == s {
			return slice // string already exists in slice
		}
	}

	return append(slice, s) // string doesn't exist, append it
}
