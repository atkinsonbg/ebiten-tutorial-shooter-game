package main

import (
	"fmt"
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/atkinsonbg/ebiten-tutorial/scenes"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
	"runtime"

	"github.com/hajimehoshi/ebiten/v2"
)

var windowWidth = 900
var windowHeight = 520

type Game struct {
	tick float64
}

func init() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !assets.RagtimeMusicPlayer.IsPlaying() {
		assets.RagtimeMusicPlayer.Play()
	}
	g.updateTick()

	scenes.BackgroundScene(screen, windowWidth, windowHeight)
	scenes.DuckScene(screen, windowWidth, windowHeight, g.tick)
	scenes.WaveScene(screen, windowWidth, windowHeight, g.tick)
	scenes.WoodScene(screen, windowWidth, windowHeight)
	scenes.CurtainsScene(screen, windowWidth, windowHeight)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v, Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB, NumGC = %v\n", ebiten.CurrentFPS(), bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func (g *Game) updateTick() {
	if g.tick >= 60 {
		g.tick = 0
	}
	g.tick++
}

func main() {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Ebiten Shooter Game Demo")
	g := &Game{}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
