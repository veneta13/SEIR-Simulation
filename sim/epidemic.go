package sim

type epidemic struct {
	contactRate  float64
	recoveryRate float64
}

func newEpidemic(contactRate float64, recoveryRate float64) *epidemic {
	newEpidemic := epidemic{
		contactRate:  contactRate,
		recoveryRate: recoveryRate,
	}
	return &newEpidemic
}
