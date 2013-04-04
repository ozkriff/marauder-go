// See LICENSE file for copyright and license details.

package game_test

import (
	"testing"
	"github.com/ozkriff/marauder/game"
)

func TestIsInboard(t *testing.T) {
	m := game.NewGameboard(game.Pos{X: 3, Y: 4})
	if !m.IsInboard(game.Pos{X: 1, Y: 1}) {
		t.Fail()
	}
	if !m.IsInboard(game.Pos{X: 0, Y: 0}) {
		t.Fail()
	}
	if !m.IsInboard(game.Pos{X: 2, Y: 3}) {
		t.Fail()
	}

	if m.IsInboard(game.Pos{X: -1, Y: 0}) {
		t.Fail()
	}
	if m.IsInboard(game.Pos{X: 0, Y: -1}) {
		t.Fail()
	}
	if m.IsInboard(game.Pos{X: 3, Y: 3}) {
		t.Fail()
	}
	if m.IsInboard(game.Pos{X: 2, Y: 4}) {
		t.Fail()
	}
}
