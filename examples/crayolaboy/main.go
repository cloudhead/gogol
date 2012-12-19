package main

import (
	"flag"
	"gogol"
	"gogol/image"
	"os"
	"time"
)

var (
	title = flag.String("title", "gogol", "window title")
)

// Speed of play
var speed float64 = 7

type Gogol struct {
	gogol.Handler

	width        int
	height       int
	seq          *image.Sequence
	title        string
	scale        float32
	isMouseLDown bool
	isMouseRDown bool
}

func (c *Gogol) Ready() {
	img := gogol.NewImage("images/gogol.tga")
	sprite := img.Sprite(16, 16)
	c.seq = sprite.Sequence(0, -1)
	c.seq.Play(speed)

	gogol.HideCursor()
}

func (c *Gogol) Reshape(w, h int) {
	c.width, c.height = w, h
}

func (c *Gogol) Display(delta time.Duration) {
	gogol.Scale(c.scale, c.scale)

	for x := 0; x <= c.width; x += 16 {
		for y := 0; y <= c.height; y += 16 {
			c.seq.DrawFrameAt(x, y)
		}
	}
}

func (c *Gogol) Keyboard(key gogol.Key, isDown bool) {
	if !isDown {
		return
	}
	switch key {
	case '.':
		c.scale++
	case ',':
		c.scale--
	case gogol.KeySpace:
		c.seq.Toggle()

		if c.seq.IsPlaying {
			gogol.HideCursor()
		} else {
			gogol.ShowCursor()
		}
	case gogol.KeyEsc:
		os.Exit(0)
	}
}

func (c *Gogol) Title() string {
	return c.title
}

func (c *Gogol) Motion(x, y int) {
	s := float32(y) / float32(c.height) * 2
	l := float32(x) / float32(c.width) * 2

	if c.isMouseLDown {
		gogol.AdjustHSL(s-1, l-1, 0)
	}
	if c.isMouseRDown {
		gogol.AdjustExp(s, l)
	}
}

func (c *Gogol) Mouse(b gogol.Mouse, isDown bool) {
	switch b {
	case gogol.MouseL:
		c.isMouseLDown = isDown
	case gogol.MouseR:
		c.isMouseRDown = isDown
	}
}

func main() {
	flag.Parse()

	handler := &Gogol{
		Handler: &gogol.DefaultHandler{*title},
		scale:   1,
	}
	gogol.Init(handler)
}
