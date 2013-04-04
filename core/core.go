// See LICENSE file for copyright and license details.

package core

import (
	"fmt"
	"my/marauder/gameboard"
	"my/marauder/pos"
)

// Core - стержень всей игры! :)
type Core struct {
	Gameboard     *gameboard.Gameboard
	Units         map[int]*Unit
	SelectedUnit  *Unit
	unitIDCounter int
}

// New создает новую карту.
//
// Принимает указатель на Gameboard.
//
func New(mapPtr *gameboard.Gameboard) *Core {
	self := Core{
		Units: map[int]*Unit{
			0: &Unit{Pos: pos.New(1, 2)},
			1: &Unit{Pos: pos.New(2, 2)},
		},
		Gameboard:     mapPtr,
		unitIDCounter: 0,
	}
	return &self
}

type NoUnitAtThisPosError struct {
	Pos pos.Pos
}

func (e NoUnitAtThisPosError) Error() string {
	return fmt.Sprintf("%v", e.Pos)
}

// UnitAt проверят, есть ли юнит по таким кординатам иили нет.
func (self *Core) UnitAt(pos pos.Pos) (*Unit, error) {
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
