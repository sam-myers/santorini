package player

import "github.com/sam-myers/santorini/board"

type Player interface {
	Select([]board.Board) board.Board
}
