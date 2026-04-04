package main

import (
    "fmt"
	"math/rand" // Import the math/rand package for random number generation
	"time"     // Import the time package to seed the random number generator
)

type Stat struct {
    strength        int range 1 to 15
    dexterity       int range 1 to 15
    endurance       int range 1 to 15
    intelligence    int range 1 to 15
    education       int range 1 to 15
    socialStatus    int range 1 to 15
}

type Character struct {
    name            string
    age             int = 18
    stat            Stat
    title           string
    rank            string
    service         string
    termsServed     int = 0
    isRetired       bool = false
    retirementPay   int
    isTASMember     bool = false
}

type Career struct {
    name        string
    enlistment  int
    draft       int
    survival    int
    commission  int
    promotion   int
    reenlist    int
}

// roll_dice simulates rolling two six-sided dice and returns their sum.
func roll_dice() int {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Roll the first die (1 to 6)
	die1 := rand.Intn(6) + 1

	// Roll the second die (1 to 6)
	die2 := rand.Intn(6) + 1

	// Calculate the sum of the two dice
	total := die1 + die2

	return total // Return only the total
}


func main() {
// Prompt the user to enter their name.
	fmt.Print("Please enter your name: ")

	// Read the user's input and assign it to the name variable.
	fmt.Scanln(&Character.name)

	&Character.stat.strength := roll_dice()
    &Character.stat.dexterity := roll_dice()
    &Character.stat.endurance := roll_dice()
    &Character.stat.intelligence := roll_dice()
    &Character.stat.education := roll_dice()
    &Character.stat.socialStatus := roll_dice()

	fmt.Printf("Name: %s\n", Character.name)
	fmt.Printf("Age: %d\n", Character.age)
	fmt.Printf("Stat: %d\n", Character.stat)
}