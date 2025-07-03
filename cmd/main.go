package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"test_game/internal/actions"
)

var gameData *actions.GameData

func handleCommand(com string) string {
	commands := strings.Split(com, " ")
	return gameData.Actions(commands)
}

func initGame() {
	gameData = actions.InitGameData()
}

func main() {
	initGame()
	for {
		commands, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Println(handleCommand(commands[:len(commands)-1]))
	}
}
