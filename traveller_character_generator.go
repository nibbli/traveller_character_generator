package main

import (
    "fmt"
	"math/rand" // Import the math/rand package for random number generation
	"time"     // Import the time package to seed the random number generator
)

type Stat struct {
    strength        int
    dexterity       int
    endurance       int
    intelligence    int
    education       int
    socialStatus    int
}

type Character struct {
    name            string
    age             int
    stat            Stat
    title           string
    rank            string
    service         string
    termsServed     int
    isRetired       bool
    retirementPay   int
    isTASMember     bool
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
    var char Character

    // Prompt the user to enter their name.
	fmt.Print("Please enter your name: ")

	// Read the user's input and assign it to the name variable.
	fmt.Scanln(&char.name)
    
    char.stat.strength = roll_dice()
    char.stat.dexterity = roll_dice()
    char.stat.endurance = roll_dice()
    char.stat.intelligence = roll_dice()
    char.stat.education = roll_dice()
    char.stat.socialStatus = roll_dice()

	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Age: %d\n", char.age)
	fmt.Printf("Stat: %d\n", char.stat)
}