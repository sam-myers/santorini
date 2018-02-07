package board

import "testing"

func TestNewBoard(t *testing.T) {
	board := NewBoard()
	if len(board) != boardYLen {
		t.Fatalf("Grid is supposed to %v long in the Y dimension, is actually %v", boardYLen, len(board))
	}
	if len(board[0]) != boardXLen {
		t.Fatalf("Grid is supposed to be %v long in the X dimension, is actually %v", boardXLen, len(board))
	}

	for y := 0; y < boardYLen; y++ {
		for x := 0; x < boardXLen; x++ {
			cell := board[y][x]
			if cell.Token != nil {
				t.Fatalf("Board is supposed to be initialized with no tokens. Token present at cell %+v", cell)
			}
			if cell.Coord.X != x {
				t.Fatalf("Coordinate of cell %+v's X is supposed to be %v, is %v", cell, x, cell.Coord.X)
			}
			if cell.Coord.Y != y {
				t.Fatalf("Coordinate of cell %+v's Y is supposed to be %v, is %v", cell, y, cell.Coord.Y)
			}
			if cell.Coord.IsInvalid() {
				t.Fatalf("Coord at cell %+v is invalid", cell)
			}
		}
	}
}

func TestBoard_Spaces(t *testing.T) {
	board := NewBoard()
	count := 0
	for range board.Spaces() {
		count++
	}
	if count != boardXLen*boardYLen {
		t.Fatalf("Iterator is supposed to return %vx%v, instead gave %v", boardXLen, boardYLen, count)
	}
}

func TestBoard_IsWin_Win(t *testing.T) {
	board := NewBoard()
	board[0][0].Token = &Token{PlayerNum: 1}
	board[0][0].Height = 3
	whoWon, gameOver := board.IsWin()
	if whoWon == nil || *whoWon != 1 || gameOver != true {
		t.Fatal("Game should be over")
	}
}

func TestBoard_IsWin_NoWin(t *testing.T) {
	board := NewBoard()
	board[0][0].Token = &Token{PlayerNum: 1}
	board[0][0].Height = 2
	whoWon, gameOver := board.IsWin()
	if whoWon != nil || gameOver != false {
		t.Fatal("Game shouldn't be over")
	}
}
