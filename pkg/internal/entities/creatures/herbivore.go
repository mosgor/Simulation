package creatures

import (
	"Simulation/pkg/internal/entities"
	"Simulation/pkg/internal/field"
)

type Herbivore struct {
	Creature
}

func (h *Herbivore) MakeMove(f *field.Field) {
	if h.Satiety == 0 {
		h.Health -= 2
		if h.Health <= 0 {
			f.RemoveEntity(h)
		}
	} else {
		h.Satiety--
	}
	pos := h.Positions()
	delete(f.ClosedCells, pos)

	path := f.FindNearest(pos, func(e field.Positionable) bool {
		_, ok := e.(*entity.Grass)
		return ok
	})
	if path == nil {
		return
	}
	next := h.Speed % len(path)
	for _, ok := f.ClosedCells[path[next]]; ok && next > 0; {
		next--
	}
	h.SetPositions(path[next][0], path[next][1])
	f.ClosedCells[path[next]] = struct{}{}
	if path[next][1] == path[len(path)-2][1] && path[next][0] == path[len(path)-2][0] {
		f.RemoveEntity(f.GetEntityAt(path[len(path)-1][0], path[len(path)-1][1]))
		h.Satiety++
	}
}
