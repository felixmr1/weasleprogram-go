package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type organism struct {
	name []rune
	fit  int
}

func main() {
	// Set up
	goal := []rune(os.Args[1])
	initial := []rune(os.Args[2])
	generations := 0

	// Seed the random generator
	rand.Seed(time.Now().UTC().UnixNano())

	// Initilize the first and goal organism
	goalOrg := organism{name: goal, fit: calcFit(goal, goal)}
	initialOrg := organism{name: initial, fit: calcFit(initial, goal)}

	winner := generateGenerations(initialOrg, goalOrg, generations)

	fmt.Printf("%s -- Goal! \n", string(winner.name))

}

func generateGenerations(org, goalOrg organism, generations int) organism {
	var offspring []organism
	generations++

	// Loop through the amount of offspring you want the next generation to have
	// Select an arbitrary chance for any given offspring to mutate
	amountOffspring := 5
	for i := 0; i < amountOffspring; i++ {
		if rand.Intn(5) == 1 { // 20% chance
			offspring = append(offspring, mutate(org, goalOrg))
		} else {
			offspring = append(offspring, org)
		}
	}

	// Get the best fitted offspring
	bestOffspring := calcBestOffs(offspring, goalOrg)

	fmt.Printf("%d: %s \n", generations, string(bestOffspring.name))

	// Check if the best offspring is equal to the goal organism
	if string(bestOffspring.name) != string(goalOrg.name) {
		// If not equal, call this function again. With best fitted offspring
		generateGenerations(bestOffspring, goalOrg, generations)
	}

	return bestOffspring
}

func mutate(org, goalOrg organism) organism {

	// If we have the goalOrg then return then dont mutate it
	if string(org.name) == string(goalOrg.name) {
		return org
	}

	// create the eco system (all possible runes)
	ecoSystem := []rune("abcdefghijklmnopqrstuvwxyz ")

	// get a random gene (rune) from the eco system
	newGene := ecoSystem[rand.Intn(len(ecoSystem))]

	// Get the indexes of all missplaced genomes
	var badGenIndex []int
	for i := range org.name {
		if org.name[i] != goalOrg.name[i] {
			badGenIndex = append(badGenIndex, i)
		}
	}

	// randomly select an index to replace
	pos := badGenIndex[rand.Intn(len(badGenIndex))]

	// doing this since i cant append 3 items
	temp := append(org.name[:pos], newGene)
	mutatedOrg := append(temp, org.name[pos+1:]...)

	return organism{name: mutatedOrg, fit: calcFit(mutatedOrg, goalOrg.name)}
}

func calcBestOffs(generation []organism, goalOrg organism) organism {

	var bestOrg organism

	for i := range generation[:len(generation)-1] {
		if generation[i].fit >= generation[i+1].fit {
			bestOrg = generation[i]
		} else {
			bestOrg = generation[i+1]
		}
	}

	return bestOrg
}

func calcFit(org, goalOrg []rune) int {
	var fit int

	for i := range goalOrg {
		if goalOrg[i] == org[i] {
			fit++
		}
	}
	return fit
}
