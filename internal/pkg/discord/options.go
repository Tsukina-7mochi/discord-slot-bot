package discord

import "github.com/bwmarrin/discordgo"

type optionMap = map[string]*discordgo.ApplicationCommandInteractionDataOption

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) optionMap {
	om := make(optionMap)
	for _, opt := range options {
		om[opt.Name] = opt
	}
	return om
}
