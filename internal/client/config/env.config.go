package config

// EnvConfig holds the active application environment name (e.g. local, production).
type EnvConfig struct {
	value string
}

func (e *EnvConfig) Get() string {
	return e.value
}

func (e *EnvConfig) Set(value string) {
	e.value = value
}
