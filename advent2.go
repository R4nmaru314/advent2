package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const file = "inputs.txt"

type Score struct {
	enemy rune
	me    rune
	score int
}

var part1 = []Score{
	{'A', 'X', 4},
	{'A', 'Y', 8},
	{'A', 'Z', 3},
	{'B', 'X', 1},
	{'B', 'Y', 5},
	{'B', 'Z', 9},
	{'C', 'X', 7},
	{'C', 'Y', 2},
	{'C', 'Z', 6},
}

var part2 = []Score{
	{'A', 'X', 3},
	{'A', 'Y', 4},
	{'A', 'Z', 8},
	{'B', 'X', 1},
	{'B', 'Y', 5},
	{'B', 'Z', 9},
	{'C', 'X', 2},
	{'C', 'Y', 6},
	{'C', 'Z', 7},
}

func main() {
	// Open the file
	file, _ := os.Open(file)
	scanner := bufio.NewScanner(file)
	scorePart1 := 0
	scorePart2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		enemy := rune(line[0])
		me := rune(line[2])

		scorePart1 += calculateScore(enemy, me, part1)
		scorePart2 += calculateScore(enemy, me, part2)
	}
	log.Println("score1:", scorePart1)
	log.Println("score2:", scorePart2)
}

func calculateScore(enemy rune, me rune, scores []Score) int {

	for _, s := range scores {
		if s.enemy == enemy && s.me == me {
			return s.score
		}
	}
	panic(fmt.Sprintf("no score found for input: %c, %c", enemy, me))
}
