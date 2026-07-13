package config

import "time"

type LibraryConfig struct {
	Type      string    `yaml:"type" mapstructure:"type"`
	Path      string    `yaml:"path" mapstructure:"path"`
	CreatedAt time.Time `yaml:"createdAt" mapstructure:"createdAt"`
}
