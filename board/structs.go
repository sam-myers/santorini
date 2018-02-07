package board

const boardXLen = 5
const boardYLen = 5

type boardRow [boardXLen]Space
type Board [boardYLen]boardRow

type Space struct {
	Token  *Token
	Coord  Coord
	Height int
	Capped bool
}

type Coord struct {
	X int
	Y int
}

type Token struct {
	PlayerNum int
}
