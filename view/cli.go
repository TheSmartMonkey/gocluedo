package main

import (
	"fmt"
	"log"
	"flag"

	c "github.com/TheSmartMonkey/gocluedo/controller"
)


func main() {
	nbPlayers := getArgs()
	initGame()
	play(nbPlayers)
	c.Game()
}

func getArgs() int {
	nbPlayers := flag.Int("n", 6, "number of players")
    flag.Parse()

    log.Println("Number of players: ", *nbPlayers)
	return *nbPlayers
}

func initGame()  {
	clueType := c.AskClueType()
	fmt.Println(clueType)
}

func play(nbPlayers int) {
	fmt.Println("text")
}
