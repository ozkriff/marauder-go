// See LICENSE file for copyright and license details.

package game_test

import (
	"testing"
	"my/marauder/game"
)

func TestNew(t *testing.T) {
	p := game.Pos{X: 0, Y: 1}
	if p.X != 0 {
		t.Errorf("X != 0")
	}
	if p.Y != 1 {
		t.Errorf("Y != 1")
	}
}

func TestAdd(t *testing.T) {
	v1 := game.Pos{X: 1, Y: 2}
	v2 := game.Pos{X: 3, Y: 4}
	real := v1.Add(v2)
	expected := game.Pos{X: 4, Y: 6}
	if real != expected {
		t.Fail()
	}
}
