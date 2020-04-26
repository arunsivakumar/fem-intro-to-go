package main

import "fmt"

func average(values ...float64) float64 {
	sum := 0.0
	for _, value := range values {
		sum += value
	}

	return sum / float64(len(values))
}

func main() {
	avg := average(10, 5, 7)
	fmt.Println(avg)
}
