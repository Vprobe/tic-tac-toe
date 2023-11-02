package messages

import (
	"fmt"

	"github.com/fatih/color"
)

var yellow = color.New(color.Bold, color.FgHiYellow).PrintlnFunc()
var cyan = color.New(color.Bold, color.FgHiCyan).PrintlnFunc()
var green = color.New(color.FgHiGreen).PrintlnFunc()
var red = color.New(color.FgHiRed).PrintlnFunc()
var magenta = color.New(color.Bold, color.FgHiMagenta).PrintlnFunc()
var white = color.New(color.FgHiWhite).PrintFunc()
var greenBold = color.New(color.Bold, color.FgHiGreen).PrintFunc()

func Title(text string) {
	yellow(text)
}

func Info(text string) {
	cyan(">>> " + text)
}

func Success(text string) {
	green("[SUCCESS] " + text)
}

func Error(text string) {
	red("[ERROR] " + text)
}

func Option(text string) {
	magenta(text)
}

func Choice() {
	white("-> ")
}

func WinCombination(text string) {
	greenBold(text)
}

func SelectGameMode() {
	Info("Select game mode [1 or 2]:")
	Option("1. Player versus Player")
	Option("2. Player versus AI")
	Choice()
}

func SelectDifficulty() {
	Info("Select AI difficulty [1 or 2]:")
	Option("1. Easy")
	Option("2. Normal")
	Choice()
}
func SelectSide() {
	Info("Select option [1 or 2]:")
	Option("1. [X] Crosses (always go first)")
	Option("2. [O] Noughts")
	Choice()
}

func Greetings() {
	Title("###############################")
	Title("####")
	Title("###      [X] Tic-Tac-Toe [O]")
	Title("##")
	Title("###########################")
}

func OutputPlayersInfo(players [2]map[string]string) {
	Title("\n=============================")
	Title("Let's roll!")
	for _, p := range players {
		Title(fmt.Sprintf("%s: %s [%s]", p["name"], p["signName"], p["sign"]))
	}
	Title("=============================")
}

func Congratulations(name, signName, sign string) {
	Title("\n====================================")
	Title("Congratulations!")
	Title(fmt.Sprintf("Winner: %s %s [%s]", name, signName, sign))
	Title("====================================")
}

func DrawGame() {
	Title("\n=====================")
	Title("It's a draw!")
	Title("No winner this time!")
	Title("=====================")
}
