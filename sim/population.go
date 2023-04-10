package sim

type population struct {
	susceptible int64
	exposed     int64
	infectious  int64
	recovered   int64
}

func createPopulation(size int64, infected int64) *population {
	currentPopulation := population{
		susceptible: size - infected,
		infectious:  infected,
		recovered:   0,
	}
	return &currentPopulation
}
