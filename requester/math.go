package requester

func sumFloat64(nn []float64) float64 {
	var sum float64
	for _, n := range nn {
		sum = sum + n
	}
	return sum
}

func avgFloat64(nn []float64) float64 {
	return sumFloat64(nn) / float64(len(nn))
}
