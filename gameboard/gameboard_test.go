// See LICENSE file for copyright and license details.

package gameboard_test

import (
	"my/marauder/gameboard"
	"my/marauder/pos"
	"testing"
)

func TestIsInboard(t *testing.T) {
	m := gameboard.New(pos.Pos{X: 3, Y: 4})
	if !m.IsInboard(pos.New(1, 1)) {
		t.Fail()
	}
	if !m.IsInboard(pos.New(0, 0)) {
		t.Fail()
	}
	if !m.IsInboard(pos.New(2, 3)) {
		t.Fail()
	}

	if m.IsInboard(pos.New(-1, 0)) {
		t.Fail()
	}
	if m.IsInboard(pos.New(0, -1)) {
		t.Fail()
	}
	if m.IsInboard(pos.New(3, 3)) {
		t.Fail()
	}
	if m.IsInboard(pos.New(2, 4)) {
		t.Fail()
	}
}
