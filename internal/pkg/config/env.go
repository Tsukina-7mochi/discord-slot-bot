package config

import (
	"errors"
	"os"
	"strings"
)

type EnvConfig struct {
	Token          string
	AppID          string
	GuildIDs       []string
	AppConfigFile  string
	DeleteCommands bool
}

func loadEnv(name string) (string, error) {
	token := os.Getenv(name)
	if len(token) == 0 {
		return "", errors.New("Environment variable \"" + name + "\" is not set.")
	}

	return token, nil
}

func EnvConfigFromEnv() (*EnvConfig, error) {
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

	// igreno error because this environment variable is optional
	deleteCommandsStr, _ := loadEnv("DELETE_COMMANDS")
	deleteCommands := deleteCommandsStr != ""

	appConfigName, err := loadEnv("APP_CONFIG_FILE")

	return &EnvConfig{
		Token:          token,
		AppID:          appID,
		GuildIDs:       guildIDs,
		AppConfigFile:  appConfigName,
		DeleteCommands: deleteCommands,
	}, nil
}
