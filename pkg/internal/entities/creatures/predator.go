package creatures

import (
	"Simulation/pkg/internal/field"
)

type Predator struct {
	Creature
	Force int
}

func (p *Predator) MakeMove(f *field.Field) {
	if p.Satiety == 0 {
		p.Health -= 2
		if p.Health <= 0 {
			f.RemoveEntity(p)
		}
	} else {
		p.Satiety--
	}
	pos := p.Positions()
	delete(f.ClosedCells, pos)

	path := f.FindNearest(pos, func(e field.Positionable) bool {
		_, ok := e.(*Herbivore)
		return ok
	})
	if path == nil {
		return
	}
	next := p.Speed % len(path)
	for _, ok := f.ClosedCells[path[next]]; ok && next > 0; {
		next--
	}
	p.SetPositions(path[next][0], path[next][1])
	f.ClosedCells[path[next]] = struct{}{}
	if path[next][1] == path[len(path)-2][1] && path[next][0] == path[len(path)-2][0] {
		herbivore := f.GetEntityAt(path[len(path)-1][0], path[len(path)-1][1]).(*Herbivore)
		herbivore.Health -= p.Force
		if herbivore.Health <= 0 {
			p.Satiety++
			f.RemoveEntity(herbivore)
		}
	}
}
