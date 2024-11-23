package simulation

import (
	"Simulation/pkg/internal/entities"
	"Simulation/pkg/internal/entities/creatures"
	"Simulation/pkg/internal/field"
	"fmt"
	"math/rand/v2"
	"time"
)

type Simulation struct {
	field.Field
	roundCounter int
}

func (s *Simulation) showMap() {
	mp := s.Render(func(e field.Positionable) rune {
		switch e.(type) {
		case *entity.Grass:
			return '🌿'
		case *entity.Tree:
			return '🌲'
		case *entity.Rock:
			return '🪨'
		case *creatures.Predator:
			return '🐺'
		case *creatures.Herbivore:
			return '🐑'
		default:
			return '.'
		}
	})
	for _, row := range mp {
		for _, cell := range row {
			fmt.Printf("%c\t", cell)
		}
		fmt.Println()
	}
}

func (s *Simulation) addGrass() {
	grassCount := rand.Int()%int(float32(s.SizeX)*float32(s.SizeY)*0.1) + 3
	for i := 0; i < grassCount; i++ {
		xPos := rand.Int() % s.SizeX
		yPos := rand.Int() % s.SizeY
		grass := entity.Grass{}
		grass.SetPositions(xPos, yPos)
		err := s.AddEntity(&grass)
		if err != nil {
			i--
		}
	}
}

func (s *Simulation) StartSimulation(sizeX, sizeY int) {
	s.roundCounter++
	fmt.Println("Round: ", s.roundCounter)
	s.Init(sizeX, sizeY)
	herbivoreCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 3
	for i := 0; i < herbivoreCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		herbivore := creatures.Herbivore{}
		herbivore.SetPositions(xPos, yPos)
		herbivore.Health = rand.Int()%5 + 5
		herbivore.Speed = rand.Int()%2 + 1
		herbivore.Satiety = 5
		err := s.AddEntity(&herbivore)
		if err != nil {
			i--
		}
	}

	predatorCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < predatorCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		predator := creatures.Predator{}
		predator.SetPositions(xPos, yPos)
		predator.Health = rand.Int()%5 + 5
		predator.Speed = rand.Int()%2 + 1
		predator.Satiety = 5
		predator.Force = rand.Int()%5 + 2
		err := s.AddEntity(&predator)
		if err != nil {
			i--
		}
	}

	treeCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < treeCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		tree := entity.Tree{}
		tree.SetPositions(xPos, yPos)
		err := s.AddEntity(&tree)
		if err != nil {
			i--
		}
	}

	rockCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < rockCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		rock := entity.Rock{}
		rock.SetPositions(xPos, yPos)
		err := s.AddEntity(&rock)
		if err != nil {
			i--
		}
	}

	s.addGrass()
	s.showMap()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		if i%5 == 0 {
			s.addGrass()
		}
		s.nextTurn()
	}
}

func (s *Simulation) nextTurn() {
	s.roundCounter++
	for _, e := range s.Entities {
		if movable, ok := e.(field.Movable); ok {
			movable.MakeMove(&s.Field)
		}
	}
	fmt.Println("\nRound: ", s.roundCounter)
	s.showMap()
}
