package slot

import (
	"math/rand/v2"
	"strings"
)

type Slot struct {
	Name  string
	Reels [][]string
}

func (s *Slot) Spin() string {
	builder := new(strings.Builder)

	for _, reel := range s.Reels {
		builder.WriteString(reel[rand.N(len(reel))])
	}

	return builder.String()
}
