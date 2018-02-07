package board

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func Import(reader io.Reader) (*Board, error) {
	boardBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("Error reading from IO: %s", err)
		return nil, err
	}

	board := &Board{}
	err = json.Unmarshal(boardBytes, board)
	if err != nil {
		log.Printf("Error deserializing from JSON: %s", err)
	}
	return board, nil
}

func (b *Board) Export(writer io.Writer) error {
	boardBytes, err := json.Marshal(b)
	if err != nil {
		log.Printf("Error serializing Board to JSON: %s", err)
		return err
	}
	_, err = writer.Write(boardBytes)
	if err != nil {
		log.Printf("Error writing JSON to %s", err)
		return err
	}
	return nil
}

func (b *Board) Print() string {
	var groundLeft string
	var groundRight string
	var token string
	var space Space
	separator := " "

	buf := bytes.Buffer{}
	for y := 0; y < boardYLen; y++ {
		for x := 0; x < boardXLen; x++ {
			space = b[y][x]

			switch {
			case space.Capped:
				token = "O"
			case space.Token != nil:
				token = "T"
			default:
				token = " "
			}

			switch {
			case space.Height == 0:
				groundLeft = "   "
				groundRight = "   "
			case space.Height == 1:
				groundLeft = "[  "
				groundRight = "  ]"
			case space.Height == 2:
				groundLeft = "[[ "
				groundRight = " ]]"
			case space.Height == 3:
				groundLeft = "[[["
				groundRight = "]]]"
			}

			buf.WriteString(groundLeft + token + groundRight + separator)
		}
		buf.WriteString("\n")
	}
	return buf.String()
}
