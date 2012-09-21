package image

import "image"

type Clip struct {
	image.Rectangle
	Image *Image
}

func NewClip(img *Image, x, y, w, h int) *Clip {
	c := &Clip{
		Image:     img,
		Rectangle: image.Rect(x, y, w, h),
	}
	return c
}

func (c *Clip) Draw() {
	c.DrawAt(0, 0)
}

func (c *Clip) DrawAt(x, y int) {
	c.Image.DrawRectangleAt(c.Rectangle, x, y)
}
