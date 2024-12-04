package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Id       int
	GameSets []GameSet
}

type GameSet struct {
	Cubes []Cube
}

type Cube struct {
	Count int
	Color string
}

var games []Game
var gameConfig = map[string]int{"red": 12, "green": 13, "blue": 14}

func main() {
	data, err := os.ReadFile("data")
	if err != nil {
		log.Fatal("Error to load the data")
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	for i, line := range lines {
		games = append(games, getGames(line, i+1))
	}

	validGames, minCubeSet := getResultValidGames()
	fmt.Printf("The sum of the valid games is: %d\n", validGames)
	fmt.Printf("The sum of the powers is: %d", minCubeSet)

}

func getResultValidGames() (validGames int, minCubeSet int) {
	for _, game := range games {
		minConfig := map[string]int{"red": 0, "green": 0, "blue": 0}
		isGameValid := true
		for _, set := range game.GameSets {
			for _, cube := range set.Cubes {

				if minConfig[cube.Color] < cube.Count || minConfig[cube.Color] == 0 {
					minConfig[cube.Color] = cube.Count
				}

				if gameConfig[cube.Color] < cube.Count {
					isGameValid = false
				}
			}
		}
		if isGameValid {
			validGames += game.Id
		}

		minCubeSet += (minConfig["red"] * minConfig["green"] * minConfig["blue"])

	}
	return
}

func getGames(line string, id int) Game {
	game := Game{Id: id}
	g := strings.Split(line, ":")
	sets := strings.Split(g[1], ";")
	for _, set := range sets {
		gameSet := GameSet{}
		setCubes := strings.Split(set, ",")
		for _, cube := range setCubes {
			c := strings.Split(strings.Trim(cube, " "), " ")
			count, err := strconv.Atoi(string(c[0]))
			if err != nil {
				log.Panic("Error parsing value", string(c[0]))
			}
			gameSet.Cubes = append(gameSet.Cubes, Cube{
				Count: int(count),
				Color: string(c[1]),
			})

		}

		game.GameSets = append(game.GameSets, gameSet)
	}
	return game
}
