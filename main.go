package main

import (
	"fmt"
	"main/algorithms"
)

func main() {
	items, capacity := algorithms.ItemsFromFile("input")

	fmt.Println(algorithms.DynamicSolution(items, capacity))
	fmt.Println(algorithms.GreedySolution(items, capacity))
	fmt.Println(algorithms.BruteForceSolution(items, capacity))
}
