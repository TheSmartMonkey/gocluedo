package controller

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"io/ioutil"
	"encoding/json"

	m "github.com/TheSmartMonkey/gocluedo/model"
)

// OpenClues open the clues.json file
func OpenClues() m.Clues {
	// Open file
	jsonFile, err := os.Open("../data/clues.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()

	// Read jsonFile
	var clues m.Clues
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &clues)

	return clues
}

// Ask input a question
func Ask(text string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text, ": ")
	response, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}

	return string(response), err
}

// AskClueType is the input question function
func AskClueType() string {
	fmt.Println("1 - persons")
	fmt.Println("2 - weapons")
	fmt.Println("3 - places")
	response, _ := Ask("Select clue type")

	if response == "1" {
		return "persons"
	} else if response == "2" {
		return "weapons"
	} else if response == "3" {
		return "places"
	} else {
		AskClueType()
	}
	return ""
}

// AskClue is the input question function
func AskClue(clueType string)  {
	clues := OpenClues()
	fmt.Println(clues)
}
