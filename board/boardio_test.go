package board

import (
	"bytes"
	"testing"
)

func TestBoard_ImportExport(t *testing.T) {
	buf := bytes.Buffer{}

	board := NewBoard()
	board.Export(&buf)
	reader := bytes.NewReader(buf.Bytes())

	newBoard, err := Import(reader)
	if newBoard == nil {
		t.Fatal("Didn't get a board back from import")
	}
	if err != nil {
		t.Fatalf("Error deserializing Board: %s", err)
	}
}

func TestBoard_Print(t *testing.T) {
	board := NewBoard()
	board[0][0].Height = 1
	board[0][1].Height = 2
	board[0][2].Height = 3
	board[0][2].Capped = true

	board[4][4].Height = 2
	board[4][4].Token = &Token{}

	expected := `[     ] [[   ]] [[[O]]]                 
                                        
                                        
                                        
                                [[ T ]] 
`
	actual := board.Print()

	if expected != actual {
		t.Fatalf("Output was supposed to be \n%s\n, was actually \n%s\n", expected, actual)
	}
}
