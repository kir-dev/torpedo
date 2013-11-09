package util

import (
	"os"
)

const (
	ENV  = "ENV"
	DEV  = "development"
	TEST = "test"
	PROD = "production"
)

// Determines whether the the application runs in a development environment. The
// application's running environment is based on the ENV environment variable.
// Development is the fallback option as well.
func IsDev() bool {
	return os.Getenv(ENV) == "" || os.Getenv(ENV) == DEV
}

// Determines whether the the application runs in a development environment. See
// details at the IsDev() method.
func IsTest() bool {
	return os.Getenv(ENV) == TEST
}
