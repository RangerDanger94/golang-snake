package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const screenWidth int = 640
const screenHeight int = 480
const gridSize int32 = 20

const fps uint32 = 60
const delayTime uint32 = 1000.0 / fps

const skipTicks uint32 = 100

var nextGameTick uint32 = 100

func borderCollision(rect *sdl.Rect) bool {
	if (rect.X < 0 || rect.X+rect.W > int32(screenWidth)) ||
		(rect.Y < 0 || rect.Y+rect.H > int32(screenHeight)) {
		return true
	}

	return false
}

func place(r *sdl.Rect) {
	cols := int32(math.Floor(float64(int32(screenWidth) / gridSize)))
	rows := int32(math.Floor(float64(int32(screenHeight) / gridSize)))

	*r = sdl.Rect{X: rand.Int31n(cols) * gridSize, Y: rand.Int31n(rows) * gridSize, W: gridSize, H: gridSize}
	fmt.Printf("Starting Position: %v,%v\n", r.X, r.Y)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	sdl.Init(sdl.INIT_EVERYTHING)

	// Create Window
	window, err := sdl.CreateWindow(
		"Snake xTreme 2016",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	// Create Renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Set up our snake & food
	s := new(0, 0, gridSize, gridSize)
	f := sdl.Rect{}
	place(&f)

	// Main Loop
	running := true
	for running {
		frameStart := sdl.GetTicks()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyDownEvent:
				switch t.Keysym.Sym {
				case sdl.K_LEFT:
					s.move(-gridSize, 0)
				case sdl.K_RIGHT:
					s.move(gridSize, 0)
				case sdl.K_UP:
					s.move(0, -gridSize)
				case sdl.K_DOWN:
					s.move(0, gridSize)
				}
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.SetDrawColor(0, 128, 255, 255)
		renderer.Clear()

		if sdl.GetTicks() > nextGameTick {
			fmt.Printf(s.string())
			// Restart
			if borderCollision(&s.body[0]) || s.dead() {
				s = new(0, 0, gridSize, gridSize)
			}

			s.update()

			if s.head().HasIntersection(&f) {
				s.eat()
				place(&f)
			}

			nextGameTick += skipTicks
		}

		renderer.SetDrawColor(100, 255, 0, 255)
		renderer.FillRects(s.body)

		renderer.SetDrawColor(255, 100, 100, 255)
		renderer.FillRect(&f)

		renderer.Present()

		if frameTime := sdl.GetTicks() - frameStart; frameTime < delayTime {
			sdl.Delay(delayTime - frameTime)
		}
	}

	// Clean Up
	sdl.Quit()
}
