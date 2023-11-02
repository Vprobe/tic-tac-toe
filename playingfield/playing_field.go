package playingfield

import (
	"fmt"

	msg "github.com/Vprobe/tic-tac-toe/messages"
	"github.com/Vprobe/tic-tac-toe/sign"
)

const (
	Col string = "column"
	Row string = "row"

	// Row & Col validation
	MinRowCol int = 1
	MaxRowCol int = 3
)

var playingField [3][3]int
var winnerCombs = [][3][3]int{
	{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 0},
	},
	{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0},
	},
	{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 1},
	},
	{
		{1, 1, 1},
		{0, 0, 0},
		{0, 0, 0},
	},
	{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0},
	},
	{
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 1},
	},
	{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	},
	{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0},
	},
}

func Draw(winComb [3][3]int) {
	fmt.Print("\n   -------------\n")
	for row := 0; row < len(playingField); row++ {
		fmt.Print("   ")
		for column := 0; column < len(playingField[row]); column++ {
			if column == 0 {
				fmt.Print("|")
			}
			if winComb[row][column] == 1 {
				msg.WinCombination(" " + sign.GetSignByNumber(playingField[row][column]) + " ")
			} else {
				fmt.Print(" " + sign.GetSignByNumber(playingField[row][column]) + " ")
			}
			if column < len(playingField[row])-1 {
				fmt.Print("|")
			}
			if column == len(playingField[row])-1 {
				fmt.Print("|")
			}
		}
		fmt.Println("")
		if row < len(playingField)-1 {
			fmt.Print("   -------------\n")
		}
	}
	fmt.Print("   -------------\n\n")
}

func FillCell(row, col, playerSign int) {
	playingField[row-1][col-1] = playerSign
}

func IsCellEmpty(row, col int) bool {
	return playingField[row-1][col-1] == sign.EmptyKey
}

func IsWinnerExist(playerSign int) (bool, [3][3]int) {
	var res bool
	var winComb [3][3]int
	for _, comb := range winnerCombs {
		i := 0
		for keyRow, row := range comb {
			for keyCell, cell := range row {
				if cell == 0 {
					continue
				}
				if playingField[keyRow][keyCell] == playerSign {
					i++
				}
			}
			if i == 3 {
				winComb = comb
				break
			}
		}
		if i == 3 {
			res = true
			break
		}
	}

	return res, winComb
}

func IsDrawGame() bool {
	draw := true
	for _, row := range &playingField {
		for _, cell := range row {
			if cell == sign.EmptyKey {
				draw = false
				break
			}
		}
		if !draw {
			break
		}
	}

	return draw
}

func GetAllEmptyCells() [][2]int {
	var cells [][2]int
	for keyRow, row := range &playingField {
		for keyCell, cell := range row {
			if cell == sign.EmptyKey {
				coord := [2]int{keyRow, keyCell}
				cells = append(cells, coord)
			}
		}
	}

	return cells
}

func GetNumberOfEmptyCells() int {
	var res int
	for _, row := range &playingField {
		for _, cell := range row {
			if cell == sign.EmptyKey {
				res++
			}
		}
	}

	return res
}

func GetShortestCombinations(playerSign int) map[int][][3][3]int {
	possibleCombs := make(map[int][][3][3]int, 0)
	for _, comb := range winnerCombs {
		i := 0
		for keyRow, row := range comb {
			hasOpponentSign := false
			for keyCell, cell := range row {
				if cell == 0 {
					continue
				}
				if playingField[keyRow][keyCell] != playerSign && playingField[keyRow][keyCell] != sign.EmptyKey {
					i = 0
					hasOpponentSign = true
					break
				}
				if playingField[keyRow][keyCell] == sign.EmptyKey {
					i++
				}
			}
			if hasOpponentSign {
				// get next comb
				break
			}
		}
		if i > 0 {
			possibleCombs[i] = append(possibleCombs[i], comb)
		}
	}

	return possibleCombs
}

func GetSignByCell(row, cell int) int {
	return playingField[row][cell]
}
