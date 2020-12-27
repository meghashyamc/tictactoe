package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	models "github.com/meghashyamc/tictactoe/models"
)

const (
	humanVersusHumanType = "human versus human"
	humanVersusAIType    = "human versus AI"
)

var inputErr = "can't read input"

var reader = bufio.NewReader(os.Stdin)

func GetHumanPlayer(id int) *models.HumanPlayer {

	fmt.Println("Enter the name of player", id)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("could not read player name", err)
	}
	name = strings.TrimRight(name, "\n")

	return &models.HumanPlayer{Name: name, ID: id}
}

func GetGameType() string {

	var (
		human = "human"
		ai    = "ai"
	)
	for {
		fmt.Println("What type of game do you want to play? (Enter '", human, "' for a human versus human game and enter '", ai, "' for a human versus AI game")
		gameType, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("could not read player name", err)
		}
		gameType = strings.TrimRight(gameType, "\n")
		switch gameType {
		case human:
			return humanVersusHumanType

		case ai:
			return humanVersusAIType
		default:
			fmt.Println("Could not understand the type you entered -", gameType)
		}

	}

}

func GetNextMove(p models.Player) (int, int) {
	fmt.Println(p.GetName()+",", "enter next move's coordinates for your symbol")

	x, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	x = strings.TrimRight(x, "\n")
	if err != nil {
		log.Fatal(err)
	}
	xint, err := strconv.Atoi(x)
	if err != nil {
		log.Fatal(err)
	}

	y, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	y = strings.TrimRight(y, "\n")
	if err != nil {
		log.Fatal(err)
	}

	yint, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
	}

	return xint, yint
}
