package main

import (
	"Simulation/pkg/simulation"
)

func main() {
	sim := simulation.Simulation{
		SizeX:        15,
		SizeY:        15,
		Rounds:       15,
		GrowthRate:   5,
		TurnDuration: 1,
	}
	sim.StartSimulation()
}
