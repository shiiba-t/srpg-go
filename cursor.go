package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func MoveCursor(g *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.CursorPosition.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.CursorPosition.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.CursorPosition.x -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.CursorPosition.x += 1
	}

	g.MoveCounter = 0
}
