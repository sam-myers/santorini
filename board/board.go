package board


func NewBoard() Board {
	board := Board{}
	for y := 0; y < boardYLen; y++ {
		for x := 0; x < boardXLen; x++ {
			board[y][x].Coord.X = x
			board[y][x].Coord.Y = y
		}
	}
	return board
}

func (b *Board) Spaces() <-chan Space {
	ch := make(chan Space)
	go func() {
		for y := 0; y < boardYLen; y++ {
			for x := 0; x < boardXLen; x++ {
				ch <- b[y][x]
			}
		}
		close(ch)
	}()
	return ch
}

// Is the board represents a win condition, give
// the player number corresponding to the winning
// token and true. Otherwise, nil and false
func (b *Board) IsWin() (*int, bool) {
	for space := range b.Spaces() {
		if space.Height == 3 && space.Token != nil {
			return &space.Token.PlayerNum, true
		}
	}
	return nil, false
}

func (b *Board) SpaceAt(c Coord) *Space {
	if c.IsInvalid() {
		return nil
	}
	return &b[c.Y][c.X]
}

func (b *Board) AddToken(c Coord, playerNum int) {
	space := b.SpaceAt(c)
	space.Token = &Token{
		PlayerNum: playerNum,
	}
}

func (b *Board) TokenPossibleMoves(c Coord) []Coord {
	height := b.SpaceAt(c).Height
	coords := make([]Coord, 0)
	for _, adjCoord := range c.Adjacent() {
		adjHeight := b.SpaceAt(adjCoord).Height
		if adjHeight < height || adjHeight+1 == height {
			coords = append(coords, adjCoord)
		}
	}
	return coords
}

func (b *Board) TokenPossibleBuilds(c Coord) []Coord {
	coords := make([]Coord, 0)
	for _, adjCoord := range c.Adjacent() {
		space := b.SpaceAt(adjCoord)
		if space.Height < 3 && !space.Capped {
			coords = append(coords, adjCoord)
		}
	}
	return coords
}

func (b *Board) TokenMove(from Coord, to Coord) {
	fromSpace := b.SpaceAt(from)
	toSpace := b.SpaceAt(to)

	token := b.SpaceAt(from).Token
	fromSpace.Token = nil
	toSpace.Token = token
}

func (b *Board) TokenBuild(site Coord) {
	b.SpaceAt(site).Height += 1
}

func (b *Board) PlayerTokenLocations(playerNum int) []Coord {
	locs := make([]Coord, 0)
	for space := range b.Spaces() {
		if space.Token != nil && space.Token.PlayerNum == playerNum {
			locs = append(locs, space.Coord)
		}
	}
	return locs
}

func (b *Board) Clone() Board {
	var newSpace Space
	nb := NewBoard()
	for oldSpace := range b.Spaces() {
		newSpace = *nb.SpaceAt(oldSpace.Coord)
		newSpace.Height = oldSpace.Height
		newSpace.Capped = oldSpace.Capped
		newSpace.Token = oldSpace.Token
	}
	return nb
}

func (b *Board) PossibleMoves(playerNum int) []Board {
	var destinationCoord Coord
	var newBoard Board
	boards := make([]Board, 0)

	for _, startingCoord := range b.PlayerTokenLocations(playerNum) {
		for _, destinationCoord = range b.TokenPossibleMoves(startingCoord) {
			newBoard = b.Clone()
			newBoard.TokenMove(startingCoord, destinationCoord)
			boards = append(boards, newBoard)
		}
	}
	return boards
}

func (b *Board) PossibleBuilds(playerCoord Coord) []Board {
	boards := make([]Board, 0)
	for _, buildCoord := range b.TokenPossibleMoves(playerCoord) {
		newBoard := b.Clone()
		newBoard.TokenBuild(buildCoord)
		boards = append(boards, newBoard)
	}
	return boards
}

func (b *Board) PossiblePlys(playerNum int) []Board {
	var destinationCoord Coord
	var buildCoord Coord
	var newBoard Board
	boards := make([]Board, 0)

	for _, startingCoord := range b.PlayerTokenLocations(playerNum) {
		for _, destinationCoord = range b.TokenPossibleMoves(startingCoord) {
			for _, buildCoord = range b.TokenPossibleMoves(destinationCoord) {
				newBoard = b.Clone()
				newBoard.TokenMove(startingCoord, destinationCoord)
				newBoard.TokenBuild(buildCoord)
				boards = append(boards, newBoard)
			}
		}
	}
	return boards
}
