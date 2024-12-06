package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
	"time"
)

type Player struct {
	Position  [2]int // x, y
	Direction int    // 0=up, 1=right, 2=down, 3=left
}

var directions = [4][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

var width = 0
var height = 0

var obstacles map[[2]int]bool = make(map[[2]int]bool)
var player Player
var playerPositionsVisited = make(map[[2]int]bool)
var originalPlayerPosition = [2]int{0, 0}

func parseInput() {
	var y = 0
	file, err := os.Open("input.txt")

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char == '#' {
				obstacles[[2]int{x, y}] = true
			}

			if char == '^' {
				player = Player{[2]int{x, y}, 0}
				originalPlayerPosition = [2]int{x, y}
			}
		}

		y++
		width = len(line)
	}

	height = y
}

var shouldPrint = true

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printGrid(ply Player, obs map[[2]int]bool) {
	if !shouldPrint {
		return
	}

	shouldPrint = false

	clearScreen()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if ply.Position[0] == x && ply.Position[1] == y {

				switch ply.Direction {
				case 0:
					fmt.Printf(color.CyanString("^"))
				case 1:
					fmt.Printf(color.CyanString(">"))
				case 2:
					fmt.Printf(color.CyanString("v"))
				case 3:
					fmt.Printf(color.CyanString("<"))
				}

			} else if _, ok := obs[[2]int{x, y}]; ok {
				fmt.Printf(color.RedString("#"))
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	time.Sleep(50 * time.Millisecond)
	shouldPrint = true
}

func isPlayerInGrid() bool {
	return player.Position[0] >= 0 && player.Position[0] < width && player.Position[1] >= 0 && player.Position[1] < height
}

func checkPhantomObstacle(x int, y int) bool {
	player.Position = originalPlayerPosition
	player.Direction = 0
	var position = [2]int{x, y}

	if _, ok := obstacles[position]; ok {
		return false // obstacle is already there
	}

	if position == originalPlayerPosition {
		return false // player is already there so no phantom obstacle
	}

	var copiedObstacles = make(map[[2]int]bool)
	for k, v := range obstacles {
		copiedObstacles[k] = v
	}
	// add phantom obstacle
	copiedObstacles[position] = true

	var positionsWentTo [][]int

	for isPlayerInGrid() {
		var newPosition = [2]int{player.Position[0] + directions[player.Direction][0], player.Position[1] + directions[player.Direction][1]}
		if _, ok := copiedObstacles[newPosition]; ok {
			player.Direction = (player.Direction + 1) % 4
			continue
		}
		// move player
		player.Position = newPosition
		//go printGrid(player, copiedObstacles)
		var timesLocationVisited = 0
		for _, pos := range positionsWentTo {
			if pos[0] == player.Position[0] && pos[1] == player.Position[1] {
				timesLocationVisited++
			}
		}

		if timesLocationVisited > 5 {
			return true
		}

		positionsWentTo = append(positionsWentTo, []int{player.Position[0], player.Position[1]})
	}

	return false
}

func main() {
	parseInput()
	var phantomObstacleLoops = 0

	fmt.Printf("Start position: %v\n", player.Position)
	fmt.Printf("Start direction: %d\n", player.Direction)
	fmt.Printf("Obstacles: %d\n", len(obstacles))
	fmt.Printf("Width: %d\n", width)
	fmt.Printf("Height: %d\n", height)

	time.Sleep(2 * time.Second)

	for isPlayerInGrid() {
		var newPosition = [2]int{player.Position[0] + directions[player.Direction][0], player.Position[1] + directions[player.Direction][1]}
		if _, ok := obstacles[newPosition]; ok {
			player.Direction = (player.Direction + 1) % 4
			continue
		}

		playerPositionsVisited[player.Position] = true
		player.Position = newPosition
	}

	fmt.Printf("Player positions visited: %d\n", len(playerPositionsVisited))

	for y := 0; y < height; y++ {
		fmt.Printf("Checking row %d, phantom obstacles found: %d\n", y, phantomObstacleLoops)
		for x := 0; x < width; x++ {
			if checkPhantomObstacle(x, y) {
				phantomObstacleLoops++
			}
		}
	}

	fmt.Printf("Phantom obstacles looped: %d\n", phantomObstacleLoops)
}
