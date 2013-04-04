// See LICENSE file for copyright and license details.

package core

import (
	"my/marauder/dir"
	"my/marauder/pos"
)

// Unit это один игровой отряд
type Unit struct {
	Pos pos.Pos
	Dir dir.Dir
}
