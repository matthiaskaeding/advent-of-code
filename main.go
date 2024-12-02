package main

import (
	"advent/day01"
	"fmt"
)

func main() {
	ids0, ids1, err := day01.ReadIds("day01/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(ids0[:5])
	fmt.Println(ids1[:5])

	distances := day01.SumDistances(ids0, ids1, err)
	fmt.Printf("Sum of distances: %v\n", distances)

	similarity_score := day01.CompSimilarityScore(ids0, ids1)
	fmt.Printf("Similarity score: %v\n", similarity_score)

}
