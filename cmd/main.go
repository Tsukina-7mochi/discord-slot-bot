package main

import (
	"log"
	"os"
	"os/signal"
	"slot-bot/internal/pkg/config"
	"slot-bot/internal/pkg/discord"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	envConfig, err := config.EnvConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	appConfig, err := config.ReadAppConfig(envConfig.AppConfigFile)
	if err != nil {
		log.Fatalf("Failed to read app config: %v", err)
	}

	session, err := discordgo.New("Bot " + envConfig.Token)
	if err != nil {
		log.Fatalf("Failed to start discord bot: %v", err)
	}

	slotHandler := discord.NewSlotHandler(*appConfig)

	session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type != discordgo.InteractionApplicationCommand {
			return
		}

		data := i.ApplicationCommandData()

		if data.Name == "spin" {
			slotHandler.HandleSpinCommand(s, i)
		}
	})

	session.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %s", r.User.String())
	})

	// delete all commands
	if envConfig.DeleteCommands {
		for _, guildID := range envConfig.GuildIDs {
			log.Printf("Deleting commands for guild %s", guildID)

			commands, err := session.ApplicationCommands(envConfig.AppID, guildID)
			if err != nil {
				log.Printf("Failed to retrieve commands for guild %s: %v", guildID, err)
				continue
			}
			for _, command := range commands {
				err = session.ApplicationCommandDelete(envConfig.AppID, guildID, command.ID)
				if err != nil {
					log.Printf("Failed to delete command %s for guild %s: %v", command.Name, guildID, err)
				}
			}
		}
	}

	// register commands
	spinCommand := slotHandler.SpinCommand()
	commands := []*discordgo.ApplicationCommand{spinCommand}
	for _, guildID := range envConfig.GuildIDs {
		log.Printf("Registering commands for guild %s", guildID)

		_, err = session.ApplicationCommandBulkOverwrite(envConfig.AppID, guildID, commands)
		if err != nil {
			log.Fatalf("Failed to register commands for guild %s: %v", guildID, err)
		}
	}

	err = session.Open()
	if err != nil {
		log.Fatalf("Failed to open session: %v", err)
	}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	<-sigch

	log.Println("Closing session...")

	err = session.Close()
	if err != nil {
		log.Printf("Failed to close session gracefully: %v", err)
	}
}
