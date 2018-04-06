package simple

import "fmt"

func CalculateFitness(member string) int {
	fitnessCount := 0
	for _, char := range member {
		if string(char) == "1" {
			fitnessCount++
		}
		message := fmt.Sprintf("Fitness level: %v", fitnessCount)
		fmt.Println(message)
	}

	return fitnessCount
}
