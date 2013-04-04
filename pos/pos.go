// See LICENSE file for copyright and license details.

// Package pos - позиция на карте.
package pos

type Pos struct {
	X int
	Y int
}

func New(x, y int) Pos {
	return Pos{X: x, Y: y}
}

// Add one vector to another
func (self Pos) Add(other Pos) Pos {
	return Pos{self.X + other.X, self.Y + other.Y}
}

func (self Pos) Subtract(other Pos) Pos {
	return Pos{self.X - other.X, self.Y - other.Y}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}

func (self Pos) Distance(other Pos) int {
	var ax int = self.X + self.Y/2
	var bx int = other.X + other.Y/2
	var dx int = bx - ax
	var dy int = other.Y - self.Y
	return (abs(dx) + abs(dy) + abs(dx-dy)) / 2
}
