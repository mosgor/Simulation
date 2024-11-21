package creatures

import (
	"Simulation/pkg/internal/entities"
)

type Movable interface {
	MakeMove()
}

type Creature struct {
	entity.Entity
	Speed, Health, Satiety int
}
