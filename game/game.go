package game

import (
	"github.com/Vprobe/tic-tac-toe/ai"
	"github.com/Vprobe/tic-tac-toe/input"
	msg "github.com/Vprobe/tic-tac-toe/messages"
	"github.com/Vprobe/tic-tac-toe/player"
	field "github.com/Vprobe/tic-tac-toe/playingfield"
	"github.com/Vprobe/tic-tac-toe/settings"
	"github.com/Vprobe/tic-tac-toe/sign"
)

func Run() {
	gs := settings.NewSettings()
	msg.Greetings()
	turnOrder := options(gs)
	outputPlayersInfo(turnOrder)
	gameLoop(turnOrder, gs)
}

func options(gs *settings.Settings) [2]*player.Player {
	input.SelectMode(gs)

	if gs.GetSelectedMode() == settings.PvAI {
		input.SelectDifficulty(gs)
	}

	side1 := input.SelectSide()
	player1 := player.NewPlayer("Player1", side1, false)

	side2 := sign.CrossKey
	if player1.GetSign() == sign.CrossKey {
		side2 = sign.NoughtKey
	}

	isAI := gs.GetSelectedMode() == settings.PvAI
	player2 := player.NewPlayer("Player2", side2, isAI)

	if player1.GetSign() == sign.CrossKey {
		turnOrder := [2]*player.Player{player1, player2}
		return turnOrder
	}

	turnOrder := [2]*player.Player{player2, player1}
	return turnOrder
}

func outputPlayersInfo(turnOrder [2]*player.Player) {
	players := [2]map[string]string{{"name": "", "signName": "", "sign": ""}, {"name": "", "signName": "", "sign": ""}}
	for i, p := range turnOrder {
		players[i]["name"] = p.GetName()
		players[i]["signName"] = sign.GetSignNameByNumber(p.GetSign())
		players[i]["sign"] = sign.GetSignByNumber(p.GetSign())
	}

	msg.OutputPlayersInfo(players)
}

func gameLoop(turnOrder [2]*player.Player, gs *settings.Settings) {
	game := true
	winComb := [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for game {
		for _, v := range turnOrder {
			field.Draw(winComb)
			getPlayersInput(v, gs)

			// make turn
			field.FillCell(v.GetRow(), v.GetColumn(), v.GetSign())

			if exist, winnerComb := field.IsWinnerExist(v.GetSign()); exist {
				msg.Congratulations(v.GetName(), sign.GetSignNameByNumber(v.GetSign()), sign.GetSignByNumber(v.GetSign()))
				field.Draw(winnerComb)
				game = false
				break
			}

			if field.IsDrawGame() {
				msg.DrawGame()
				field.Draw(winComb)
				game = false
				break
			}
		}
	}
}

func getPlayersInput(player *player.Player, gs *settings.Settings) {
	for {
		if player.IsAI() {
			// AI generate input
			row, cell := ai.MakeTurn(gs.GetSelectedDifficulty(), player.GetSign())
			player.SetRow(row + 1)
			player.SetColumn(cell + 1)
			msg.Info("AI turn")
		} else {
			input.ReadPlayerInput(player)
		}

		// validate players input
		if !field.IsCellEmpty(player.GetRow(), player.GetColumn()) {
			msg.Error("Cell is not empty")
			continue
		}

		break
	}
}
