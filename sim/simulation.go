package sim

import "math/rand"

type Simulation struct {
	population Population
	world      World
	clock      int64
}

func simCellSusceptible(x int64, y int64, readWorld *World, writeWorld *World) {
	probabilityOfInfection := 0

	for tempX := x - 1; tempX <= x+1; tempX++ {
		for tempY := y - 1; tempY <= y+1; tempY++ {
			if readWorld[tempX][tempY].cellType == Infected {
				probabilityOfInfection++
			}
		}
	}

	if rand.Intn(8) < probabilityOfInfection {
		writeWorld[x][y].cellType = Exposed
		writeWorld[x][y].countdown = IncubationPeriod
	}
}

func simCellExposed(x int64, y int64, readWorld *World, writeWorld *World) {
	if readWorld[x][y].countdown == 1 {
		writeWorld[x][y].cellType = Infected
		writeWorld[x][y].countdown = SpreadPeriod
		return
	}

	writeWorld[x][y].cellType = Exposed
	writeWorld[x][y].countdown = readWorld[x][y].countdown - 1
}

func simCell(x int64, y int64, readWorld *World, writeWorld *World) {
	if readWorld[x][y].cellType == Susceptible {
		simCellSusceptible(x, y, readWorld, writeWorld)
		return
	}

	if readWorld[x][y].cellType == Exposed {
		simCellExposed(x, y, readWorld, writeWorld)
		return
	}
}

func PlaySimulation() {
	return
}
