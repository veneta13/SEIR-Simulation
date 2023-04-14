package sim

import "math/rand"

type World struct {
	board [WorldX][WorldY]Cell
}

func initRandomWorld(population *Population) *World {
	world := World{
		board: [WorldX][WorldY]Cell{},
	}

	var current int64 = 0
	for current < population.infectious {
		x := rand.Int63n(WorldX)
		y := rand.Int63n(WorldY)

		if world.board[x][y] == Free {
			world.board[x][y] = Infected
			current++
		}
	}

	return &world
}
