package creatures

import "Simulation/pkg/internal/field"

type Predator struct {
	Creature
	Force int
}

func (p *Predator) MakeMove(f *field.Field) {
	p.Satiety--
	pos := p.Positions()
	delete(f.ClosedCells, pos)

	nearest := f.FindNearest([2]int{pos[1], pos[0]}, func(e field.Positionable) bool {
		_, ok := e.(*Herbivore)
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
		next := path[p.Speed]
		p.SetPositions(next[0], next[1])
	}
	pos = p.Positions()
	f.ClosedCells[pos] = struct{}{}
	if pos[1] == goal[1] && pos[0] == goal[0] {
		herbivore := nearest.(*Herbivore)
		herbivore.Health -= p.Force
		p.Satiety++
		if herbivore.Health <= 0 {
			f.RemoveEntity(herbivore)
		}
	}
}
