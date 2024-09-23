package discord

import (
	"log"
	"slot-bot/internal/pkg/config"
	"slot-bot/internal/pkg/slot"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type SlotHandler struct {
	Slots []*slot.Slot
}

func NewSlotHandler(appConfig config.AppConfig) *SlotHandler {
	slots := make([]*slot.Slot, 0, len(appConfig.Slots))
	for _, s := range appConfig.Slots {
		slots = append(slots, s.Slot())
	}
	return &SlotHandler{
		Slots: slots,
	}
}

func (h *SlotHandler) SpinCommand() *discordgo.ApplicationCommand {
	choices := make([]*discordgo.ApplicationCommandOptionChoice, 0, len(h.Slots))
	for i, slot := range h.Slots {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  slot.Name,
			Value: i,
		})
	}

	command := discordgo.ApplicationCommand{
		Name:        "spin",
		Description: "Spin a slot",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "name",
				Description: "Slot name",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Choices:     choices,
				Required:    true,
			},
			{
				Name:        "spins",
				Description: "Number of spins",
				Type:        discordgo.ApplicationCommandOptionInteger,
				Required:    false,
			},
		},
	}

	return &command
}

func (h *SlotHandler) HandleSpinCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	defer func() {
		if rec := recover(); rec != nil {
			log.Printf("[Command %s] Error in HandleSpinCommand: %v", i.ID, rec)

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Something weng wrong :(",
					Flags:   1 << 6, // ephemeral
				},
			})
			if err != nil {
				log.Printf("[Command %s] Failed to respond interaction: %s", i.ID, err)
				return
			}
		}
	}()

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := i.ApplicationCommandData()
	if data.Name != "spin" {
		return
	}

	options := parseOptions(data.Options)
	index := options["name"].IntValue()

	spins := 1
	if options["spins"] != nil {
		spins = int(options["spins"].IntValue())
	}

	var builder strings.Builder
	for i := 0; i < spins; i++ {
		result := h.Slots[index].Spin()
		builder.WriteString(result)
		builder.WriteString("\n")
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: builder.String(),
		},
	})
	if err != nil {
		log.Printf("[Command %s] Failed to respond interaction: %s", i.ID, err)
		return
	}
}
