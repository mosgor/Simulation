package entity

type Entity struct {
	xPos, yPos int
}

func (e *Entity) Positions() [2]int {
	return [2]int{e.xPos, e.yPos}
}

func (e *Entity) SetPositions(x, y int) {
	e.xPos, e.yPos = x, y
}
