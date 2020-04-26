package main

import "fmt"

func main() {
	name := "Arun Sivakumar"

	for index, letter := range name {
		if index%2 == 1 {
			fmt.Println(string(letter))
		}
	}
}
