package main

import (
	"crayola"
	"crayola/image"
	"flag"
	"os"
	"time"
)

var (
	title = flag.String("title", "crayola", "window title")
)

// Speed of play
var speed float64 = 7

type Cray struct {
	width        int
	height       int
	seq          *image.Sequence
	title        string
	scale        float32
	isMouseLDown bool
	isMouseRDown bool
}

func (c *Cray) Ready() {
	c.seq.Play(speed)
	crayola.HideCursor()
	crayola.FullScreen()
}

func (c *Cray) Reshape(w, h int) {
	c.width, c.height = w, h
}

func (c *Cray) Display(delta time.Duration) {
	crayola.Scale(c.scale, c.scale)

	for x := 0; x <= c.width; x += 16 {
		for y := 0; y <= c.height; y += 16 {
			c.seq.DrawFrameAt(x, y)
		}
	}
}

func (c *Cray) Keyboard(key crayola.Key, isDown bool) {
	if !isDown {
		return
	}
	switch key {
	case '.':
		c.scale++
	case ',':
		c.scale--
	case crayola.KeySpace:
		c.seq.Toggle()

		if c.seq.IsPlaying {
			crayola.HideCursor()
		} else {
			crayola.ShowCursor()
		}
	case crayola.KeyEsc:
		os.Exit(0)
	}
}

func (c *Cray) Title() string {
	return c.title
}

func (c *Cray) Motion(x, y int) {
	s := float32(y) / float32(c.height) * 2
	l := float32(x) / float32(c.width) * 2

	if c.isMouseLDown {
		crayola.AdjustHSL(s-1, l-1, 0)
	}
	if c.isMouseRDown {
		crayola.AdjustExp(s, l)
	}
}

func (c *Cray) Mouse(b crayola.Mouse, isDown bool, x, y int) {
	switch b {
	case crayola.MouseL:
		c.isMouseLDown = isDown
	case crayola.MouseR:
		c.isMouseRDown = isDown
	}
}

func (c *Cray) Entry(e bool) {}

func main() {
	flag.Parse()

	img := crayola.NewImage("images/crayola.bmp")
	sprite := img.Sprite(16, 16)
	seq := sprite.Sequence(0, -1)

	handler := &Cray{seq: seq, title: *title, scale: 1}
	crayola.Init(handler)
}
