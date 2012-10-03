package gogol

import "time"

// DefaultHandler implements Handler
type DefaultHandler struct {
	WindowTitle string
}

func (d *DefaultHandler) Ready()                               {}
func (d *DefaultHandler) Reshape(w, h int)                     {}
func (d *DefaultHandler) Display(delta time.Duration)          {}
func (d *DefaultHandler) DisplayPost()                         {}
func (d *DefaultHandler) Keyboard(key Key, isDown bool)        {}
func (d *DefaultHandler) Motion(x, y int)                      {}
func (d *DefaultHandler) Mouse(b Mouse, isDown bool, x, y int) {}
func (d *DefaultHandler) Entry(e bool)                         {}
func (d *DefaultHandler) Title() string                        { return d.WindowTitle }
