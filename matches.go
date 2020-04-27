package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func display(matches int) {

	var i int = 0
	fmt.Println("number of matches left : ", matches)
	for i < matches {
		fmt.Print("|")
		i++
	}
	fmt.Print("\n")
}

func main() {

	var matches int = 13
	var player bool = true

	fmt.Print("Welcome to the game of matches !\n")
	display(matches)

	for matches != 0 {
		if player == true {
			fmt.Print("Player one turn !\n")
		} else {
			fmt.Print("Player two turn !\n")
		}
	takematches:
		fmt.Print("Select beetwen 1 and 3 matches to take\n")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		take, _ := strconv.Atoi(scanner.Text())

		if take < 1 || take > 3 {
			fmt.Print("You can take only 1,2 or 3 matches ! \n")
			goto takematches
		} else if take > matches {
			fmt.Println("There only", matches, "matches left.")
			display(matches)
			goto takematches
		} else {
			matches = matches - take
			display(matches)
			if player == true {
				player = false
			} else {
				player = true
			}
		}
		if matches == 0 {
			if player == true {
				fmt.Print("Player one win !\n")
			} else {
				fmt.Print("Player two win !\n")
			}
		}
		fmt.Println("Matches you have take : ", take)
	}
}
