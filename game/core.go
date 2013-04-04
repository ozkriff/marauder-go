// See LICENSE file for copyright and license details.

package game

import (
	"fmt"
)

// Core - стержень всей игры! :)
type Core struct {
	Gameboard     *Gameboard
	Units         map[int]*Unit
	SelectedUnit  *Unit
	unitIDCounter int
}

// NewCore ...
//
// Принимает указатель на Gameboard.
//
func NewCore(mapPtr *Gameboard) *Core {
	self := Core{
		Units: map[int]*Unit{
			0: &Unit{Pos: Pos{X: 1, Y: 2}},
			1: &Unit{Pos: Pos{X: 2, Y: 2}},
		},
		Gameboard:     mapPtr,
		unitIDCounter: 0,
	}
	return &self
}

type NoUnitAtThisPosError struct {
	Pos Pos
}

func (e NoUnitAtThisPosError) Error() string {
	return fmt.Sprintf("%v", e.Pos)
}

// UnitAt проверят, есть ли юнит по таким кординатам иили нет.
func (self *Core) UnitAt(pos Pos) (*Unit, error) {
	for _, unit := range self.Units {
		if unit.Pos == pos {
			return unit, nil
		}
	}
	return nil, NoUnitAtThisPosError{Pos: pos}
}

// GenerateNewUnitID создает новый id юнита.
//
// новый уникальный для этой сессии ID для юнита.
//
func (self *Core) GenerateNewUnitID() int {
	self.unitIDCounter++
	return self.unitIDCounter
}
