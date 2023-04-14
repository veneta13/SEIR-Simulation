package sim

import "math/rand"

type World [gridSize][gridSize]Cell

func initRandomWorld(population *Population) *World {
	world := World{}

	var current int64 = 0
	for current < population.infectious {
		x := rand.Int63n(gridSize)
		y := rand.Int63n(gridSize)

		if world[x][y].cellType == Free {
			world[x][y].cellType = Infected
			current++
		}
	}

	return &world
}
