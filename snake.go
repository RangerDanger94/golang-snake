package main

import "github.com/veandco/go-sdl2/sdl"
import "fmt"

// Snake struct
type Snake struct {
	x, y, width, height int32
	xSpeed, ySpeed      int32
	body                []sdl.Rect
	length              int
}

// New snake
func new(x, y, width, height int32) (s Snake) {
	s.xSpeed, s.ySpeed = width, 0
	s.x, s.y = x, y
	s.width, s.height = width, height
	s.length = 1
	s.body = make([]sdl.Rect, 1)
	s.body[0] = sdl.Rect{X: s.x, Y: s.y, W: s.width, H: s.height}
	return
}

func (s *Snake) dead() bool {
	for i := 1; i < len(s.body); i++ {
		if selfHarm := s.head().HasIntersection(&s.body[i]); selfHarm == true {
			return true
		}
	}
	return false
}

func (s *Snake) eat() {
	s.length += 3
}

func (s *Snake) update() {
	s.x, s.y = s.x+s.xSpeed, s.y+s.ySpeed

	newHead := []sdl.Rect{
		{X: s.x, Y: s.y, W: s.width, H: s.height},
	}

	shiftedBody := s.body[:len(s.body)-1]
	if s.length != len(s.body) {
		shiftedBody = s.body[:len(s.body)]
	}

	s.body = append(newHead, shiftedBody...)
}

func (s *Snake) head() *sdl.Rect {
	return &s.body[0]
}

func (s *Snake) string() string {
	var r string

	for i, v := range s.body {
		r += fmt.Sprintf("index %d = X:%d, Y:%d W:%d, H:%d\n", i, v.X, v.Y, v.W, v.H)
	}

	return r
}

func (s *Snake) move(dirX, dirY int32) {
	if !((s.xSpeed != 0 && dirX != 0) || (s.ySpeed != 0 && dirY != 0)) {
		s.xSpeed = dirX
		s.ySpeed = dirY
	}
}
