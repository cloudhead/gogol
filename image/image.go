package image

//
// #include <GL/gl.h>
// #include "texture.h"
// #cgo LDFLAGS: -lglut -lGL
//
import "C"

import (
	"image"
	"log"
	"os"
	"unsafe"
)

// Register bmp format
import _ "code.google.com/p/go.image/bmp"

type Image struct {
	image.Image
	FilePath string
	W, H     int
	id       C.uint
}

var cache = map[string]*Image{}

// New initializes a new image from the given file path.
func New(path string) *Image {
	if cache[path] != nil {
		return cache[path]
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := m.(*image.RGBA); !ok {
		log.Fatalf("image %s: invalid format, must be RGBA", path)
	}
	bounds := m.Bounds()

	img := &Image{
		Image:    m,
		FilePath: path,
		W:        bounds.Max.X,
		H:        bounds.Max.Y,
	}
	cache[path] = img

	return img
}

func (img *Image) Gen() {
	ptr := unsafe.Pointer(&img.Image.(*image.RGBA).Pix[0])
	img.id = C.textureGen(C.int(img.W), C.int(img.H), (*C.uint)(ptr))
}

// Draw is DrawAt(0, 0)
func (img *Image) Draw() {
	img.DrawAt(0, 0)
}

// DrawAt draws the image at the specified position.
func (img *Image) DrawAt(x, y int) {
	img.DrawRectAt(0, 0, img.W, img.H, x, y)
}

// DrawRectangle is DrawRectangleAt with a zero position.
func (img *Image) DrawRectangle(rect image.Rectangle) {
	img.DrawRectangleAt(rect, 0, 0)
}

// DrawRectangleAt draws a portion of the image defined by an image.Rectangle,
// at the specified position.
func (img *Image) DrawRectangleAt(rect image.Rectangle, x, y int) {
	img.DrawRectAt(rect.Min.X, rect.Min.Y, rect.Dx(), rect.Dy(), x, y)
}

// DrawRectAt draws a portion of the image specified w
// at the specified position.
func (img *Image) DrawRectAt(rx, ry, rw, rh, x, y int) {
	C.textureDraw(img.id,
		C.int(img.W), C.int(img.H),
		C.int(rx), C.int(ry),
		C.int(rw), C.int(rh),
		C.float(x), C.float(y))
}

// Sprite creates a sprite from the image, with the specified grid.
func (img *Image) Sprite(w, h int) *Sprite {
	return NewSprite(img, w, h)
}
