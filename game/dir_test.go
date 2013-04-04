// See LICENSE file for copyright and license details.

package game_test

import (
	"testing"
	"my/marauder/game"
)

func TestDiff(t *testing.T) {
	var d1 game.Dir = 1
	var d2 game.Dir = 3
	if d1.Diff(d2) != 2 {
		t.Fail()
	}
}

func TestOpposite(t *testing.T) {
	var d game.Dir = 1
	if d.Opposite() != 1+3 {
		t.Fail()
	}
}

func TestGetDir(t *testing.T) {
	from := game.Pos{X: 0, Y: 0}
	to := game.Pos{X: 1, Y: 0}
	dir1, err := game.GetDirFromPosToPos(from, to)
	if err != nil {
		t.Errorf("returned error!")
	}
	if dir1 != 1 {
		t.Errorf("bad dir")
	}
}

func GetDirFromPosToPos(t *testing.T) {
	from := game.Pos{X: 0, Y: 0}
	to := game.Pos{X: 2, Y: 0}
	_, err := game.GetDirFromPosToPos(from, to)
	if err == nil {
		t.Errorf("where is my error?")
	}
}
