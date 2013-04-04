// See LICENSE file for copyright and license details.

package game

import (
	"fmt"
	"github.com/ozkriff/marauder/queue"
	"log"
)

// TODO: remove print funcs from here!
// ----------------------------------------------------

func PrintForEachTile(m Gameboard, f func(*Tile)) {
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

func PrintMapIsWalkable(m Gameboard) {
	PrintForEachTile(m, func(tile *Tile) {
		if tile.IsWalkable {
			fmt.Printf("..  ")
		} else {
			fmt.Printf("[]  ")
		}
	})
}

func PrintMapCost(m Gameboard) {
	PrintForEachTile(m, func(tile *Tile) {
		if tile.Cost == 200 {
			fmt.Printf(".   ")
		} else {
			fmt.Printf("%-2v  ", tile.Cost)
		}
	})
}

func PrintPath(m Gameboard) {
	PrintForEachTile(m, func(tile *Tile) {
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
	m *Gameboard,
	u *Unit,
	p Pos) Dir {

	tile := m.Tile(p)
	if tile.Cost == 0 {
		return u.Dir
	}
	return tile.Parent.Opposite()
}

func getTileCost(
	originalPos Pos,
	neighbourPos Pos) int {

	// int diff = Dir(t, nb).diff(getParentDir(u, t))
	// int maxAP = u.type().actionPoints - 1
	// int additionalCost[] = {3, 4, maxAP, maxAP}
	// assert(diff >= 0 && diff <= 3)
	// return 1 + additionalCost[diff]

	return 1 // TODO: convert full version from C++
}

func processNeighbourPos(
	m *Gameboard,
	originalPos Pos,
	neighbourPos Pos,
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
		parent, err := GetDirFromPosToPos(
			neighbourPos, originalPos)
		if err != nil {
			log.Fatal()
		}
		t2.Parent = parent
		// t2.dir = Dir(neighbourPos, originalPos)
	}
}

func tryToPushNeighbours(
	m *Gameboard,
	// const Unit& u,
	p Pos,
) {
	if !m.IsInboard(p) {
		log.Fatalf("p(%#v) isn't inboard", p)
	}
	for i := 0; i < 6; i++ {
		neighbourPos := GetNeighbourPos(p, Dir(i))
		if m.IsInboard(neighbourPos) {
			processNeighbourPos(m, p, neighbourPos)
		}
	}
}

func cleanMap(m *Gameboard) {
	m.ForEachTilePos(func(p Pos) {
		m.Tile(p).Cost = 200
	})
}

// Fill оценивает каждую проходимую клетку карты.
//
// Т.е. заполняет поле Cost в клетках.
//
func Fill(m *Gameboard, startPos Pos) {
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
		p := pathQueue.Pop().(Pos)
		tryToPushNeighbours(m, p)
	}
}

func GetPath(m *Gameboard, target Pos) []Pos {
	if !m.IsInboard(target) {
		log.Fatalf("bad target position %#v", target)
	}

	p := target

	// std::vector<V2i> path;
	path := make([]Pos, 0)

	for m.Tile(p).Cost != 0 {
		path = append(path, p) // path.push_back(p);
		p = GetNeighbourPos(p, m.Tile(p).Parent)
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
	reversedPath := make([]Pos, pathLen)
	for k, v := range path {
		reversedPath[pathLen-1-k] = v
	}

	return reversedPath
}
