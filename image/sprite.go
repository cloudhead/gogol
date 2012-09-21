package image

type Sprite struct {
	Clips        []*Clip
	ClipW, ClipH int
}

func NewSprite(img *Image, w, h int) *Sprite {
	s := &Sprite{
		ClipW: w,
		ClipH: h,
		Clips: []*Clip{},
	}

	for y := 0; y <= img.H-h; y += h {
		for x := 0; x <= img.W-w; x += w {
			clip := NewClip(img, x, y, x+w, y+h)
			s.Clips = append(s.Clips, clip)
		}
	}
	return s
}

func (s *Sprite) DrawClip(index int) {
	s.DrawClipAt(index, 0, 0)
}

func (s *Sprite) DrawClipAt(index, x, y int) {
	s.Clips[index].DrawAt(x, y)
}

func (s *Sprite) Sequence(from, to int) *Sequence {
	return NewSequence(s, from, to)
}
