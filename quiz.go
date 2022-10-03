package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
* Globals
* - lengthOfQuiz: easy access for tracking length of quiz across the program.
 */
var lengthOfQuiz int = 0
var flag_timer int

func init() {
	// Default timer is 30 seconds
	flag.IntVar(&flag_timer, "timer", 30, "Set the time duration (in seconds) for the quiz.")
	flag.Parse()
}

func main() {
	fmt.Printf("Timer flag is: %v\n", flag_timer)
	var scoreTracker int = 0

	consoleInput := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	text, err := consoleInput.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	if strings.Compare(text, "\r\n") == 0 { // `text` is ONLY "\r\n" for an Enter keypress
		fmt.Println("Starting timer...")
	}

	/*
	* Setup a timer that will "tick" in the background while the quiz takes place.
	 */
	timer := time.NewTimer(time.Duration(flag_timer) * time.Second) // Setup timer.
	go func() {                                                     // Use a goroutine to increment timer on separate
		<-timer.C
		fmt.Println("Timer 2 fired")

		// Timer has ended. Need to exit program.
		exitProgram(scoreTracker, true)
	}()

	// Start reading through the CSV for the quiz information.
	csvReader(&scoreTracker)
}

/*
* csvReader()
* - Reading the CSV file happens in here
* - Using a for-loop to make the quiz "work"
 */
func csvReader(scoreTracker *int) {
	inputReader := bufio.NewReader(os.Stdin)

	// Open file
	recordFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(recordFile) // Initialize reader
	records, _ := reader.ReadAll()      // Read all records
	lengthOfQuiz = len(records)         // update length of quiz tracker

	// Iterate through records
	// Format: record -> [0] - question, [1] - answer
	for i := 0; i < len(records); i++ {
		fmt.Println("What is: " + records[i][0])

		input, err := inputReader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1) // Remove the delimiting chars on OS input
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		// Check answer for correctness. Increment score if correct.
		if strings.Compare(input, records[i][1]) == 0 {
			*scoreTracker = *scoreTracker + 1
		}
	}

	// Exit the program because we completed the quiz
	exitProgram(*scoreTracker, false)
}

/*
* exitProgram()
* - Forcefully exits the program
* - Prints the final message
 */
func exitProgram(scoreTracker int, outOfTime bool) {
	if outOfTime {
		// Timer expired before finishing quiz!
		fmt.Println("TIME UP!")
		fmt.Printf("You got %v right, out of %v", scoreTracker, lengthOfQuiz)
	} else {
		fmt.Printf("Finished! You got %v right, out of %v", scoreTracker, lengthOfQuiz)
	}

	fmt.Println()
	os.Exit(3)
}

func detectOS() {
	// TODO : Replace ending character based on detected OS.
	// if OS == Windows {
	// 	input = strings.Replace(input, "\r\n", "", -1)
	// } else if OS == Linux {
	// 	input = strings.Replace(input, "\n", "", -1)
	// }
}
