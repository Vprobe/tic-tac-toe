package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	msg "github.com/Vprobe/tic-tac-toe/messages"
	"github.com/Vprobe/tic-tac-toe/player"
	field "github.com/Vprobe/tic-tac-toe/playingfield"
	"github.com/Vprobe/tic-tac-toe/settings"
	"github.com/Vprobe/tic-tac-toe/sign"
	"github.com/Vprobe/tic-tac-toe/validation"
)

func SelectMode(gs *settings.Settings) {
	for {
		r := bufio.NewReader(os.Stdin)
		msg.SelectGameMode()
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)

		if validation.IsEmptyString(s) {
			msg.Error("Value can not be empty")
			continue
		}

		i, _ := strconv.Atoi(s)
		if i <= 0 || i > len(gs.GetModes()) {
			msg.Error(fmt.Sprintf("Value must be in range from %s to %s",
				strconv.Itoa(1),
				strconv.Itoa(len(gs.GetModes()))))
			continue
		}

		msg.Success("Game mode is: " + gs.GetModes()[i-1])
		gs.SetSelectedMode(gs.GetModes()[i-1])
		break
	}
}

func SelectDifficulty(gs *settings.Settings) {
	for {
		r := bufio.NewReader(os.Stdin)
		msg.SelectDifficulty()
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)

		if validation.IsEmptyString(s) {
			msg.Error("Value can not be empty")
			continue
		}

		i, _ := strconv.Atoi(s)
		if i <= 0 || i > len(gs.GetDifficulties()) {
			msg.Error(fmt.Sprintf("Value must be in range from %s to %s",
				strconv.Itoa(1),
				strconv.Itoa(len(gs.GetDifficulties()))))
			continue
		}

		msg.Success("Game difficulty is: " + gs.GetDifficulties()[i-1])
		gs.SetSelectedDifficulty(gs.GetDifficulties()[i-1])
		break
	}
}

func SelectSide() int {
	var res int
	side := [2]string{sign.CrossVal, sign.NoughtVal}

	for {
		r := bufio.NewReader(os.Stdin)
		msg.SelectSide()
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)

		if validation.IsEmptyString(s) {
			msg.Error("Value can not be empty")
			continue
		}

		i, _ := strconv.Atoi(s)
		if validation.IsIntOutOfRange(i, sign.MinSide, sign.MaxSide) {
			msg.Error(fmt.Sprintf("Value must be in range from %s to %s",
				strconv.Itoa(sign.MinSide),
				strconv.Itoa(sign.MaxSide)))
			continue
		}

		res = i
		msg.Success("Chosen side is: [" + side[i-1] + "]")
		break
	}

	return res
}

func ReadPlayerInput(p *player.Player) {
	setInput(field.Col, p)
	setInput(field.Row, p)
}

func setInput(val string, p *player.Player) {
	for {
		r := bufio.NewReader(os.Stdin)
		msg.Info(fmt.Sprintf("%s [%s]: type %s number [from 1 to 3]:",
			sign.GetSignNameByNumber(p.GetSign()),
			sign.GetSignByNumber(p.GetSign()),
			strings.ToUpper(val)))
		msg.Choice()
		s, _ := r.ReadString('\n')
		s = strings.TrimSpace(s)

		if validation.IsEmptyString(s) {
			msg.Error("Value can not be empty")
			continue
		}

		i, _ := strconv.Atoi(s)
		if validation.IsIntOutOfRange(i, field.MinRowCol, field.MaxRowCol) {
			msg.Error(fmt.Sprintf("Value must be in range from %s to %s",
				strconv.Itoa(field.MinRowCol),
				strconv.Itoa(field.MaxRowCol)))
			continue
		}

		msg.Success(strings.ToUpper(val[:1]) + val[1:] + " set: " + s)
		if val == field.Col {
			p.SetColumn(i)
		} else {
			p.SetRow(i)
		}

		break
	}
}
