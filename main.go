// See LICENSE file for copyright and license details.

// Marauder is 2D turn-based hexagonal strategy game.
//
package main

import (
	"fmt"
	"github.com/ozkriff/marauder/game"
	userInterface "github.com/ozkriff/marauder/ui"
)

func setSomeUnwalkableTiles(m *game.Gameboard) {
	coordinates := []game.Pos{
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
	mapSize := game.Pos{X: 10, Y: 20}
	board := game.NewGameboard(mapSize)
	setSomeUnwalkableTiles(board)
	if true {
		// тест путенахождения
		pathfinder := game.Pathfinder{
			Board: board,
		}
		pathfinder.Fill(game.Pos{X: 1, Y: 2})
		game.PrintMapIsWalkable(*board)
		game.PrintMapCost(*board)
		p := pathfinder.GetPath(game.Pos{X: 9, Y: 9})
		fmt.Printf("%v\n", p)
		game.PrintPath(*board)
	} else {
		// нормальная игра
		core := game.NewCore(board)
		ui := userInterface.NewUserInterface(core)
		ui.Run()
		ui.Close()
	}
}
