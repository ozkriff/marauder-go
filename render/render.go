// See LICENSE file for copyright and license details.

package render

import (
	"github.com/banthar/Go-SDL/sdl"
	"my/marauder/core"
	"my/marauder/pos"
)

// Некоторая позиция в пикселях относительно левого верхнего края окна.
type ScreenPos struct {
	X int
	Y int
}

// Render, вы не поверите, рисует
type Render struct {
	// ядро - всему голова
	Core *core.Core

	// сдвиг левого верхнего края карты
	// относительно левого верхнего края окна
	MapOffset ScreenPos

	// закончить главный цикл?
	IsFinished bool

	// показывается в окне
	screen *sdl.Surface

	imageTile         *sdl.Surface
	imageSelectedTile *sdl.Surface
	imageUnit         *sdl.Surface

	// текущие координаты курсора мыши в пикселях
	MousePos ScreenPos

	// текущая выбранная позиция на карте
	SelectedMapPos pos.Pos
}

func (self *Render) mapPosToScreenPos(mapPos pos.Pos) ScreenPos {
	const ImageSize = 96
	// TODO: 92??
	screenPos := ScreenPos{
		// X: self.MapOffset.X + mapPos.X*(ImageSize) + 92/2,
		X: self.MapOffset.X + mapPos.X*(ImageSize+1) + 96/2,
		Y: self.MapOffset.Y + mapPos.Y*(ImageSize*3/4+1) + ImageSize/2,
	}
	if isOdd(mapPos.Y) {
		screenPos.X -= ImageSize / 2
	}
	return screenPos
}

// PickTile принимает экранные координаты,
// а возвращает соответствующую позицию на карте
//
func (self *Render) PickTile(MousePos ScreenPos) pos.Pos {
	var closestMapPos pos.Pos
	minimalDistance := 9000
	self.Core.Gameboard.ForEachTilePos(func(position pos.Pos) {
		screenPos := self.mapPosToScreenPos(position)
		if distance(MousePos, screenPos) < minimalDistance {
			minimalDistance = distance(MousePos, screenPos)
			closestMapPos = position
		}
	})
	return closestMapPos
}

func (self *Render) drawImageAt(image *sdl.Surface, position ScreenPos) {
	rect := sdl.Rect{
		X: int16(position.X - int(image.W)/2),
		Y: int16(position.Y - int(image.H)/2),
		W: 0,
		H: 0}
	self.screen.Blit(&rect, image, nil)
}

func (self *Render) drawUnitAt(image *sdl.Surface, position pos.Pos) {
	screenPos := self.mapPosToScreenPos(position)
	// screenPos1 := screenPos.Add(ScreenPos{15, 15})
	screenPos1 := screenPos
	screenPos1.X += 15
	screenPos1.Y += 15
	screenPos2 := screenPos
	screenPos2.X -= 15
	screenPos2.Y += 15
	screenPos3 := screenPos
	screenPos3.X -= 15
	screenPos3.Y -= 15
	screenPos4 := screenPos
	screenPos4.X += 15
	screenPos4.Y -= 15
	self.drawImageAt(self.imageUnit, screenPos1)
	self.drawImageAt(self.imageUnit, screenPos2)
	self.drawImageAt(self.imageUnit, screenPos3)
	self.drawImageAt(self.imageUnit, screenPos4)
}

// Draw выполняет непосредственно отрисовку
func (self *Render) Draw() {
	self.screen.FillRect(nil, 0x000000)
	self.Core.Gameboard.ForEachTilePos(func(position pos.Pos) {
		self.drawImageAt(self.imageTile,
			self.mapPosToScreenPos(position))
	})
	for _, unit := range self.Core.Units {
		self.drawUnitAt(self.imageUnit, unit.Pos)
	}
	if self.Core.SelectedUnit != nil {
		self.drawImageAt(self.imageSelectedTile,
			self.mapPosToScreenPos(self.Core.SelectedUnit.Pos))
	}
	self.drawImageAt(self.imageSelectedTile,
		self.mapPosToScreenPos(self.SelectedMapPos))
	self.screen.Flip()
}

func (self *Render) setVideoMode(w, h int) {
	self.screen = sdl.SetVideoMode(w, h, 32, sdl.RESIZABLE)
	if self.screen != nil {
		// self.reshape(w, h)
		// TODO: update?
	} else {
		panic("Couldn't set ### GL video mode: " + sdl.GetError() + "\n")
	}
}

// Resize изменяет размеры поверхности screen.
func (self *Render) Resize(w, h int) {
	self.screen.Free()
	self.screen = sdl.SetVideoMode(w, h, 32, sdl.RESIZABLE)
}

// New создает новый экземпляр Render.
//
// Принимает указатель на ядро.
//
func New(core *core.Core) *Render {
	initSDL()
	self := Render{
		Core:              core,
		IsFinished:        false,
		MapOffset:         ScreenPos{X: 10, Y: 10},
		imageTile:         loadImage("img/rocks.png"),
		imageSelectedTile: loadImage("img/selectedTile.png"),
		imageUnit:         loadImage("img/rifleSquad.png"),
	}
	self.setVideoMode(640, 480)
	return &self
}

func initSDL() {
	sdl.Init(sdl.INIT_EVERYTHING)
	sdl.WM_SetCaption("Marauder", "marauder")
}

// Close освобождает ресурсы
func (self *Render) Close() {
	sdl.Quit()
}
