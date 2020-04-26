package main

func average(a, b, c int) float32 {
	return float32((a + b + c) / 3)
}
func main() {
	avg := average(2, 3, 4)
	println(avg)
}
