package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const zero = 0b000000000
const one = 0b000000001
const two = 0b000000010
const three = 0b000000100
const four = 0b000001000
const five = 0b000010000
const six = 0b000100000
const seven = 0b001000000
const eight = 0b010000000
const nine = 0b100000000

type Board struct {
	values [9][9]int
}

func initialize() Board {
	rand.Seed(time.Now().Unix())

	return Board{}
}

func main() {
	board := initialize()

	bit2Int := map[int]int{
		zero:  0,
		one:   1,
		two:   2,
		three: 3,
		four:  4,
		five:  5,
		six:   6,
		seven: 7,
		eight: 8,
		nine:  9,
	}

	int2Bit := map[int]int{
		0: zero,
		1: one,
		2: two,
		3: three,
		4: four,
		5: five,
		6: six,
		7: seven,
		8: eight,
		9: nine,
	}

	// Populate board
	fmt.Println("Populate the board left to right, top to bottom:")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print("(", i, ",", j, "-> ")
			text, _ := reader.ReadString('\n')
			intValue, _ := strconv.Atoi(text)
			if intValue, ok := int2Bit[intValue]; ok {
				board.values[i][j] = int2Bit[intValue]
			} else {
				fmt.Println("1-9 only")
			}
		}
	}
}
