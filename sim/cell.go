package sim

type Cell int64

const (
	Free Cell = iota
	Susceptible
	Exposed
	Infected
	Recovered
	Dead
)
