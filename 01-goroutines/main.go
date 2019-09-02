package main

import (
	"fmt"
	"math/rand"

	//"sync"
	"time"
)

func findGoldenNumber(randGenerator *rand.Rand, magicNumber int) int {
	numTries := 0
	for {
		numTries++
		i := randGenerator.Intn(10000)

		if i == magicNumber {
			break
		}
	}
	return numTries
}

func main() {

	magicNumbers := []int{42, 606, 81, 233}

	// wg := sync.WaitGroup{}
	// wg.Add(len(magicNumbers))

	for _, magicNumber := range magicNumbers {

		// create a new random number generator
		randGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

		//go func(magicNumber int) {
		numTries := findGoldenNumber(randGenerator, magicNumber)
		fmt.Printf("Number of tries for %d: %d\n", magicNumber, numTries)
		//wg.Done()
		//}(magicNumber)
	}

	//wg.Wait()
}
