package config

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

type Config struct {
}

func Load(env string) (*Config, error) {
	return &Config{}, nil
}
