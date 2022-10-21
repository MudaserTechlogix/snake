package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/DyegoCosta/snake-game/snake"
	"github.com/sqweek/dialog"
)

func main() {
	var My_value_2 string
	My_value_2 = ""
	// var playerList [10]string

	// var index int
	// index = 0
	// var scoreList [10]int
	readFile, err := os.Open("score.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {

		// s := strings.Split(fileScanner.Text(), ",")
		// playerList[index] = s[0]
		// i, _ := strconv.Atoi(s[1])
		// scoreList[index] = i
		My_value_2 += fileScanner.Text() + "\n"
		// index = index + 1
	}

	readFile.Close()
	fmt.Println()
	if dialog.Message("%s", My_value_2).YesNo() {
		snake.NewGame().Start()
		for dialog.Message("%s", "Do you want to continue?").YesNo() {
			snake.NewGame().Start()
		}
	}

	dialog.Message("%s", My_value_2).YesNo()

}
