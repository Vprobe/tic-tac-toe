package ai

import (
	"math/rand"
	"time"

	field "github.com/Vprobe/tic-tac-toe/playingfield"
	"github.com/Vprobe/tic-tac-toe/settings"
	"github.com/Vprobe/tic-tac-toe/sign"
)

func MakeTurn(difficulty string, playerSign int) (int, int) {
	if difficulty == settings.Easy {
		return easyAI()
	}

	return normalAI(playerSign)
}

func easyAI() (int, int) {
	return generateCellRandomly()
}

func normalAI(playerSign int) (int, int) {
	// first turn: try filling center cell otherwise generating cell randomly
	if field.GetNumberOfEmptyCells() >= 8 {
		if field.GetSignByCell(1, 1) == sign.EmptyKey {
			return 1, 1
		}
		return generateCellRandomly()
	}

	// find shortest way to win for both players
	opponentSign := sign.CrossKey
	if playerSign == sign.CrossKey {
		opponentSign = sign.NoughtKey
	}

	aiShortestCombs := field.GetShortestCombinations(playerSign)
	playerShortestCombs := field.GetShortestCombinations(opponentSign)
	aiShortestTurn := getShortestTurn(aiShortestCombs)
	playerShortestTurn := getShortestTurn(playerShortestCombs)

	if playerShortestTurn < aiShortestTurn || (playerShortestTurn > 0 && aiShortestTurn == 0) {
		// block opponent's combination if:
		// - opponent's combination is shorter
		// - opponent has winning combination & you are not
		combs := playerShortestCombs[playerShortestTurn]
		comb := getFinalCombination(combs)
		possibleTurns := getPossibleTurns(comb)
		return getFinalTurn(possibleTurns)
	}

	if aiShortestTurn == 0 && playerShortestTurn == 0 {
		// there are no win combinations: generate cell randomly
		return generateCellRandomly()
	}

	// get your winning combination to finish
	combs := aiShortestCombs[aiShortestTurn]
	comb := getFinalCombination(combs)
	possibleTurns := getPossibleTurns(comb)

	return getFinalTurn(possibleTurns)
}

func generateCellRandomly() (int, int) {
	cells := field.GetAllEmptyCells()
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(cells))

	return cells[n][0], cells[n][1]
}

func getShortestTurn(combinations map[int][][3][3]int) int {
	res := 4 // impossible combination length: (max 3, min 0 - means there are no combinations left)
	if len(combinations) == 0 {
		res = 0
		return res
	}
	for k := range combinations {
		if len(combinations) > 0 && k < res {
			res = k
		}
	}

	return res
}

func getFinalCombination(combs [][3][3]int) [3][3]int {
	if len(combs) > 1 {
		// if there are more then 1 combs, choose one comb randomly
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(combs))
		return combs[n]
	}

	return combs[0]
}

func getPossibleTurns(comb [3][3]int) [][2]int {
	var possibleTurns [][2]int
	for keyRow, row := range comb {
		for keyCell, cell := range row {
			if cell == 1 {
				cellSign := field.GetSignByCell(keyRow, keyCell)
				if cellSign == sign.EmptyKey {
					var emptyCell = [2]int{keyRow, keyCell}
					possibleTurns = append(possibleTurns, emptyCell)
				}
			}
		}
	}

	return possibleTurns
}

func getFinalTurn(turns [][2]int) (int, int) {
	if len(turns) > 1 {
		// if there are more then 1 combs, choose one comb randomly
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(len(turns))
		return turns[n][0], turns[n][1]
	}

	return turns[0][0], turns[0][1]
}
