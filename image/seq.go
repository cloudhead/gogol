package image

import (
	"math"
	"time"
)

type Sequence struct {
	Sprite              *Sprite
	Speed               float64
	StartTime, StopTime time.Time
	From, To, Curr      int
	IsPlaying           bool
}

func NewSequence(s *Sprite, from, to int) *Sequence {
	if from < 0 {
		from = len(s.Clips) + from
	}
	if to < 0 {
		to = len(s.Clips) + to
	}
	return &Sequence{
		Sprite: s,
		Speed:  1,
		From:   from,
		To:     to,
		Curr:   0,
	}
}

func (s *Sequence) DrawFrameAt(x, y int) {
	if s.IsPlaying {
		elapsed := float64(time.Now().Sub(s.StartTime).Nanoseconds()) / 1e9
		fraction := s.Speed * float64(elapsed)
		s.Curr = s.From + int(math.Floor(fraction))%(s.To-s.From+1)
	}
	s.Sprite.DrawClipAt(s.Curr, x, y)
}

func (s *Sequence) Play(speed float64) {
	s.StartTime = time.Now()
	s.IsPlaying = true
	s.Speed = speed
}

func (s *Sequence) Pause() {
	s.StopTime = time.Now()
	s.IsPlaying = false
}

func (s *Sequence) Toggle() {
	if s.IsPlaying {
		s.Pause()
	} else {
		s.Play(s.Speed)
	}
}
