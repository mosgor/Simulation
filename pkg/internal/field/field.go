package field

import (
	"errors"
)

type Positionable interface {
	Positions() [2]int
	SetPositions(int, int)
}

type Movable interface {
	MakeMove(f *Field)
}

type Field struct {
	SizeX, SizeY int
	Entities     []Positionable
	ClosedCells  map[[2]int]struct{}
}

func (f *Field) Init(sizeX, sizeY int) {
	if sizeX*sizeY < 24 {
		sizeX = 5
		sizeY = 5
	}
	f.SizeX = sizeX
	f.SizeY = sizeY
	f.Entities = make([]Positionable, 0, 24)
	f.ClosedCells = make(map[[2]int]struct{}, 24)
}

func (f *Field) AddEntity(e Positionable) error {
	coords := e.Positions()
	if _, ok := f.ClosedCells[coords]; ok {
		return errors.New("this position is already closed")
	}
	f.Entities = append(f.Entities, e)
	f.ClosedCells[coords] = struct{}{}
	return nil
}

func (f *Field) RemoveEntity(e Positionable) {
	for i, v := range f.Entities {
		if v == e {
			coords := e.Positions()
			delete(f.ClosedCells, coords)
			f.Entities = append(f.Entities[i:], f.Entities[i+1:]...)
			break
		}
	}
}

func (f *Field) GetEntityAt(x, y int) Positionable {
	for _, e := range f.Entities {
		pos := e.Positions()
		if pos[0] == y && pos[1] == x {
			return e
		}
	}
	return nil
}

func (f *Field) FindNearest(start [2]int, match func(Positionable) bool) Positionable {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][2]int{start}
	visited := make(map[[2]int]struct{})
	visited[start] = struct{}{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, d := range directions {
			neighbor := [2]int{current[0] + d[0], current[1] + d[1]}
			if neighbor[0] < 0 || neighbor[1] < 0 || neighbor[0] >= f.SizeX || neighbor[1] >= f.SizeY {
				continue
			}
			entity := f.GetEntityAt(neighbor[0], neighbor[1])
			if entity != nil {
				if match(entity) {
					return entity
				} else {
					visited[neighbor] = struct{}{}
					continue
				}
			}
			if _, ok := visited[neighbor]; ok {
				continue
			}
			queue = append(queue, neighbor)
			visited[neighbor] = struct{}{}
		}
	}
	return nil
}

func (f *Field) FindPathBFS(start, goal [2]int) [][2]int {
	directions := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][2]int{start}
	visited := make(map[[2]int]struct{})
	parent := make(map[[2]int][2]int)
	visited[start] = struct{}{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == goal {
			var path [][2]int
			for at := goal; at != start; at = parent[at] {
				path = append([][2]int{at}, path...)
			}
			path = append([][2]int{start}, path...)
			return path
		}
		for _, d := range directions {
			neighbor := [2]int{current[0] + d[0], current[1] + d[1]}
			if neighbor[0] < 0 || neighbor[1] < 0 || neighbor[0] >= f.SizeX || neighbor[1] >= f.SizeY {
				continue
			}
			if _, ok := visited[neighbor]; ok || f.GetEntityAt(neighbor[0], neighbor[1]) != nil {
				continue
			}
			queue = append(queue, neighbor)
			visited[neighbor] = struct{}{}
			parent[neighbor] = current
		}
	}
	return nil
}

func (f *Field) Render(renderFunc func(e Positionable) rune) [][]rune {
	var renderedMap = make([][]rune, f.SizeY)
	for i := 0; i < f.SizeY; i++ {
		renderedMap[i] = make([]rune, f.SizeX)
		for j := 0; j < f.SizeX; j++ {
			renderedMap[i][j] = '.'
		}
	}
	for _, v := range f.Entities {
		renderedMap[v.Positions()[0]][v.Positions()[1]] = renderFunc(v)
	}
	return renderedMap
}
