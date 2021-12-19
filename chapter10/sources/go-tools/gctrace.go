package main

func main() {
	m := make(map[interface{}]interface{})
	v := "hello, gc"

	for i := 1; i < 1<<31; i++ {
		m[i] = v
	}
}
