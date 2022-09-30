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
	consoleInput := bufio.NewReader(os.Stdin)
	fmt.Println("Press Enter to continue...")
	text, err := consoleInput.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Printf(text) // `text` is ONLY "\r\n" for an Enter keypress
	if strings.Compare(text, "\r\n") == 0 {
		fmt.Print("Starting timer...")
	}

	timer := time.NewTimer(5 * time.Second)

	// notifying the channel
	<-timer.C

	fmt.Println("Timer is inactivated")

	// csvReader()
}

func csvReader() {
	var score int = 0
	inputReader := bufio.NewReader(os.Stdin)

	// 1. open file
	recordFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}

	// 2. initialize reader
	reader := csv.NewReader(recordFile)

	// 3. read all records
	records, _ := reader.ReadAll()

	// 4. iterate through records
	for i := 0; i < len(records); i++ {
		fmt.Println("What is: " + records[i][0])

		input, err := inputReader.ReadString('\n')
		input = strings.Replace(input, "\r\n", "", -1) // Remove the delimiting chars on OS input

		// TODO : Replace ending character based on detected OS.
		// if OS == Windows {
		// 	input = strings.Replace(input, "\r\n", "", -1)
		// } else if OS == Linux {
		// 	input = strings.Replace(input, "\n", "", -1)
		// }

		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		if strings.Compare(input, records[i][1]) == 0 {
			score = score + 1
		}
	}

	fmt.Printf("Finished! Your score is: %v/%v", score, len(records))
}
