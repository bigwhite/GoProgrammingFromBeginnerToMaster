package bar

func Bar(step, count int) int {
	sum := 1
	for i := 0; i < count; i++ {
		sum *= step
	}
	return sum
}
