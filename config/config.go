package config

type (
	// Config stores the configuration settings.
	Config struct {
		HTTP struct {
			Host string
			Port string `envconfig:"PORT" default:"8080"`
			Root string `default:"/"`
		}
		Facebook struct {
			AppID   string `envconfig:"TEST_FACEBOOK_APP_ID"`
			Secret  string
			Version string
		}
	}
)
