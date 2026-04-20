package main

import (
	"bufio"
	"fmt"
	"math/rand/v2" // Import the math/rand package for random number generation
	"os"
	"strconv"
)

type Stat struct {
	strength     int
	dexterity    int
	endurance    int
	intelligence int
	education    int
	socialStatus int
}

type Character struct {
	name          string
	age           int
	stat          Stat
	title         string
	rank          string
	service       string
	termsServed   int
	isRetired     bool
	retirementPay int
	isTASMember   bool
}

type Career struct {
	name       string
	enlistment int
	draft      int
	survival   int
	commission int
	promotion  int
	reenlist   int
}

// roll_dice simulates rolling two six-sided dice and returns their sum.
func roll_dice() int {

	// Roll the first die (1 to 6)
	die1 := rand.IntN(6) + 1

	// Roll the second die (1 to 6)
	die2 := rand.IntN(6) + 1

	// Calculate the sum of the two dice
	total := die1 + die2

	return total // Return only the total
}

func read_input() string {
	var input string
	reader := bufio.NewReader(os.Stdin) // Create a buffered reader

	for {
		// Read a line from the reader until a newline is encountered.
		// The '\n' argument tells bufio to stop reading when it sees a newline character.
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input:", err) // Handle errors (e.g., EOF)
			break                                    // Exit loop if there's an error
		}

		// Remove the trailing newline character (important!)
		input = line[:len(line)-1]

		break // Exit loop after successful read
	}
	return input
}

func setCareer(draftNumber int, career *Career, char *Character, enlistmentDM *int, survivalDM *int, commisionDM *int, promotionDM *int) {
	switch draftNumber {
	case 1:
		char.service = "Navy"
		career.name = "Navy"
		career.enlistment = 8
		career.draft = 1
		career.survival = 5
		career.commission = 10
		career.promotion = 8
		career.reenlist = 6
		if char.stat.intelligence >= 8 {
			*enlistmentDM += 1
		}
		if char.stat.education >= 9 {
			*enlistmentDM += 2
		}
		if char.stat.intelligence >= 7 {
			*survivalDM += 2
		}
		if char.stat.socialStatus >= 9 {
			*commisionDM += 1
		}
		if char.stat.education >= 8 {
			*promotionDM += 1
		}
	case 2:
		char.service = "Marines"
		career.name = "Marines"
		career.enlistment = 9
		career.draft = 2
		career.survival = 6
		career.commission = 9
		career.promotion = 9
		career.reenlist = 6
		if char.stat.intelligence >= 8 {
			*enlistmentDM += 1
		}
		if char.stat.strength >= 8 {
			*enlistmentDM += 2
		}
		if char.stat.endurance >= 8 {
			*survivalDM += 2
		}
		if char.stat.education >= 7 {
			*commisionDM += 1
		}
		if char.stat.socialStatus >= 8 {
			*promotionDM += 1
		}
	case 3:
		char.service = "Army"
		career.name = "Army"
		career.enlistment = 5
		career.draft = 3
		career.survival = 5
		career.commission = 5
		career.promotion = 6
		career.reenlist = 7
		if char.stat.dexterity >= 6 {
			*enlistmentDM += 1
		}
		if char.stat.endurance >= 5 {
			*enlistmentDM += 2
		}
		if char.stat.education >= 6 {
			*survivalDM += 2
		}
		if char.stat.endurance >= 7 {
			*commisionDM += 1
		}
		if char.stat.education >= 7 {
			*promotionDM += 1
		}
	case 4:
		char.service = "Scouts"
		career.name = "Scouts"
		career.enlistment = 7
		career.draft = 4
		career.survival = 7
		career.commission = 0
		career.promotion = 0
		career.reenlist = 3
		if char.stat.intelligence >= 6 {
			*enlistmentDM += 1
		}
		if char.stat.strength >= 8 {
			*enlistmentDM += 2
		}
		if char.stat.endurance >= 9 {
			*survivalDM += 2
		}
	case 5:
		char.service = "Merchants"
		career.name = "Merchants"
		career.enlistment = 7
		career.draft = 5
		career.survival = 5
		career.commission = 4
		career.promotion = 10
		career.reenlist = 4
		if char.stat.strength >= 7 {
			*enlistmentDM += 1
		}
		if char.stat.intelligence >= 6 {
			*enlistmentDM += 2
		}
		if char.stat.intelligence >= 7 {
			*survivalDM += 2
		}
		if char.stat.intelligence >= 6 {
			*commisionDM += 1
		}
		if char.stat.intelligence >= 9 {
			*promotionDM += 1
		}
	case 6:
		char.service = "Other"
		career.name = "Other"
		career.enlistment = 3
		career.draft = 6
		career.survival = 5
		career.commission = 0
		career.promotion = 0
		career.reenlist = 5
		if char.stat.intelligence >= 9 {
			*survivalDM += 2
		}
	}
}

func main() {
	char := Character{age: 18}
	var career Career

	// Prompt the user to enter their name.
	fmt.Print("Please enter your name: ")
	char.name = "Marcus Deltoid"
	// char.name = read_input()

	char.stat.strength = roll_dice()
	char.stat.dexterity = roll_dice()
	char.stat.endurance = roll_dice()
	char.stat.intelligence = roll_dice()
	char.stat.education = roll_dice()
	char.stat.socialStatus = roll_dice()

	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Age: %d\n", char.age)
	fmt.Printf("Stat: %+v\n", char.stat)

	var enlistmentDM int // enlistment dice roll modifier
	var survivalDM int   // survival dice roll modifier
	var commisionDM int  // commission dice roll modifier
	var promotionDM int  // promotion dice roll modifier

	for {
		fmt.Print("What service do you want to enlist in? 1:Navy, 2:Marines, 3:Army, 4:Scouts, 4:Merchants, 6:Other (1-6): ")
		serviceChoice := "1"
		// serviceChoice := read_input()

		if !(serviceChoice == "1" || serviceChoice == "2" || serviceChoice == "3" || serviceChoice == "4" || serviceChoice == "5" || serviceChoice == "6") {
			continue
		}

		serviceChoiceInt, _ := strconv.Atoi(serviceChoice)

		setCareer(serviceChoiceInt, &career, &char, &enlistmentDM, &survivalDM, &commisionDM, &promotionDM)

		enlistmentRoll := roll_dice()

		fmt.Printf("enlistmentRoll %d\n", enlistmentRoll)
		fmt.Printf("enlistmentDM %d\n", enlistmentDM)
		fmt.Printf("career.enlistment %d\n", career.enlistment)

		if enlistmentRoll+enlistmentDM < career.enlistment {
			fmt.Printf("You were not enlisted in the %s\n", career.name)
			draftRoll := rand.IntN(6) + 1
			setCareer(draftRoll, &career, &char, &enlistmentDM, &survivalDM, &commisionDM, &promotionDM)
		}
		break
	}

	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Age: %d\n", char.age)
	fmt.Printf("Stat: %+v\n", char.stat)
	fmt.Printf("Service: %s\n", char.service)
	fmt.Printf("Career: %+v\n", career)
}
