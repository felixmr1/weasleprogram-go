package main

import (
	"fmt" // A package in the Go standard library.
	"math/rand"
	"os"
)

type organism struct {
	name string
	fit  int
}

func main() {
	// Set up
	inital := os.Args[1]
	goal := os.Args[2]

	// Initilize the first and last organism
	initalOrg := organism{name: inital, fit: calcFit(inital, goal)}
	goalOrg := organism{name: goal, fit: calcFit(goal, goal)}

	generation := nextGen(initalOrg, goalOrg)

	bestOrg := calcBestChild(generation)

	fmt.Print(initalOrg)

}

func nextGen(org, goalOrg organism) []organism {
	var nextGenOrg organism
	var nextGen []organism

	for i := 1; i <= 5; i++ {
		if rand.Intn(10) == 1 {
			nextGenOrg := mutate(org, goalOrg)
		} else {
			nextGenOrg := org
		}
		nextGen := append(nextGen, nextGenOrg)
	}

	return nextGen
}

func mutate(org, goal organism) organism {
	// create the eco system (all possible chars)
	ecoSystem := append([]rune(org.name), []rune(goal.name)...)

	// get a random gene (char) from the eco system
	newGene := ecoSystem[rand.Intn(len(ecoSystem)-1)]

	// get a random gene (char) *position* to replace
	pos := rand.Intn(len(org.name) - 1)

	mutatedOrg := org.name[:pos] + string(newGene) + org.name[pos+1:]

	return organism{name: mutatedOrg, fit: calcFit(mutatedOrg, goal.name)}
}

func calcBestChild(generation []organism) organism {
	var bestOrg organism

	return bestOrg
}

func calcFit(current, goal string) int {
	var fit int

	if goal == goal {

	} else {
		for k, v := range goal {
			if goal[k] == current[k] {
				fit++
			}
		}
	}
	return fit
}
