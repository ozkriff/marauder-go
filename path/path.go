// See LICENSE file for copyright and license details.

// Package path реализует путенахождение
//
package path

import (
	"fmt"
	"log"
	"my/marauder/core"
	"my/marauder/dir"
	"my/marauder/gameboard"
	"my/marauder/pos"
	`my/marauder/queue`
)

// TODO: remove print funcs from here!
// ----------------------------------------------------

func PrintForEachTile(
	m gameboard.Gameboard,
	f func(*gameboard.Tile),
) {
	for y := range m.Tiles {
		if y%2 == 0 {
			fmt.Printf("  ")
		}
		for x := range m.Tiles[y] {
			f(&m.Tiles[y][x])
		}
		fmt.Printf("\n")
	}
}

func PrintMapIsWalkable(m gameboard.Gameboard) {
	PrintForEachTile(m, func(tile *gameboard.Tile) {
		if tile.IsWalkable {
			fmt.Printf("..  ")
		} else {
			fmt.Printf("[]  ")
		}
	})
}

func PrintMapCost(m gameboard.Gameboard) {
	PrintForEachTile(m, func(tile *gameboard.Tile) {
		if tile.Cost == 200 {
			fmt.Printf(".   ")
		} else {
			fmt.Printf("%-2v  ", tile.Cost)
		}
	})
}

func PrintPath(m gameboard.Gameboard) {
	PrintForEachTile(m, func(tile *gameboard.Tile) {
		if tile.Visited {
			fmt.Printf("[]  ")
		} else {
			fmt.Printf("..  ")
		}
	})
}

// ----------------------------------------------------

var pathQueue = queue.Queue{}

// If this is first(start) tile - no parent tile
func getParentDir(
	m *gameboard.Gameboard,
	u *core.Unit,
	p pos.Pos) dir.Dir {

	tile := m.Tile(p)
	if tile.Cost == 0 {
		return u.Dir
	}
	return tile.Parent.Opposite()
}

func getTileCost(
	originalPos pos.Pos,
	neighbourPos pos.Pos) int {

	// int diff = Dir(t, nb).diff(getParentDir(u, t))
	// int maxAP = u.type().actionPoints - 1
	// int additionalCost[] = {3, 4, maxAP, maxAP}
	// assert(diff >= 0 && diff <= 3)
	// return 1 + additionalCost[diff]

	return 1 // TODO: convert full version from C++
}

func processNeighbourPos(
	m *gameboard.Gameboard,
	originalPos pos.Pos,
	neighbourPos pos.Pos,
) {
	t1 := m.Tile(originalPos)
	t2 := m.Tile(neighbourPos)

	// if (mCore.isUnitAt(neighbourPos) || t2.obstacle) {
	if !t2.IsWalkable {
		return
	}

	newcost := t1.Cost + getTileCost(originalPos, neighbourPos)

	// int ap = u.actionPoints()

	if t2.Cost > newcost /* && newcost <= ap */ {
		pathQueue.Push(neighbourPos)

		// update neighbour tile info
		t2.Cost = newcost
		parent, err := dir.GetDirFromPosToPos(
			neighbourPos, originalPos)
		if err != nil {
			log.Fatal()
		}
		t2.Parent = parent
		// t2.dir = Dir(neighbourPos, originalPos)
	}
}

func tryToPushNeighbours(
	m *gameboard.Gameboard,
	// const Unit& u,
	p pos.Pos,
) {
	if !m.IsInboard(p) {
		log.Fatalf("p(%#v) isn't inboard", p)
	}
	for i := 0; i < 6; i++ {
		neighbourPos := dir.GetNeighbourPos(p, dir.Dir(i))
		if m.IsInboard(neighbourPos) {
			processNeighbourPos(m, p, neighbourPos)
		}
	}
}

func cleanMap(m *gameboard.Gameboard) {
	m.ForEachTilePos(func(p pos.Pos) {
		m.Tile(p).Cost = 200
	})
}

// Fill оценивает каждую проходимую клетку карты.
//
// Т.е. заполняет поле Cost в клетках.
//
func Fill(m *gameboard.Gameboard, startPos pos.Pos) {
	if !pathQueue.IsEmpty() {
		log.Fatal("queue is not empty\n")
	}

	cleanMap(m)

	// Push start position
	pathQueue.Push(startPos)

	tile := m.Tile(startPos)
	tile.Cost = 0
	// TODO: t.parent = DirID::NONE
	tile.Parent = 0
	// TODO: ... t.dir = u.dir()

	for !pathQueue.IsEmpty() {
		// fmt.Printf("len: %d\n", pathQueue.Len())
		p := pathQueue.Pop().(pos.Pos)
		tryToPushNeighbours(m, p)
	}
}

func GetPath(m *gameboard.Gameboard, target pos.Pos) []pos.Pos {
	if !m.IsInboard(target) {
		log.Fatalf("bad target position %#v", target)
	}

	p := target

	// std::vector<V2i> path;
	path := make([]pos.Pos, 0)

	for m.Tile(p).Cost != 0 {
		path = append(path, p) // path.push_back(p);
		p = dir.GetNeighbourPos(p, m.Tile(p).Parent)
		if !m.IsInboard(p) {
			log.Fatalf("bad position %#v", p)
		}
		m.Tile(p).Visited = true
	}

	// Add start position
	path = append(path, p)

	m.Tile(p).Visited = true

	// reverse list
	// std::reverse(path.begin(), path.end());
	pathLen := len(path)
	reversedPath := make([]pos.Pos, pathLen)
	for k, v := range path {
		reversedPath[pathLen-1-k] = v
	}

	return reversedPath
}
