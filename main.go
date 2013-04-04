// See LICENSE file for copyright and license details.

// Marauder is 2D turn-based hexagonal strategy game.
//
package main

import (
	"fmt"
	"my/marauder/core"
	"my/marauder/gameboard"
	"my/marauder/path"
	"my/marauder/pos"
	userInterface "my/marauder/ui"
)

func setSomeUnwalkableTiles(m *gameboard.Gameboard) {
	coordinates := []pos.Pos{
		{X: 5, Y: 5},
		{X: 4, Y: 6},
		{X: 4, Y: 7},
		{X: 4, Y: 8},
		{X: 4, Y: 9},
		{X: 4, Y: 10},
		{X: 4, Y: 11},
		{X: 4, Y: 12},
		{X: 4, Y: 13},
		{X: 4, Y: 14},
		{X: 5, Y: 5},
		{X: 6, Y: 5},
		{X: 7, Y: 5},
		{X: 8, Y: 5},
		{X: 9, Y: 5},
	}
	for _, coord := range coordinates {
		m.Tile(coord).IsWalkable = false
	}
}

func main() {
	mapSize := pos.Pos{X: 10, Y: 20}
	board := gameboard.New(mapSize)
	setSomeUnwalkableTiles(board)
	if true {
		// тест путенахождения
		path.Fill(board, pos.New(1, 2))
		path.PrintMapIsWalkable(*board)
		path.PrintMapCost(*board)
		p := path.GetPath(board, pos.Pos{X: 9, Y: 9})
		fmt.Printf("%v\n", p)
		path.PrintPath(*board)
	} else {
		// нормальная игра
		core := core.New(board)
		ui := userInterface.New(core)
		ui.Run()
		ui.Close()
	}
}
