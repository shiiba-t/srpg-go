package main

import (
	_ "embed"
	_ "image/png"
	"log"

	// "github.com/golang/freetype/truetype"
	// "golang.org/x/image/font"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/fonts/x12y20pxScanLine.ttf
// var sampleBytes []byte

// var (
// 	gamerFontS font.Face
// )

// マップ上の座標を表す構造体
// 左上を(0, 0)とする
type Position struct {
	x int // x座標
	y int // y座標
}

//ゲーム全体に必要なデータを格納
type Game struct {
	CursorPosition Position // カーソルの位置
	FrameCount     int      // フレーム数(カーソル点滅用)
	MoveCounter    int      // カーソルの動きを制限する用
}

// func init() {
// 	tt, err := truetype.Parse(sampleBytes)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	const dpi = 72
// 	gamerFontS = truetype.NewFace(tt, &truetype.Options{
// 		Size:    20,
// 		DPI:     dpi,
// 		Hinting: font.HintingFull,
// 	})
// }

// ゲーム構造体のコンストラクタ
func NewGame() *Game {
	g := &Game{}
	g.CursorPosition = Position{x: 0, y: 0}
	g.FrameCount = 0
	g.MoveCounter = 0
	return g
}

//Update is called each tic.
func (g *Game) Update() error {
	if g.MoveCounter > 5 {
		MoveCursor(g)
	}

	g.FrameCount++
	g.MoveCounter++
	return nil
}

//Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	floor, _, err := ebitenutil.NewImageFromFile("assets/map/grasslands.png")
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 30; i++ {
		for j := 0; j < 15; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(32*i), float64(32*j))
			screen.DrawImage(floor, op)
		}
	}

	// カーソル位置表示(0.5秒点滅)
	if (g.FrameCount % 60) > 30 {
		cursor, _, err := ebitenutil.NewImageFromFile("assets/map/selected.png")
		if err != nil {
			log.Fatal(err)
		}
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(32*g.CursorPosition.x), float64(32*g.CursorPosition.y))
		screen.DrawImage(cursor, op)
	}

	// text.Draw(screen, fmt.Sprintf("ABC"), gamerFontS, 0, 0, color.White)
	// text.Draw(screen, fmt.Sprintf("DEF"), gamerFontS, 0, 32, color.White)
	// text.Draw(screen, fmt.Sprintf("GHI"), gamerFontS, 32, 0, color.White)
	// text.Draw(screen, fmt.Sprintf("JKL"), gamerFontS, 32, 32, color.White)

	// img, _, err := ebitenutil.NewImageFromFile("assets/leaf.png")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(float64(64), float64(0))
	// screen.DrawImage(img, op)
}

//Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	return 960, 480
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(960, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("SRPG")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
