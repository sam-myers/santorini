package board

import "testing"

func TestCoord_Adjacent_Middle(t *testing.T) {
	c := Coord{X: boardXLen / 2, Y: boardYLen / 2}
	numAdjacent := len(c.Adjacent())
	if numAdjacent != 8 {
		t.Fatalf("There 8 spaces adjacent to a center tile, not %v", numAdjacent)
	}
}

func TestCoord_Adjacent_Edge(t *testing.T) {
	c := Coord{X: 0, Y: boardYLen / 2}
	numAdjacent := len(c.Adjacent())
	if numAdjacent != 5 {
		t.Fatalf("There 5 spaces adjacent to an edge tile, not %v", numAdjacent)
	}
}

func TestCoord_Adjacent_Corner(t *testing.T) {
	c := Coord{X: boardXLen - 1, Y: boardYLen - 1}
	numAdjacent := len(c.Adjacent())
	if numAdjacent != 3 {
		t.Fatalf("There 3 spaces adjacent to a corner tile, not %v", numAdjacent)
	}
}

func TestCoord_IsValid(t *testing.T) {
	invalid := Coord{X: -1, Y: -1}
	if invalid.IsValid() {
		t.Fatal("Coord is invalid")
	}
}

func TestCoord_IsInvalid(t *testing.T) {
	valid := Coord{X: 0, Y: 0}
	if valid.IsInvalid() {
		t.Fatal("Coord is invalid")
	}
}
