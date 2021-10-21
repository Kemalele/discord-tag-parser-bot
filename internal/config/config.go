package config

import "github.com/kelseyhightower/envconfig"

const (
	// ServiceName contains a service name prefix which used in ENV variables.
	ServiceName = "BOT"
)

// Config contains ENV variables.
type Config struct {
	Token string `split_words:"true" required:"true"`
	// BotName  string `split_words:"true" required:"true"`
}

// Load settles ENV variables into Config structure.
func (c *Config) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
