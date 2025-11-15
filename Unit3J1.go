// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-15
// This program checks if a contestant may enter the pie eating contest
// based on their weight.

package main

import (
	"fmt"
)

func main() {

	// Constants
	const MIN_WEIGHT float64 = 77.0
	const MAX_WEIGHT float64 = 105.0

	var weight float64

	// Input
	fmt.Print("How much do you weigh? ")
	fmt.Scan(&weight)

	// Decision
	if weight >= MIN_WEIGHT && weight <= MAX_WEIGHT {
		fmt.Println("You may enter the contest.")
	} else {
		fmt.Println("You may NOT enter the contest.")
	}

	// Done message
	fmt.Println("\nDone.")
}
