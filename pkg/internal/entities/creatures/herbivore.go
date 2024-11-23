package creatures

import (
	"Simulation/pkg/internal/entities"
	"Simulation/pkg/internal/field"
)

type Herbivore struct {
	Creature
}

func (h *Herbivore) MakeMove(f *field.Field) {
	h.Satiety--
	pos := h.Positions()
	delete(f.ClosedCells, pos)

	nearest := f.FindNearest([2]int{pos[1], pos[0]}, func(e field.Positionable) bool {
		_, ok := e.(*entity.Grass)
		return ok
	})
	if nearest == nil {
		return
	}
	goalPos := nearest.Positions()
	start := [2]int{pos[1], pos[0]}
	goal := [2]int{goalPos[1], goalPos[0]}
	path := f.FindPathBFS(start, goal)
	if len(path) > 1 {
		next := path[h.Speed]
		h.SetPositions(next[0], next[1])
	}
	pos = h.Positions()
	f.ClosedCells[pos] = struct{}{}
	if pos[1] == goal[1] && pos[0] == goal[0] {
		f.RemoveEntity(nearest)
		h.Satiety++
	}
}
