package simulation

import (
	"Simulation/pkg/internal/field"
	"fmt"
)

type Simulation struct {
	field.Field
	roundCounter int
}

func (s *Simulation) ShowMap() {
	rendered := s.Render()
	//for _, row := range s.Rendered {
	for _, row := range rendered {
		for _, cell := range row {
			fmt.Printf("%c\t", cell)
		}
		fmt.Println()
	}
}

func (s *Simulation) StartSimulation(xSize, ySize int) {
	s.Init(xSize, ySize)
	s.ShowMap()
}
