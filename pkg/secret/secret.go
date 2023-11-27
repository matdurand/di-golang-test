package secret

import (
	"fmt"
	"os"
)

type Secret interface {
	Value() (*string, error)
}

type SecretProvider interface {
	GetSecret(name string) Secret
}

type SecretFromEnvProvider struct{}

type secretFromEnv struct {
	envName string
}

func (s secretFromEnv) Value() (*string, error) {
	val := os.Getenv(s.envName)
	if val == "" {
		return nil, fmt.Errorf("env variable %s is not set", s.envName)
	}
	return &val, nil
}

func (s *SecretFromEnvProvider) GetSecret(name string) Secret {
	return secretFromEnv{
		envName: name,
	}
}
