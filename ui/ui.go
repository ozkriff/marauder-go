// See LICENSE file for copyright and license details.

package ui

import (
	"github.com/banthar/Go-SDL/sdl"
	"github.com/ozkriff/marauder/game"
)

// UserInterface позволяет рисовать происходящее в игре
// и получать ввод от пользователя
//
type UserInterface struct {
	renderer     *Render
	inputHandler *InputHandler
}

// New это конструктор
//
// Принимает указатель на game.Core, т.к. надо же чего-то отображать.
//
func NewUserInterface(core *game.Core) *UserInterface {
	renderer := NewRender(core)
	inputHandler := NewInputHandler(renderer, core)
	self := &UserInterface{
		renderer:     renderer,
		inputHandler: inputHandler,
	}
	return self
}

func (self *UserInterface) mainloop() {
	// TODO: render??
	for !self.renderer.IsFinished {
		self.inputHandler.ProcessSDLEvents()
		self.renderer.Draw()
		sdl.Delay(90)
		// sdl.Delay(1)
	}
}

// Run starts game's main loop
func (self *UserInterface) Run() {
	self.mainloop()
}

// Close освобождает захапаные ресурсы
func (self *UserInterface) Close() {
	self.renderer.Close()
}
