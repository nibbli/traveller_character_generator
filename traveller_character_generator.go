package main

import (
	"bufio"
	"fmt"
	"math/rand" // Import the math/rand package for random number generation
	"os"
	"time" // Import the time package to seed the random number generator
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

	for {
		fmt.Print("What service? 1:Navy, 2:Marines, 3:Army, 4:Scouts, 4:Merchants, 6:Other (1-6): ")
		service_choice := read_input()

		if !(service_choice == "1" || service_choice == "2" || service_choice == "3" || service_choice == "4" || service_choice == "5" || service_choice == "6") {
			continue
		}
		switch service_choice {
		case "1":
			char.service = "Navy"
			career = Career{name: "Navy", enlistment: 8, draft: 1, survival: 5, commission: 10, promotion: 8, reenlist: 6}
		case "2":
			char.service = "Marines"
			career = Career{name: "Marines", enlistment: 9, draft: 2, survival: 6, commission: 9, promotion: 9, reenlist: 6}
		case "3":
			char.service = "Army"
			career = Career{name: "Army", enlistment: 5, draft: 3, survival: 5, commission: 5, promotion: 6, reenlist: 7}
		case "4":
			char.service = "Scouts"
			career = Career{name: "Scouts", enlistment: 7, draft: 4, survival: 7, commission: 0, promotion: 0, reenlist: 3}
		case "5":
			char.service = "Merchants"
			career = Career{name: "Merchants", enlistment: 7, draft: 5, survival: 5, commission: 4, promotion: 10, reenlist: 4}
		case "6":
			char.service = "Other"
			career = Career{name: "Other", enlistment: 3, draft: 6, survival: 5, commission: 0, promotion: 0, reenlist: 5}
		}
		break
	}

	fmt.Printf("Name: %s\n", char.name)
	fmt.Printf("Age: %d\n", char.age)
	fmt.Printf("Stat: %+v\n", char.stat)
	fmt.Printf("Service: %s\n", char.service)
	fmt.Printf("Career: %+v\n", career)

}
