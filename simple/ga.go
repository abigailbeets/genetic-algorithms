package simple

import (
	"fmt"
	"math/rand"
	"time"
)

// thought - make population map that stores string and fitness value
// in order to easily locate the least/most fit members and to
// avoid checking fitness more than once for a member

// var populationMap = make(map[string]int)

const (
	MemberLength int = 8
)

func SimpleGeneticAlgorithm(population []string) {
	var hasConverged bool
	generation := 1
	for !hasConverged {
		// time.Sleep(time.Second * 1)
		highestFitnessScore := getHighestFitnessScore(population)
		message := fmt.Sprintf("Generation: %v --- Highest Fitness: %v, Population Size: %v", generation, highestFitnessScore, len(population))
		fmt.Println(message)

		// select most fit pair from population
		parents := SelectMostFitPair(population)

		// have parents reproduce children
		children := Crossover([]byte(parents[0]), []byte(parents[1]))

		// check if children should mutate, and do so
		for index, child := range children {
			if shouldMutate() {
				children[index] = MutateMember(child)
			}
		}

		// get fittest child and least fittest member to compare
		fittestChild, fittestChildFitness := getFittestChild(children[0], children[1])
		leastFitMemberIndex, leastFitMemberFitness := getLeastFitMember(population)

		if fittestChildFitness > leastFitMemberFitness {
			population = replaceLeastFitMember(fittestChild, leastFitMemberIndex, population)
		}

		for _, member := range population {
			fitness := CalculateFitness(member)
			if fitness == MemberLength {
				fmt.Println("Found child with maximum fitness score!")
				hasConverged = true
				break
			}
		}

		if !hasConverged {
			generation++
		}
	}

	fmt.Println("Finished on generation: ", generation)
}

func CalculateFitness(member string) int {
	fitnessCount := 0
	for _, char := range member {
		if string(char) == "1" {
			fitnessCount++
		}
	}
	// message := fmt.Sprintf("Fitness level: %v", fitnessCount)
	// fmt.Println(message)

	return fitnessCount
}

func SelectMostFitPair(population []string) []string {
	var mostFit string
	var secondMostFit string
	var mostFitFitness int
	var secondMostFitFitness int

	for _, member := range population {
		// fmt.Println("Member: ", member)
		fitness := CalculateFitness(member)
		if fitness > mostFitFitness {
			// message := fmt.Sprintf("Current most fit member: %v with fitness score: %v", member, fitness)
			// fmt.Println(message)

			// move former most fit down a peg - loser
			secondMostFit = mostFit
			secondMostFitFitness = mostFitFitness

			// make current member most fit
			mostFit = member
			mostFitFitness = fitness
		} else if fitness > secondMostFitFitness {
			// message := fmt.Sprintf("Current second most fit member: %v with fitness score: %v", member, fitness)
			// fmt.Println(message)

			secondMostFit = member
			secondMostFitFitness = fitness
		}
	}

	mostFitPair := []string{mostFit, secondMostFit}
	return mostFitPair
}

func Crossover(parent1, parent2 []byte) []string {
	// select random crossover point
	// assumes parents are same length
	// swap characters up to, not including, crossover point
	// return two children (parent1 and parent2 crossed over)
	min := 2
	max := MemberLength

	rand.Seed(time.Now().Unix())
	crossoverPoint := rand.Intn(max-min) + min

	// fmt.Println("Max: ", max)
	// fmt.Println("Crossover Point: ", crossoverPoint)

	// instantiate empty children
	child1 := make([]byte, MemberLength)
	child2 := make([]byte, MemberLength)

	// add swapped values to children up until crossover point
	for i := 0; i < crossoverPoint; i++ {
		child1[i] = parent2[i]
		child2[i] = parent1[i]
	}

	// add remaining un-swapped values to children
	for i := crossoverPoint; i < MemberLength; i++ {
		child1[i] = parent1[i]
		child2[i] = parent2[i]
	}

	// fmt.Println("Child1: ", string(child1))
	// fmt.Println("Child2: ", string(child2))

	//parents should remain unchanged
	// fmt.Println("Parent1: ", string(parent1))
	// fmt.Println("Parent2: ", string(parent2))

	children := []string{string(child1), string(child2)}
	return children
}

func MutateMember(member string) string {
	min := 0
	max := MemberLength - 3

	rand.Seed(time.Now().Unix())
	mutateStartingIndex := rand.Intn(max-min+1) + min
	mutateLength := rand.Intn(3-1) + 1

	// fmt.Println("Mutate starting index: ", mutateStartingIndex)
	// fmt.Println("Mutate length: ", mutateLength)

	for i := 0; i <= mutateLength; i++ {
		flipped := flip(string(member[mutateStartingIndex+i]))
		member = member[:mutateStartingIndex+i] + flipped + member[mutateStartingIndex+i+1:]
	}

	return member
}

// helper functions

func getFittestChild(child1, child2 string) (string, int) {
	child1Fitness := CalculateFitness(child1)
	child2Fitness := CalculateFitness(child2)

	if child1Fitness >= child2Fitness {
		return child1, child1Fitness
	}
	return child2, child2Fitness
}

func getLeastFitMember(population []string) (int, int) {
	var leastFitMemberIndex int
	leastFitMemberFitness := MemberLength

	for index, member := range population {
		fitness := CalculateFitness(member)
		if fitness < leastFitMemberFitness {
			leastFitMemberIndex = index
			leastFitMemberFitness = fitness
		}
	}

	return leastFitMemberIndex, leastFitMemberFitness
}

func getHighestFitnessScore(population []string) int {
	highestFitness := 0
	for _, member := range population {
		fitness := CalculateFitness(member)
		if fitness > highestFitness {
			highestFitness = fitness
		}
	}

	return highestFitness
}

func replaceLeastFitMember(newMember string, index int, population []string) []string {
	// fmt.Println("Old member: ", population[index])
	population[index] = newMember

	// fmt.Println("New member: ", population[index])

	return population
}

func flip(char string) string {
	if char == "1" {
		return "0"
	}
	return "1"
}

func shouldMutate() bool {
	min := 1
	max := 100

	// 5% chance of mutation in child
	rand.Seed(time.Now().Unix())
	result := rand.Intn(max-min) + min

	if result <= 5 {
		return true
	}
	return false
}
