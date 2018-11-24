package config

import "github.com/kelseyhightower/envconfig"

// Load loads the configuration from the environment.
func Load() (Config, error) {
	config := Config{}
	err := envconfig.Process("test", &config)
	return config, err
}

// MustLoad loads the configuration from the environmnet
// and panics if an error is encountered.
func MustLoad() Config {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	return config
}
