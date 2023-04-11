package sim

type Population struct {
	susceptible int64
	exposed     int64
	infectious  int64
	recovered   int64
	dead        int64
}

func (population *Population) size() int64 {
	return population.susceptible +
		population.exposed +
		population.infectious +
		population.recovered
}

func createPopulation(size int64, infected int64) *Population {
	currentPopulation := Population{
		susceptible: size - infected,
		exposed:     0,
		infectious:  infected,
		recovered:   0,
		dead:        0,
	}
	return &currentPopulation
}
