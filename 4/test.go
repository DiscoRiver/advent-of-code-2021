package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var pickedNumbers []int

type bingoBoard struct {
	boardString [25]int

	rows [5][11]int
	columns [5][11]int

	won bool
	calledNum int

	rin []int
	cin []int

	mu sync.Mutex
}

func NewBingoBoard(in [25]int) *bingoBoard {
	bb := &bingoBoard{
		boardString: in,

		rows: [5][11]int{},
		columns: [5][11]int{},

		rin: []int{},
		cin: []int{},
	}

	bb.setRows()
	bb.setColumns()

	return bb
}

func (bb *bingoBoard) setRows() {
	rows := [5][11]int{}
	for i := 0; i < 25; i=i+5 {
		rows[i/5] = [11]int{0, bb.boardString[i], bb.boardString[i+1], bb.boardString[i+2], bb.boardString[i+3], bb.boardString[i+4]}
	}
	bb.rows = rows
}

func (bb *bingoBoard) setColumns() {
	columns := [5][11]int{}
	for i := 0; i < 5; i++ {
		columns[i] = [11]int{0, bb.boardString[i], bb.boardString[i+5], bb.boardString[i+10], bb.boardString[i+15], bb.boardString[i+20]}
	}
	bb.columns = columns
}

func (bb *bingoBoard) getUnmarkedNumbers() []int {
	var unmarkedNums []int

	for j := range bb.rows {
		for i := 6; i < 11; i++ {
			if bb.rows[j][i] == 0 {
				unmarkedNums = append(unmarkedNums, bb.rows[j][i-5])
			}
		}
	}
	return unmarkedNums
}

func (bb *bingoBoard) hasWon() bool {
	return bb.won
}

func (bb *bingoBoard) processNumberSet(nums []int) {
	for n := range nums {
		for i := 5; i < 11; i++ {
			for j := range bb.rows {
				for k := 1; k < 6; k++ {
					if bb.rows[j][k] == nums[n] {
						bb.rows[j][k+5] = 1
					}

					win := bb.rows[j][6:]
					count := 0

					for w := range win {
						if win[w] == 1 {
							count++
						}
					}

					if count == 5 {
						bb.rows[j][0] = 1
						bb.calledNum = nums[n]
						bb.won = true
						return
					}
				}
			}

			for j := range bb.columns {
				for k := 1; k < 6; k++ {
					if bb.columns[j][k] == nums[n] {
						bb.columns[j][k+5] = 1
					}

					win := bb.columns[j][6:]

					count := 0
					for w := range win {
						if win[w] == 1 {
							count++
						}
					}

					if count == 5 {
						bb.columns[j][0] = 1
						bb.calledNum = nums[n]
						bb.won = true
						return
					}
				}
			}
		}
	}
}

func main() {
	boards := readInput()

	numBoards := len(boards)
	wonBoards := map[int]struct{}{}

	for n := 0; n < len(pickedNumbers); n = n + 5 {
		for i := range boards {
			boards[i].processNumberSet(pickedNumbers[n : n+5])

			if boards[i].hasWon() {
				if len(wonBoards) == 0 {
					fmt.Println("**** FIRST WINNER ****")
					fmt.Println("Board ", i)

					um := boards[i].getUnmarkedNumbers()
					sum := 0

					for i := range um {
						sum += um[i]
					}

					fmt.Println("Unmarked Numbers: ", um)
					fmt.Println("Last Called: ", boards[i].calledNum)
					fmt.Println("Sum of uncalled numbers: ", sum)
					fmt.Println("lastCalled * uncalled: ", sum*boards[i].calledNum)
				}
				wonBoards[i] = struct{}{}
			}

			if numBoards == len(wonBoards) {
				lastWon := NewBingoBoard(boards[i].boardString)
				for n := 0; n < len(pickedNumbers); n = n + 5 {
					lastWon.processNumberSet(pickedNumbers[n : n+5])

					if lastWon.won {
						fmt.Println("\n**** LAST WINNER ****")

						um := lastWon.getUnmarkedNumbers()
						sum := 0

						for i := range um {
							sum += um[i]
						}

						fmt.Println("Unmarked Numbers: ", um)
						fmt.Println("Last Called: ", lastWon.calledNum)
						fmt.Println("Sum of uncalled numbers: ", sum)
						fmt.Println("lastCalled * uncalled: ", lastWon.calledNum*sum)
						os.Exit(0)
					}
				}

			}
		}
	}
}

func readInput() []*bingoBoard {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	var bbs []*bingoBoard

	scanner := bufio.NewScanner(file)
	lastLine := 0
	var block []int
	for scanner.Scan() {
		lastLine++

		if lastLine == 1 {
			s := strings.Split(scanner.Text(), ",")
			var tmp []int
			for i := range s {
				num, err := strconv.Atoi(s[i])
				if err != nil {
					panic(err)
				}
				tmp = append(tmp, num)
			}
			pickedNumbers = tmp
			continue
		} else {
			if scanner.Text() != "" {
				s := strings.Split(scanner.Text(), " ")
				for i := range s {
					num, err := strconv.Atoi(s[i])
					if err != nil {
						panic(err)
					}
					block = append(block, num)
				}
			}
		}

		if len(block) == 25 {
			var blockArr [25]int
			copy(blockArr[:], block)

			board := NewBingoBoard(blockArr)
			bbs = append(bbs, board)

			block = nil
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return bbs
}