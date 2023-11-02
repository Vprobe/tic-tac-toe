package player

import (
	"fmt"
	"strconv"

	field "github.com/Vprobe/tic-tac-toe/playingfield"
	"github.com/Vprobe/tic-tac-toe/validation"
)

type Player struct {
	name   string
	row    int
	column int
	sign   int
	ai     bool
}

func NewPlayer(name string, sign int, ai bool) *Player {
	if ai {
		name += " [AI]"
	} else {
		name += " [Human]"
	}
	return &Player{name: name, sign: sign, ai: ai, row: 1, column: 1}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetRow() int {
	return p.row
}

func (p *Player) GetColumn() int {
	return p.column
}

func (p *Player) GetSign() int {
	return p.sign
}

func (p *Player) IsAI() bool {
	return p.ai
}

func (p *Player) SetRow(r int) {
	if validation.IsIntOutOfRange(r, field.MinRowCol, field.MaxRowCol) {
		panic(fmt.Sprintf("Player: Row must be in range from %s to %s",
			strconv.Itoa(field.MinRowCol),
			strconv.Itoa(field.MaxRowCol)))
	}
	p.row = r
}

func (p *Player) SetColumn(c int) {
	if validation.IsIntOutOfRange(c, field.MinRowCol, field.MaxRowCol) {
		panic(fmt.Sprintf("Player: Column must be in range from %s to %s",
			strconv.Itoa(field.MinRowCol),
			strconv.Itoa(field.MaxRowCol)))
	}
	p.column = c
}
