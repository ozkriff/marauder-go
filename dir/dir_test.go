// See LICENSE file for copyright and license details.

package dir_test

import (
	"my/marauder/dir"
	"my/marauder/pos"
	"testing"
)

func TestDiff(t *testing.T) {
	var d1 dir.Dir = 1
	var d2 dir.Dir = 3
	if d1.Diff(d2) != 2 {
		t.Fail()
	}
}

func TestOpposite(t *testing.T) {
	var d dir.Dir = 1
	if d.Opposite() != 1+3 {
		t.Fail()
	}
}

func TestGetDir(t *testing.T) {
	from := pos.Pos{X: 0, Y: 0}
	to := pos.Pos{X: 1, Y: 0}
	dir1, err := dir.GetDirFromPosToPos(from, to)
	if err != nil {
		t.Errorf("returned error!")
	}
	if dir1 != 1 {
		t.Errorf("bad dir")
	}
}

func GetDirFromPosToPos(t *testing.T) {
	from := pos.Pos{X: 0, Y: 0}
	to := pos.Pos{X: 2, Y: 0}
	_, err := dir.GetDirFromPosToPos(from, to)
	if err == nil {
		t.Errorf("where is my error?")
	}
}
