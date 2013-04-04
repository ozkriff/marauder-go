// See LICENSE file for copyright and license details.

package gameboard

import (
	"my/marauder/dir"
	"my/marauder/pos"
)

type Tile struct {
	Cost       int
	IsWalkable bool
	Parent     dir.Dir
	Visited    bool
}

type Gameboard struct {
	Size  pos.Pos
	Tiles [][]Tile
}

// New создает новый экземпляр карты с заданным размером
func New(size pos.Pos) *Gameboard {
	self := Gameboard{
		Size: size,
	}
	self.Tiles = make([][]Tile, size.Y)
	for y := 0; y < size.Y; y++ {
		self.Tiles[y] = make([]Tile, size.X)
		for x := 0; x < size.X; x++ {
			self.Tiles[y][x] = Tile{
				Cost:       0,
				IsWalkable: true,
			}
		}
	}
	return &self
}

func (self *Gameboard) Tile(p pos.Pos) *Tile {
	return &self.Tiles[p.Y][p.X]
}

func (self *Gameboard) IsInboard(p pos.Pos) bool {
	return p.X >= 0 &&
		p.X < self.Size.X &&
		p.Y >= 0 &&
		p.Y < self.Size.Y
}

// ForEachTilePos выполняет некую функцию для каждой клетки,
// передавая ее координаты в функцию.
func (self *Gameboard) ForEachTilePos(f func(position pos.Pos)) {
	for y := 0; y < self.Size.Y; y++ {
		for x := 0; x < self.Size.X; x++ {
			f(pos.New(x, y))
		}
	}
}
