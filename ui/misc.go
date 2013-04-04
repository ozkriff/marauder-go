// See LICENSE file for copyright and license details.

package ui

import (
	"github.com/banthar/Go-SDL/sdl"
	"math"
)

func loadImage(filename string) *sdl.Surface {
	image := sdl.Load(filename)
	if image == nil {
		panic(sdl.GetError())
	}
	return image
}

// Get distance between two screen points.
func distance(a, b ScreenPos) int {
	dx := math.Abs(float64(b.X - a.X))
	dy := math.Abs(float64(b.Y - a.Y))
	result := math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
	return int(result)
}

func isOdd(n int) bool {
	return (n % 2) != 0
}
