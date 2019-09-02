package main

import (
	"fmt"
	"math/rand"

	"time"
)

func avg(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total / len(nums)
}

func findGoldenNumber(randGenerator *rand.Rand, magicNumber int) int {
	numTries := 0
	for {
		numTries++
		i := randGenerator.Intn(1000000)

		if i == magicNumber {
			break
		}
	}
	return numTries
}

func main() {

	magicNumbers := []int{42, 606, 81, 233}

	resultsChan := make(chan int, len(magicNumbers))
	defer close(resultsChan)

	for _, magicNumber := range magicNumbers {

		// create a new random number generator
		randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

		go func(magicNumber int) {

			numTries := findGoldenNumber(randGenerator, magicNumber)
			resultsChan <- numTries

		}(magicNumber)
	}

	totals := []int{}

	for i := 0; i < len(magicNumbers); i++ {
		result := <-resultsChan
		totals = append(totals, result)
	}

	fmt.Println("Total number of tries:", totals)
	fmt.Println("Average number of tries:", avg(totals))
}
