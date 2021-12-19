package foo

func Foo(step, count int) int {
	sum := 0
	for i := 0; i < count; i++ {
		sum += step
	}
	return sum
}
