// See LICENSE file for copyright and license details.

package pos_test

import (
	"my/marauder/pos"
	"testing"
)

func TestNew(t *testing.T) {
	p := pos.New(0, 1)
	if p.X != 0 {
		t.Errorf("X != 0")
	}
	if p.Y != 1 {
		t.Errorf("Y != 1")
	}
}

func TestAdd(t *testing.T) {
	v1 := pos.Pos{X: 1, Y: 2}
	v2 := pos.Pos{X: 3, Y: 4}
	real := v1.Add(v2)
	expected := pos.Pos{X: 4, Y: 6}
	if real != expected {
		t.Fail()
	}
}
