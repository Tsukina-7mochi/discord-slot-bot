package config

import (
	"errors"
	"os"
	"strings"
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

	guildIDStr, err := loadEnv("GUILD_IDS")
	if err != nil {
		return nil, err
	}
	guildIDs := strings.Split(guildIDStr, ",")

	appConfigName, err := loadEnv("APP_CONFIG_FILE")

	return &Config{
		Token:         token,
		AppID:         appID,
		GuildIDs:      guildIDs,
		AppConfigFile: appConfigName,
	}, nil
}
