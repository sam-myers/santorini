package main

import (
	"github.com/sam-myers/santorini/board"
	"github.com/sam-myers/santorini/player"
	"fmt"
)

func main() {
	game := board.NewBoard()
	game.AddToken(board.Coord{X:0, Y:0}, 0)
	game.AddToken(board.Coord{X:1, Y:0}, 0)
	game.AddToken(board.Coord{X:0, Y:3}, 1)
	game.AddToken(board.Coord{X:2, Y:3}, 1)

	p0 := player.DumbPlayer{}
	p1 := player.DumbPlayer{}

	for i:=0; i<10; i++ {
		game = p0.Select(game.PossiblePlys(0))
		fmt.Print(game.Print())
		game = p1.Select(game.PossiblePlys(1))
		fmt.Print(game.Print())
	}
}
