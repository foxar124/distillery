package config

type Provider struct {
	Provider string `yaml:"provider" toml:"provider"`
	BaseURL  string `yaml:"base_url" toml:"base_url"`
}
