package image

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"image"
	"log"
)

type Image struct {
	*glfw.Image

	Texture  gl.Texture
	FilePath string
	W, H     int
}

var cache = map[string]*Image{}

// New initializes a new image from the given file path.
func New(path string) *Image {
	if cache[path] != nil {
		return cache[path]
	}

	i, err := glfw.ReadImage(path, glfw.OriginUlBit|glfw.NoRescaleBit)
	if err != nil {
		log.Fatalf("image %s: failed to read: %s", path, err)
	}

	img := &Image{
		Image:    i,
		FilePath: path,
		W:        i.Width(),
		H:        i.Height(),
	}
	cache[path] = img

	return img
}

func (img *Image) Gen() {
	texture := gl.GenTexture()
	texture.Bind(gl.TEXTURE_2D)

	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameterf(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)

	if !img.LoadTextureImage2D(0x0) {
		log.Fatalf("image %s: failed to load texture", img.FilePath)
	}
	img.Texture = texture
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
	img.drawRectAt(float64(img.W), float64(img.H), float64(rx), float64(ry), float64(rw), float64(rh), float64(x), float64(y))
}

// Sprite creates a sprite from the image, with the specified grid.
func (img *Image) Sprite(w, h int) *Sprite {
	return NewSprite(img, w, h)
}

func (img *Image) drawRectAt(tw, th, x, y, w, h, sx, sy float64) {
	gl.PushMatrix()

	gl.MatrixMode(gl.MODELVIEW)
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	img.Texture.Bind(gl.TEXTURE_2D)

	rx, ry := x/tw, y/th
	rw, rh := w/tw, h/th

	gl.Begin(gl.QUADS)

	gl.TexCoord2d(rx, ry)
	gl.Vertex2d(sx, sy)

	gl.TexCoord2d(rx+rw, ry)
	gl.Vertex2d(sx+w, sy)

	gl.TexCoord2d(rx+rw, (ry + rh))
	gl.Vertex2d(sx+w, sy+h)

	gl.TexCoord2d(rx, (ry + rh))
	gl.Vertex2d(sx, sy+h)

	gl.End()

	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.BLEND)

	gl.PopMatrix()
}
