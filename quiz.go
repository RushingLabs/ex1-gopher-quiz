package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// var timerLength int64 = 5 // 30
	var scoreTracker int = 0

	consoleInput := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	text, err := consoleInput.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	if strings.Compare(text, "\r\n") == 0 { // `text` is ONLY "\r\n" for an Enter keypress
		fmt.Print("Starting timer...")
	}

	timer2 := time.NewTimer(45 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")

		// Timer has ended. Need to exit program.
		exitProgram(scoreTracker)
	}()

	csvReader(&scoreTracker)
}

func csvReader(scoreTracker *int) {
	// var score int = 0
	inputReader := bufio.NewReader(os.Stdin)

	// 1. open file
	recordFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(recordFile) // 2. initialize reader
	records, _ := reader.ReadAll()      // 3. read all records

	// 4. iterate through records
	for i := 0; i < len(records); i++ {
		fmt.Println("What is: " + records[i][0])

		input, err := inputReader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1) // Remove the delimiting chars on OS input

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		if strings.Compare(input, records[i][1]) == 0 {
			*scoreTracker = *scoreTracker + 1
		}
	}

	exitProgram(*scoreTracker)
}

func exitProgram(scoreTracker int) {
	// fmt.Printf("Finished! Your score is: %v/%v", scoreTracker, len(records))
	fmt.Printf("Finished! You got %v right.", scoreTracker)
	fmt.Println("Program exit. Status code: 3")
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
