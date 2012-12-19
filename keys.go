package gogol

import "github.com/go-gl/glfw"

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
	KeyLeft          = keySpecial | glfw.KeyLeft
	KeyUp            = keySpecial | glfw.KeyUp
	KeyRight         = keySpecial | glfw.KeyRight
	KeyDown          = keySpecial | glfw.KeyDown
	KeyShiftL        = keySpecial | 112
	KeyShiftR        = keySpecial | 113
	KeyCtrlL         = keySpecial | 114
	KeyCtrlR         = keySpecial | 115
	KeyAltL          = keySpecial | 116
	KeyAltR          = keySpecial | 117
	KeyF1            = keySpecial | glfw.KeyF1
	KeyF2            = keySpecial | glfw.KeyF2
	KeyF3            = keySpecial | glfw.KeyF3
	KeyF4            = keySpecial | glfw.KeyF4
	KeyF5            = keySpecial | glfw.KeyF5
	KeyF6            = keySpecial | glfw.KeyF6
	KeyF7            = keySpecial | glfw.KeyF7
	KeyF8            = keySpecial | glfw.KeyF8
	KeyF9            = keySpecial | glfw.KeyF9
	KeyF10           = keySpecial | glfw.KeyF10
	KeyF11           = keySpecial | glfw.KeyF11
	KeyF12           = keySpecial | glfw.KeyF12
)
