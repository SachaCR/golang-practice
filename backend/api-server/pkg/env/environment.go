package env

import "fmt"

type AppEnvironment int64

const (
	Production AppEnvironment = iota
	Staging
	Localhost
	Test
)

func (s AppEnvironment) String() string {
	switch s {

	case Production:
		return "production"

	case Staging:
		return "staging"

	case Localhost:
		return "localhost"

	case Test:
		return "test"
	}

	return "unknown"
}

func ParseEnvFromString(envString string) AppEnvironment {
	switch envString {

	case "production":
		return Production

	case "staging":
		return Staging

	case "localhost":
		return Localhost

	case "test":
		return Test
	}

	panic(fmt.Errorf("UNKNOWN ENVIRONMENT: " + envString))
}
