package config

import (
	"errors"
	"os"
)

func loadEnv(name string) (string, error) {
	token := os.Getenv(name)
	if len(token) == 0 {
		return "", errors.New("Environment variable \"" + name + "\" is not set.")
	}

	return token, nil
}

func FromEnv() (*Config, error) {
	token, err := loadEnv("TOKEN")
	if err != nil {
		return nil, err
	}

	appID, err := loadEnv("APP_ID")
	if err != nil {
		return nil, err
	}

	guildID, err := loadEnv("GUILD_ID")
	if err != nil {
		return nil, err
	}

	return &Config{
		Token:   token,
		AppID:   appID,
		GuildID: guildID,
	}, nil
}
