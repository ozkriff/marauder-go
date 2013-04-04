// See LICENSE file for copyright and license details.

package ui

import (
	"github.com/banthar/Go-SDL/sdl"
	// "github.com/0xe2-0x9a-0x9b/Go-SDL/sdl" // ?
	"github.com/ozkriff/marauder/game"
)

// InputHandler обрабатывает ввод пользователя и всякие
// менее важные события, навроде изменения размера окна
type InputHandler struct {
	render *Render
	core   *game.Core
}

func (self *InputHandler) processSDLResizeEvent(event *sdl.ResizeEvent) {
	self.render.Resize(int(event.W), int(event.H))
}

func (self *InputHandler) processSDLKeyboardEvent(event *sdl.KeyboardEvent) {
	sym := event.Keysym.Sym
	println(sym, ": ", sdl.GetKeyName(sdl.Key(sym)))
	if event.Type == sdl.KEYDOWN {
		switch event.Keysym.Sym {
		case sdl.K_ESCAPE, sdl.K_q:
			self.render.IsFinished = true
		case sdl.K_UP:
			self.render.MapOffset.Y += 50
		case sdl.K_DOWN:
			self.render.MapOffset.Y -= 50
		case sdl.K_RIGHT:
			self.render.MapOffset.X -= 50
		case sdl.K_LEFT:
			self.render.MapOffset.X += 50
		}
	}
}

func (self *InputHandler) processSDLMouseMotionEvent(event *sdl.MouseMotionEvent) {
	self.render.MousePos = ScreenPos{
		X: int(event.X),
		Y: int(event.Y),
	}
	self.render.SelectedMapPos = self.render.PickTile(self.render.MousePos)
	// println(SelectedMapPos.X, SelectedMapPos.Y)
}

func (self *InputHandler) processSDLButtonEvent(event *sdl.MouseButtonEvent) {
	// TODO: err?
	unit, _ := self.core.UnitAt(self.render.SelectedMapPos)
	// fmt.Printf("%#v : %#v\n", unit, err)
	if unit != nil {
		self.core.SelectedUnit = unit
	} else if self.core.SelectedUnit != nil {
		self.core.SelectedUnit.Pos = self.render.SelectedMapPos
	}
}

func (self *InputHandler) processSDLEvent(abstractEvent sdl.Event) {
	switch event := abstractEvent.(type) {
	case *sdl.ResizeEvent:
		self.processSDLResizeEvent(event)
	case *sdl.KeyboardEvent:
		self.processSDLKeyboardEvent(event)
	case *sdl.MouseMotionEvent:
		self.processSDLMouseMotionEvent(event)
	case *sdl.MouseButtonEvent:
		self.processSDLButtonEvent(event)
	case *sdl.QuitEvent:
		self.render.IsFinished = true
	}
}

// processSDLEvents вытаскивает все новые события из SDL
// и обрабатывает их.
func (self *InputHandler) ProcessSDLEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		self.processSDLEvent(event)
	}
}

// New создает новый экземпляр InputHandler'а
func NewInputHandler(render *Render, core *game.Core) *InputHandler {
	self := InputHandler{
		render: render,
		core:   core,
	}
	return &self
}
