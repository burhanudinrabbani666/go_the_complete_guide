package main

import (
	"fmt"
)

type FloatMap map[string]float64

func (floatMap FloatMap) output() {
	fmt.Println(floatMap)
}

func main() {

	coursesRatings := make(FloatMap, 3)

	coursesRatings["Go"] = 4.7
	coursesRatings["React"] = 4.8
	coursesRatings["Angular"] = 4.8

	// Index becoming Key when in map strcture
	for index, value := range coursesRatings {
		fmt.Println(index, value)
	}
}
