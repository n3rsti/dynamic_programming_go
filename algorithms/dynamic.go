package algorithms

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	Weight float64
	Value  float64
}

func ItemsFromFile(filename string) ([]Item, float64) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)

	itemCount, _ := strconv.Atoi(parts[0])
	capacity, _ := strconv.ParseFloat(parts[1], 64)

	items := make([]Item, 0, itemCount)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		weight, _ := strconv.ParseFloat(line[0], 64)
		value, _ := strconv.ParseFloat(line[1], 64)

		items = append(items, Item{weight, value})
	}

	return items, capacity
}

func (i *Item) profitability() float64 {
	return i.Value / i.Weight
}

func DynamicSolution(items []Item, W float64) int {
	n := len(items)
	K := make([][]int, n+1)

	for i := range K {
		K[i] = make([]int, int(W)+1)
	}

	for i := 0; i <= n; i++ {
		for w := 0; w <= int(W); w++ {
			if i == 0 || w == 0 {
				K[i][w] = 0
			} else if int(items[i-1].Weight) <= w {
				K[i][w] = max(K[i-1][w], K[i-1][w-int(items[i-1].Weight)]+int(items[i-1].Value))
			} else {
				K[i][w] = K[i-1][w]
			}
		}
	}

	return K[n][int(W)]
}

func GreedySolution(items []Item, W float64) ([]Item, float64) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].profitability() > items[j].profitability()
	})

	fmt.Println(items)

	var result []Item
	weight := 0.0
	value := 0.0

	for _, item := range items {
		if weight+item.Weight <= W {
			result = append(result, item)
			weight += item.Weight
			value += item.Value
		}
	}

	return result, value
}

func generateCombinations(n int) [][]int {
	var combinations [][]int
	total := 1 << n

	for i := 0; i < total; i++ {
		var combination []int
		for j := 0; j < n; j++ {
			if (i>>j)&1 == 1 {
				combination = append(combination, 1)
			} else {
				combination = append(combination, 0)
			}
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

func BruteForceSolution(items []Item, capacity float64) ([]int, float64) {
	n := len(items)
	maxValue := 0.0
	var bestCombination []int

	combinations := generateCombinations(n)

	for _, combination := range combinations {
		totalValue := 0.0
		totalWeight := 0.0

		for i, include := range combination {
			if include == 1 {
				totalValue += items[i].Value
				totalWeight += items[i].Weight
			}
		}

		if totalWeight <= capacity && totalValue > maxValue {
			maxValue = totalValue
			bestCombination = combination
		}
	}

	return bestCombination, maxValue
}
