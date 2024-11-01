package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	cellSize     = 10
	gridWidth    = screenWidth / cellSize
	gridHeight   = screenHeight / cellSize
	fps          = 50
)

var (
	bgColor = color.RGBA{0, 0, 0, 0}
	fgColor = color.RGBA{255, 255, 255, 255}
)

type Game struct {
	grid    [][]bool
	nextGen [][]bool
}

func NewGame() *Game {
	g := &Game{
		grid:    make([][]bool, gridWidth),
		nextGen: make([][]bool, gridWidth),
	}

	// Initialize grids
	for i := range g.grid {
		g.grid[i] = make([]bool, gridHeight)
		g.nextGen[i] = make([]bool, gridHeight)
	}

	// Randomize initial state
	for i := range g.grid {
		for j := range g.grid[i] {
			g.grid[i][j] = rand.Float32() < 0.5
		}
	}

	return g
}

func (g *Game) Update() error {
	// Calculate next generation
	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			neighbors := g.countNeighbors(x, y)
			current := g.grid[x][y]

			// Apply Conway's rules
			g.nextGen[x][y] = false
			if current && (neighbors == 2 || neighbors == 3) {
				g.nextGen[x][y] = true
			} else if !current && neighbors == 3 {
				g.nextGen[x][y] = true
			}
		}
	}

	// Swap grids
	g.grid, g.nextGen = g.nextGen, g.grid
	return nil
}

func (g *Game) countNeighbors(x, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// wrap around the edges
			nx := (x + i + gridWidth) % gridWidth
			ny := (y + j + gridHeight) % gridHeight
			if g.grid[nx][ny] {
				count++
			}
		}
	}
	return count
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	for x := 0; x < gridWidth; x++ {
		for y := 0; y < gridHeight; y++ {
			if g.grid[x][y] {
				vector.DrawFilledRect(
					screen,
					float32(x*cellSize),
					float32(y*cellSize),
					float32(cellSize-1),
					float32(cellSize-1),
					fgColor,
					false,
				)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")
	ebiten.SetWindowDecorated(false)
	ebiten.SetTPS(fps) // Set the ticks per second

	if err := ebiten.RunGameWithOptions(NewGame(), &ebiten.RunGameOptions{
		ScreenTransparent: true,
	}); err != nil {
		log.Fatal(err)
	}
}
