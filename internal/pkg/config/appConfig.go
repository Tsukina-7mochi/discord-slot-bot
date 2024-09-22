package config

import (
	"encoding/json"
	"os"
	"slot-bot/internal/pkg/slot"
)

type Slot struct {
	Name  string     `json:"name"`
	Reels [][]string `json:"reels"`
}

func (s *Slot) Slot() *slot.Slot {
	return &slot.Slot{
		Name:  s.Name,
		Reels: s.Reels,
	}
}

type AppConfig struct {
	Slots []Slot `json:"slots"`
}

func ReadAppConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg AppConfig
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *AppConfig) GetSlots() *[]slot.Slot {
	slots := make([]slot.Slot, 0, len(cfg.Slots))
	for _, s := range cfg.Slots {
		slots = append(slots, *s.Slot())
	}

	return &slots
}
