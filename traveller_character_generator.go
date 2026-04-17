package main

import (
	"bufio"
	"fmt"
	"math/rand/v2" // Import the math/rand package for random number generation
	"os"
	// Import the time package to seed the random number generator
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

func main() {
	char := Character{age: 18}
	var career Career

	// Prompt the user to enter their name.
	fmt.Print("Please enter your name: ")
	char.name = read_input()

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
		fmt.Print("What service do you wnat to enlist in? 1:Navy, 2:Marines, 3:Army, 4:Scouts, 4:Merchants, 6:Other (1-6): ")
		serviceChoice := read_input()

		if !(serviceChoice == "1" || serviceChoice == "2" || serviceChoice == "3" || serviceChoice == "4" || serviceChoice == "5" || serviceChoice == "6") {
			continue
		}

		switch serviceChoice {
		case "1":
			career = Career{name: "Navy", enlistment: 8, draft: 1, survival: 5, commission: 10, promotion: 8, reenlist: 6}
			if char.stat.intelligence >= 8 {
				enlistmentDM += 1
			}
			if char.stat.education >= 9 {
				enlistmentDM += 2
			}
			if char.stat.intelligence >= 7 {
				survivalDM += 2
			}
			if char.stat.socialStatus >= 9 {
				commisionDM += 1
			}
			if char.stat.education >= 8 {
				promotionDM += 1
			}
		case "2":
			career = Career{name: "Marines", enlistment: 9, draft: 2, survival: 6, commission: 9, promotion: 9, reenlist: 6}
			if char.stat.intelligence >= 8 {
				enlistmentDM += 1
			}
			if char.stat.strength >= 8 {
				enlistmentDM += 2
			}
			if char.stat.endurance >= 8 {
				survivalDM += 2
			}
			if char.stat.education >= 7 {
				commisionDM += 1
			}
			if char.stat.socialStatus >= 8 {
				promotionDM += 1
			}
		case "3":
			career = Career{name: "Army", enlistment: 5, draft: 3, survival: 5, commission: 5, promotion: 6, reenlist: 7}
			if char.stat.dexterity >= 6 {
				enlistmentDM += 1
			}
			if char.stat.endurance >= 5 {
				enlistmentDM += 2
			}
			if char.stat.education >= 6 {
				survivalDM += 2
			}
			if char.stat.endurance >= 7 {
				commisionDM += 1
			}
			if char.stat.education >= 7 {
				promotionDM += 1
			}
		case "4":
			career = Career{name: "Scouts", enlistment: 7, draft: 4, survival: 7, commission: 0, promotion: 0, reenlist: 3}
			if char.stat.intelligence >= 6 {
				enlistmentDM += 1
			}
			if char.stat.strength >= 8 {
				enlistmentDM += 2
			}
			if char.stat.endurance >= 9 {
				survivalDM += 2
			}
		case "5":
			career = Career{name: "Merchants", enlistment: 7, draft: 5, survival: 5, commission: 4, promotion: 10, reenlist: 4}
			if char.stat.strength >= 7 {
				enlistmentDM += 1
			}
			if char.stat.intelligence >= 6 {
				enlistmentDM += 2
			}
			if char.stat.intelligence >= 7 {
				survivalDM += 2
			}
			if char.stat.intelligence >= 6 {
				commisionDM += 1
			}
			if char.stat.intelligence >= 9 {
				promotionDM += 1
			}
		case "6":
			career = Career{name: "Other", enlistment: 3, draft: 6, survival: 5, commission: 0, promotion: 0, reenlist: 5}
			if char.stat.intelligence >= 9 {
				survivalDM += 2
			}
		}
		enlistmentRoll := roll_dice()
		fmt.Printf("enlistmentRoll %d\n", enlistmentRoll)
		fmt.Printf("enlistmentDM %d\n", enlistmentDM)
		fmt.Printf("career.enlistment %d\n", career.enlistment)
		if enlistmentRoll+enlistmentDM < career.enlistment {
			fmt.Printf("You were not enlisted in the %s\n", career.name)
			draftRoll := rand.UintN(6) + 1
			switch draftRoll {
			case 1:
				char.service = "Navy"
				career = Career{name: "Navy", enlistment: 8, draft: 1, survival: 5, commission: 10, promotion: 8, reenlist: 6}
				if char.stat.intelligence >= 7 {
					survivalDM += 2
				}
				if char.stat.socialStatus >= 9 {
					commisionDM += 1
				}
				if char.stat.education >= 8 {
					promotionDM += 1
				}
			case 2:
				char.service = "Marines"
				career = Career{name: "Marines", enlistment: 9, draft: 2, survival: 6, commission: 9, promotion: 9, reenlist: 6}
				if char.stat.endurance >= 8 {
					survivalDM += 2
				}
				if char.stat.education >= 7 {
					commisionDM += 1
				}
				if char.stat.socialStatus >= 8 {
					promotionDM += 1
				}
			case 3:
				char.service = "Army"
				career = Career{name: "Army", enlistment: 5, draft: 3, survival: 5, commission: 5, promotion: 6, reenlist: 7}
				if char.stat.education >= 6 {
					survivalDM += 2
				}
				if char.stat.endurance >= 7 {
					commisionDM += 1
				}
				if char.stat.education >= 7 {
					promotionDM += 1
				}
			case 4:
				char.service = "Scouts"
				career = Career{name: "Scouts", enlistment: 7, draft: 4, survival: 7, commission: 0, promotion: 0, reenlist: 3}
				if char.stat.endurance >= 9 {
					survivalDM += 2
				}
			case 5:
				char.service = "Merchants"
				career = Career{name: "Merchants", enlistment: 7, draft: 5, survival: 5, commission: 4, promotion: 10, reenlist: 4}
				if char.stat.intelligence >= 7 {
					survivalDM += 2
				}
				if char.stat.intelligence >= 6 {
					commisionDM += 1
				}
				if char.stat.intelligence >= 9 {
					promotionDM += 1
				}
			case 6:
				char.service = "Other"
				career = Career{name: "Other", enlistment: 3, draft: 6, survival: 5, commission: 0, promotion: 0, reenlist: 5}
				if char.stat.intelligence >= 9 {
					survivalDM += 2
				}
			}
		}
		break
	}

	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Age: %d\n", char.age)
	fmt.Printf("Stat: %+v\n", char.stat)
	fmt.Printf("Service: %s\n", char.service)
	fmt.Printf("Career: %+v\n", career)
}
