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

//
// #include "crayola.h"
// #include <GL/glut.h>
// #cgo LDFLAGS: -lglut -lGL
//
import "C"

import (
	"crayola/image"
	"time"
)

var (
	images   = []*image.Image{}
	lastTick = 0
)

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

	// Keyboard is called whenever a key is pressed or released.
	// The key is passed as an argument, as well as the state
	// of t
	Keyboard(k Key, isDown bool)

	// Mouse is called whenever the mouse is clicked.
	Mouse(button Mouse, isDown bool, x, y int)

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
	C.init(C.CString(handler.Title()))
}

// Translate moves the drawing position to the specified coordinates.
// Subsequent drawing operations will be relative to these coordinates.
func Translate(x, y float32) {
	C.glTranslatef(C.GLfloat(x), C.GLfloat(y), 0)
}

// Scale scales the view by x and y.
func Scale(x, y float32) {
	if x > 0 && y > 0 {
		C.glScalef(C.GLfloat(x), C.GLfloat(y), 1.0)
	}
}

// AdjustHSL adjusts the scene's hue, saturation & lightness.
// Values can range between 0 and 1.
func AdjustHSL(h, s, l float32) {
	C.adjustHSL(C.float(h), C.float(s), C.float(l))
}

// AdjustExp adjusts the scene's exposure. The first parameter
// specifies the exposure (default is 1.0), the second specifies
// the maximum brightness.
func AdjustExp(exp, max float32) {
	C.adjustExp(C.float(exp), C.float(max))
}

// ShowCursor shows the mouse cursor.
func ShowCursor() { C.glutSetCursor(C.GLUT_CURSOR_INHERIT) }

// HideCursor hides the mouse cursor.
func HideCursor() { C.glutSetCursor(C.GLUT_CURSOR_NONE) }

// FullScreen enables full-screen mode.
func FullScreen() { C.glutFullScreen() }

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

//export goReady
//
// goReady is called right before 'glutMainLoop'.
//
func goReady() {
	for _, img := range images {
		img.Gen()
	}
	handler.Ready()
}

//export goReshape
//
// goReshape is called whenever the window is resized.
//
func goReshape(w, h C.int) {
	handler.Reshape(int(w), int(h))
}

//export goDisplay
//
// goDisplay is called when the view is about to be redrawn.
//
func goDisplay(now C.int) {
	delta := int(now) - lastTick
	handler.Display(time.Millisecond * time.Duration(delta))
	lastTick = int(now)
}

//export goKeyboard
//
// goKeyboard is called when an ordinary key is pressed.
//
func goKeyboard(key C.uchar, _, _ C.int) {
	handler.Keyboard(Key(key), true)
}

//export goKeyboardUp
//
// goKeyboardUp is called when an ordinary key is released.
//
func goKeyboardUp(key C.uchar, _, _ C.int) {
	handler.Keyboard(Key(key), false)
}

//export goSpecial
//
// goSpecial is called when a special key is pressed.
//
func goSpecial(k, _, _ C.int) {
	handler.Keyboard(Key(k|keySpecial), true)
}

//export goSpecialUp
//
// goSpecialUp is called when a special key is released.
//
func goSpecialUp(k, _, _ C.int) {
	handler.Keyboard(Key(k|keySpecial), false)
}

//export goMouse
//
// goMouse is called when a mouse button is pressed or released.
//
func goMouse(button, state, x, y C.int) {
	handler.Mouse(Mouse(button), state == 0, int(x), int(y))
}

//export goMotion
//
// goMotion is called when the mouse is moved.
//
func goMotion(x, y C.int) {
	handler.Motion(int(x), int(y))
}

//export goEntry
//
// goEntry is called when the mouse enters/exits the view.
//
func goEntry(e C.int) {
	handler.Entry(e == 1)
}
