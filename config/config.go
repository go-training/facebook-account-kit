package config

type (
	// Config stores the configuration settings.
	Config struct {
		Facebook struct {
			AppID   string `envconfig:"TEST_FACEBOOK_APP_ID"`
			Secret  string
			Version string
		}
	}
)
