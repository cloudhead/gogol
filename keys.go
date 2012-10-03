package gogol

//
// #include <GL/glut.h>
// #cgo LDFLAGS: -lglut -lGL
//
import "C"

type Key uint

const (
	keyChar    = 0x0
	keySpecial = 0xf00
)

const (
	KeyBackspace Key = keyChar | 8
	KeyTab           = keyChar | 9
	KeyReturn        = keyChar | 13
	KeyEsc           = keyChar | 27
	KeySpace         = keyChar | 32
	KeyLeft          = keySpecial | C.GLUT_KEY_LEFT
	KeyUp            = keySpecial | C.GLUT_KEY_UP
	KeyRight         = keySpecial | C.GLUT_KEY_RIGHT
	KeyDown          = keySpecial | C.GLUT_KEY_DOWN
	KeyShiftL        = keySpecial | 112
	KeyShiftR        = keySpecial | 113
	KeyCtrlL         = keySpecial | 114
	KeyCtrlR         = keySpecial | 115
	KeyAltL          = keySpecial | 116
	KeyAltR          = keySpecial | 117
	KeyF1            = keySpecial | C.GLUT_KEY_F1
	KeyF2            = keySpecial | C.GLUT_KEY_F2
	KeyF3            = keySpecial | C.GLUT_KEY_F3
	KeyF4            = keySpecial | C.GLUT_KEY_F4
	KeyF5            = keySpecial | C.GLUT_KEY_F5
	KeyF6            = keySpecial | C.GLUT_KEY_F6
	KeyF7            = keySpecial | C.GLUT_KEY_F7
	KeyF8            = keySpecial | C.GLUT_KEY_F8
	KeyF9            = keySpecial | C.GLUT_KEY_F9
	KeyF10           = keySpecial | C.GLUT_KEY_F10
	KeyF11           = keySpecial | C.GLUT_KEY_F11
	KeyF12           = keySpecial | C.GLUT_KEY_F12
)
