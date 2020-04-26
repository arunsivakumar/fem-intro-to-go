package main

import "fmt"

func avgscores(scores [5]float64) float64 {
	sum := 0.0
	for _, value := range scores {
		sum += value
	}
	return sum / float64(len(scores))
}

var initialPet map[string]string = map[string]string{
	"fido":  "dog",
	"nancy": "cat",
}

var initialGroceries = []string{"beans", "peas"}

func doesPetExisit(petName string) bool {
	_, ok := initialPet[petName]
	return ok
}

func add(groceries ...string) []string {
	food := initialGroceries
	for _, value := range groceries {
		food = append(food, value)
	}
	return food
}

func main() {
	scores := [5]float64{2.0, 3.0, 4.0, 5.0, 6.0}
	fmt.Println(avgscores(scores))

	petExists := doesPetExisit("lido")
	fmt.Println(petExists)

	foods := add("bread", "milk")
	fmt.Println(foods)
}
