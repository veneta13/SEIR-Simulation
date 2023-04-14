package sim

import "math/rand"

type World struct {
	board [gridSize][gridSize]Cell
}

func initRandomWorld(population *Population) *World {
	world := World{
		board: [gridSize][gridSize]Cell{},
	}

	var current int64 = 0
	for current < population.infectious {
		x := rand.Int63n(gridSize)
		y := rand.Int63n(gridSize)

		if world.board[x][y].cellType == Free {
			world.board[x][y].cellType = Infected
			current++
		}
	}

	return &world
}
