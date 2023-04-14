package sim

type CellType int64

const (
	Free CellType = iota
	Susceptible
	Exposed
	Infected
	Recovered
	Dead
)

type Cell struct {
	cellType CellType
	lifetime int64
}
