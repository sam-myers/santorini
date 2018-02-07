package player

import "github.com/sam-myers/santorini/board"

type DumbPlayer struct {}

func (p *DumbPlayer) Select(possibilities []board.Board) board.Board {
	return possibilities[0]
}
