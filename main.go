package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomSlices := make([]int, 100)

	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(100)
		randomSlices[i] = randomNumber
	}

	fmt.Printf("Random Slices Value is %v\n", randomSlices)
	fmt.Printf("Random Slices Length is %v\n", len(randomSlices))
	fmt.Printf("Random Slices Capacity is %v\n", cap(randomSlices))

	fmt.Println("Looping through the Random Slice")

	for v := range randomSlices {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
		} else if v%2 == 0 {
			fmt.Println("Two!")

		} else if v%3 == 0 {
			fmt.Println("Three!")
		} else {
			fmt.Println("Never Mind")
		}
	}

}
