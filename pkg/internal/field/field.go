package field

import (
	"Simulation/pkg/internal/entities"
	"Simulation/pkg/internal/entities/creatures"
	"math/rand"
)

type Field struct {
	sizeX, sizeY int
	Entities     [][]entity.Positionable
	//Entities     map[[2]int]entity.Positionable
	//Rendered     [][]rune
}

func (f *Field) Init(sizeX, sizeY int) {
	if sizeX*sizeY < 24 {
		sizeX = 5
		sizeY = 5
	}
	f.sizeX = sizeX
	f.sizeY = sizeY
	//f.Entities = make(map[[2]int]entity.Positionable, 24)
	for i := 0; i < sizeY; i++ {
		//f.Rendered = append(f.Rendered, make([]rune, sizeX))
		f.Entities = append(f.Entities, make([]entity.Positionable, sizeX))
		//for j := 0; j < sizeX; j++ {
		//f.Rendered[i][j] = '.'
		//}
	}

	herbivoreCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 3
	for i := 0; i < herbivoreCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		//_, ok := f.Entities[[2]int{yPos, xPos}]
		//if ok {
		//	i--
		//	continue
		//}

		if f.Entities[yPos][xPos] != nil {
			i--
			continue
		}

		herbivore := creatures.Herbivore{}
		herbivore.SetPositions(xPos, yPos)
		herbivore.Health = rand.Int()%5 + 5
		herbivore.Speed = rand.Int()%2 + 1
		herbivore.Satiety = 5
		//f.Entities[[2]int{yPos, xPos}] = &herbivore
		f.Entities[yPos][xPos] = &herbivore
	}

	predatorCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < predatorCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		//_, ok := f.Entities[[2]int{yPos, xPos}]
		//if ok {
		//	i--
		//	continue
		//}

		if f.Entities[yPos][xPos] != nil {
			i--
			continue
		}

		predator := creatures.Predator{}
		predator.SetPositions(xPos, yPos)
		predator.Health = rand.Int()%5 + 5
		predator.Speed = rand.Int()%2 + 1
		predator.Satiety = 5
		predator.Force = rand.Int()%5 + 2
		//f.Entities[[2]int{yPos, xPos}] = &predator
		f.Entities[yPos][xPos] = &predator
	}

	treeCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < treeCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		//_, ok := f.Entities[[2]int{yPos, xPos}]
		//if ok {
		//	i--
		//	continue
		//}

		if f.Entities[yPos][xPos] != nil {
			i--
			continue
		}

		tree := entity.Tree{}
		tree.SetPositions(xPos, yPos)
		//f.Entities[[2]int{yPos, xPos}] = &tree
		f.Entities[yPos][xPos] = &tree
	}

	rockCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 2
	for i := 0; i < rockCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		//_, ok := f.Entities[[2]int{yPos, xPos}]
		//if ok {
		//	i--
		//	continue
		//}

		if f.Entities[yPos][xPos] != nil {
			i--
			continue
		}

		rock := entity.Rock{}
		rock.SetPositions(xPos, yPos)
		//f.Entities[[2]int{yPos, xPos}] = &rock
		f.Entities[yPos][xPos] = &rock
	}

	grassCount := rand.Int()%int(float32(sizeX)*float32(sizeY)*0.1) + 3
	for i := 0; i < grassCount; i++ {
		xPos := rand.Int() % sizeX
		yPos := rand.Int() % sizeY
		//_, ok := f.Entities[[2]int{yPos, xPos}]
		//if ok {
		//	i--
		//	continue
		//}

		if f.Entities[yPos][xPos] != nil {
			i--
			continue
		}

		grass := entity.Grass{}
		grass.SetPositions(xPos, yPos)
		//f.Entities[[2]int{yPos, xPos}] = &grass
		f.Entities[yPos][xPos] = &grass
	}
}

func (f *Field) Render() [][]rune {
	var renderedMap = make([][]rune, f.sizeY)
	for i := 0; i < f.sizeY; i++ {
		renderedMap[i] = make([]rune, f.sizeX)
	}
	//for k, v := range f.Entities {
	for i := 0; i < f.sizeY; i++ {
		for j := 0; j < f.sizeX; j++ {
			//switch v.(type) {
			switch f.Entities[i][j].(type) {
			case *entity.Grass:
				//f.Rendered[k[0]][k[1]] = '🌿'
				renderedMap[i][j] = '🌿'
			case *entity.Tree:
				//f.Rendered[k[0]][k[1]] = '🌲'
				renderedMap[i][j] = '🌲'
			case *entity.Rock:
				//f.Rendered[k[0]][k[1]] = '🪨'
				renderedMap[i][j] = '🪨'
			case *creatures.Predator:
				//f.Rendered[k[0]][k[1]] = '🐺'
				renderedMap[i][j] = '🐺'
			case *creatures.Herbivore:
				//f.Rendered[k[0]][k[1]] = '🐑'
				renderedMap[i][j] = '🐑'
			case nil:
				renderedMap[i][j] = '.'
			}
		}
	}
	return renderedMap
}
