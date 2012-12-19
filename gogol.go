//
// Package gogol provides an abstract interface over
// bitmap graphic rendering and animation.
//
// Create an empty event handler and initialize gogol
// with it:
//
//    handler := gogol.DefaultHandler{}
//    gogol.Init(handler)
//
// Create an animation from a sprite, and draw the current
// frame:
//
//    sprite := gogol.NewImage("goblin.bmp").Sprite(16, 16)
//    animation := sprite.Sequence(0, -1)
//    ...
//    animation.DrawFrameAt(0, 0)
//
package gogol

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"gogol/image"
	"log"
	"time"
)

var (
	images    = []*image.Image{}
	lastTick  = 0
	winWidth  = 0
	winHeight = 0
)

// Mouse can be either MouseL, MouseR or MouseR.
type Mouse int

const (
	MouseL Mouse = 0
	MouseM       = 1
	MouseR       = 2
)

var handler Handler

type Handler interface {
	// Ready is called once, after the window is created,
	// but before anything is drawn to the screen.
	Ready()

	// Reshape is called every time the window is resized.
	// The new width and height are passed as arguments.
	Reshape(w, h int)

	// Display is called on every frame, right before the
	// view is re-drawn. The time delta since the last frame
	// is passed as an argument.
	//
	// This is the recommend place to perform all drawing
	// operations.
	Display(delta time.Duration)

	// DisplayPost is called on every frame, after Display,
	// and after any post-processing effects have been applied.
	// This function is useful for drawing UI elements which
	// shouldn't be affected by image filters.
	DisplayPost()

	// Keyboard is called whenever a key is pressed or released.
	// The key is passed as an argument, as well as the state
	// of t
	Keyboard(k Key, isDown bool)

	// Mouse is called whenever the mouse is clicked.
	Mouse(button Mouse, isDown bool)

	// Motion is called whenever the mouse is moved, with the x
	// and y position of the cursor.
	Motion(x, y int)

	// Entry is called when the mouse enters or exits the window.
	// true is passed if the mouse is entering, false is passed
	// if it is exiting.
	Entry(isEnter bool)

	// Title is called to set the title of the window.
	Title() string
}

// Init initializes the engine and creates a window
// with the specified title.
func Init(h Handler) {
	handler = h

	if err := glfw.Init(); err != nil {
		log.Fatal(err)
	}
	defer glfw.Terminate()

	if err := glfw.OpenWindow(0, 0, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
		log.Fatal(err)
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1) // Vsync
	glfw.SetWindowTitle(h.Title())
	glfw.SetKeyCallback(goKeyboard)
	glfw.SetWindowSizeCallback(goReshape)
	glfw.SetMouseButtonCallback(goMouse)
	glfw.SetMousePosCallback(goMotion)

	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)

	handler.Ready()

	for _, img := range images {
		img.Gen()
	}

	lastTick := glfw.Time()

	for {
		now := glfw.Time()
		delta := time.Duration(now - lastTick)
		lastTick = now

		render(delta)
	}
}

// Translate moves the drawing position to the specified coordinates.
// Subsequent drawing operations will be relative to these coordinates.
func Translate(x, y float32) {
	gl.Translatef(x, y, 0)
}

// Scale scales the view by x and y.
func Scale(x, y float32) {
	if x > 0 && y > 0 {
		gl.Scalef(x, y, 1.0)
	}
}

// WindowSize returns the current window width & height
func WindowSize() (int, int) {
	return winWidth, winHeight
}

// AdjustHSL adjusts the scene's hue, saturation & lightness.
// Values can range between 0 and 1.
func AdjustHSL(h, s, l float32) {
	panic("not implemented")
}

// AdjustExp adjusts the scene's exposure. The first parameter
// specifies the exposure (default is 1.0), the second specifies
// the maximum brightness.
func AdjustExp(exp, max float32) {
	panic("not implemented")
}

// ShowCursor shows the mouse cursor.
func ShowCursor() { glfw.Enable(glfw.MouseCursor) }

// HideCursor hides the mouse cursor.
func HideCursor() { glfw.Disable(glfw.MouseCursor) }

// FullScreen enables full-screen mode.
func FullScreen() { panic("not implemented") }

// NewImage creates a *image.Image from the specified path.
func NewImage(path string) *image.Image {
	img := image.New(path)
	images = append(images, img)

	return img
}

// NewClip creates a *image.Clip from the given *image.Image.
// x and y specify the top-left corner on the image, w and h
// specify the width and height.
func NewClip(img *image.Image, x, y, w, h int) *image.Clip {
	return image.NewClip(img, x, y, w, h)
}

// NewSprite creates a *image.Sprite from the given *image.Image.
// w and h specify the width and height of the clips.
func NewSprite(img *image.Image, w, h int) *image.Sprite {
	return image.NewSprite(img, w, h)
}

// goReshape is called whenever the window is resized.
func goReshape(w, h int) {
	winWidth, winHeight = w, h

	gl.ClearColor(1, 1, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, float64(w), float64(h), 0, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	// fbo.Bind(gl.TEXTURE_2D)
	// gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, winWidth, winHeight, 0, gl.RGBA, gl.UNSIGNED_BYTE, nil)
	// fbo.Unbind(gl.TEXTURE_2D)

	glfw.SwapBuffers()

	handler.Reshape(w, h)
}

// goDisplay is called when the view is about to be redrawn.
func goDisplay(delta time.Duration) {
	handler.Display(time.Millisecond * delta)
}

func goDisplayPost(now int) {
	handler.DisplayPost()
}

// goKeyboard is called when an ordinary key is pressed.
func goKeyboard(key, state int) {
	handler.Keyboard(Key(key), true)
}

// goKeyboardUp is called when an ordinary key is released.
func goKeyboardUp(key byte, _, _ int) {
	handler.Keyboard(Key(key), false)
}

// goSpecial is called when a special key is pressed.
func goSpecial(k, _, _ int) {
	handler.Keyboard(Key(k|keySpecial), true)
}

// goSpecialUp is called when a special key is released.
func goSpecialUp(k, _, _ int) {
	handler.Keyboard(Key(k|keySpecial), false)
}

// goMouse is called when a mouse button is pressed or released.
func goMouse(button, state int) {
	handler.Mouse(Mouse(button), state == 0)
}

// goMotion is called when the mouse is moved.
func goMotion(x, y int) {
	handler.Motion(int(x), int(y))
}

// goEntry is called when the mouse enters/exits the view.
func goEntry(e int) {
	handler.Entry(e == 1)
}

func render(delta time.Duration) {
	gl.ClearColor(1, 1, 1, 0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.Viewport(0, 0, winWidth, winHeight)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0.0, float64(winWidth), float64(winHeight), 0, -1, 1)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	goDisplay(delta)

	glfw.SwapBuffers()
}
