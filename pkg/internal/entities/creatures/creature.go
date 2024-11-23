package creatures

import (
	"Simulation/pkg/internal/entities"
)

type Creature struct {
	entity.Entity
	Speed, Health, Satiety int
}
