package main

import "genetic-algorithms/simple"

func main() {
	// commands to run specified genetic algorithm
	// mostFitPair := simple.SelectMostFitPair(population)
	// log.Printf("Most fit: %v, Second most fit: %v", mostFitPair[0], mostFitPair[1])

	// member := "11001001"
	// mutated := simple.MutateMember(member)
	// fmt.Println("Mutated member: ", mutated)

	// parent1 := "11110000"
	// parent2 := "00001111"
	// simple.Crossover([]byte(parent1), []byte(parent2))

	// newPop := simple.ReplaceLeastFitMember("11111111", 2, population)
	// for _, member := range newPop {
	// 	fmt.Println(member)
	// }

	population := []string{"11001001", "10010000", "00000000", "11101110", "10101010", "01010101"}
	simple.SimpleGeneticAlgorithm(population)
}
