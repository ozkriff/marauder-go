// See LICENSE file for copyright and license details.

// Package dir provides hexagonal directions.
//
// Модуль direction реализует работу с направлениями 
// в гексагональной сетке.
//
package game

import (
	"errors"
	"log"
)

// Dir обозначает некоторое направления.
// Например, движение в определенную сторону
// или направление взгляда юнита.
//
type Dir int

// Константы направлений.
const (
	NorthEast = iota
	East
	SouthEast
	SouthWest
	West
	NorthWest
	None
)

// dirToPosDiff это вспомогательная таблица
// для преобразования из позиции в направление.
//
var dirToPosDiff = [2][6]Pos{
	{
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 0},
		{0, -1},
	},
	{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
	},
}

// Diff returns difference between this and some other direction.
func (self Dir) Diff(other Dir) int {
	d := self - other
	if d < 0 {
		d = -d
	}
	if d > 6/2 {
		d = 6 - d
	}
	return int(d)
}

// Opposite return opposite direction.
func (self Dir) Opposite() Dir {
	directionIndex := self + 6/2
	if directionIndex >= 6 {
		directionIndex -= 6
	}
	return directionIndex
}

func getTableIndex(p Pos) int {
	var isOddRow bool = (p.Y%2 == 1)
	var subtableIndex int
	if isOddRow {
		subtableIndex = 1
	} else {
		subtableIndex = 0
	}
	return subtableIndex
}

// GetNeighbourPos возвращает соседнюю позицию в определенном направлении.
func GetNeighbourPos(p Pos, i Dir) Pos {
	subtableIndex := getTableIndex(p)
	if i >= 6 {
		log.Fatal("bad direction")
	}
	difference := dirToPosDiff[subtableIndex][i]
	return p.Add(difference)
}

// GetDirFromPosToPos принимает две прилежащих
// позиции и возвращает индекс направления.
//
func GetDirFromPosToPos(
	a Pos, b Pos,
) (Dir, error) {
	if a.Distance(b) != 1 {
		return 0, errors.New("distance != 1")
	}
	diff := b.Subtract(a)
	for i := 0; i < 6; i++ {
		if diff == dirToPosDiff[a.Y%2][i] {
			return Dir(i), nil
		}
	}
	return 0, errors.New("bad positions")
}
