package main

import (
	"fmt" // A package in the Go standard library.
	"math/rand"
	"os"
	"unicode/utf8"
)

type organism struct {
	name string
	fit  int
}

func main() {
	// Set up
	inital := os.Args[1]
	goal := os.Args[2]

	// Initilize the first and goal organism
	initalOrg := organism{name: inital, fit: calcFit(inital, goal)}
	goalOrg := organism{name: goal, fit: calcFit(goal, goal)}

	lastOrg := generateGenerations(initalOrg, goalOrg)

	fmt.Println(lastOrg.name + "<-- Winner")

}

func generateGenerations(org, goalOrg organism) organism {
	var offspring []organism

	for i := 1; i <= 5; i++ {
		if rand.Intn(10) == 1 {
			offspring = append(offspring, mutate(org, goalOrg))
		} else {
			offspring = append(offspring, org)
		}
	}

	bestOffspring := calcBestOffs(offspring, goalOrg)

	if bestOffspring.name != goalOrg.name {
		fmt.Println(bestOffspring.name)
		generateGenerations(bestOffspring, goalOrg)
	}
	return bestOffspring
}

func mutate(org, goal organism) organism {
	// create the eco system (all possible chars)
	ecoSystem := append([]rune(org.name), []rune(goal.name)...)

	// get a random gene (char) from the eco system
	newGene := ecoSystem[rand.Intn(len(ecoSystem)-1)]

	// get a random gene (char) *position* to replace
	pos := rand.Intn(utf8.RuneCountInString(org.name) - 1)

	mutatedOrg := org.name[:pos] + string(newGene) + org.name[pos+1:]

	return organism{name: mutatedOrg, fit: calcFit(mutatedOrg, goal.name)}
}

func calcBestOffs(generation []organism, goalOrg organism) organism {
	var bestOrg organism

	return bestOrg
}

func calcFit(current, goal string) int {
	var fit int

	for i := range goal {
		if goal[i] == current[i] {
			fit++
		}
	}
	return fit
}
