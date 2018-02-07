package board

func (c *Coord) Adjacent() []Coord {
	adjacent := make([]Coord, 0)
	var newX, newY int
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			newX = c.X + i
			newY = c.Y + j
			if i == 0 && j == 0 {
				continue
			}
			if newX < 0 || newX > boardXLen-1 {
				continue
			}
			if newY < 0 || newY > boardYLen-1 {
				continue
			}

			adjacent = append(adjacent, Coord{
				X: newX,
				Y: newY,
			})
		}
	}
	return adjacent
}

func (c *Coord) IsValid() bool {
	return c.X >= 0 && c.X < boardXLen && c.Y >= 0 && c.Y < boardYLen
}

func (c *Coord) IsInvalid() bool {
	return !c.IsValid()
}
