package main

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/atkinsonbg/ebiten-tutorial/scenes"
	"github.com/atkinsonbg/ebiten-tutorial/utils"
	_ "image/png"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var windowWidth = 900
var windowHeight = 520

type Game struct {
	tick float64
	lastClickedAt time.Time
	currentScore *utils.Score
}

func init() {

}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.currentScore == nil {
		g.currentScore = &utils.Score{}
	}
	var y = 0
	var x = 0
	g.lastClickedAt, x, y = utils.CheckMouseClick(g.lastClickedAt)

	if !assets.RagtimeMusicPlayer.IsPlaying() {
		assets.RagtimeMusicPlayer.Play()
	}
	g.updateTick()

	scenes.BackgroundScene(screen, windowWidth, windowHeight)
	scenes.DuckScene(screen, windowWidth, windowHeight, g.tick, x, y, g.currentScore)
	scenes.WaveScene(screen, windowWidth, windowHeight, g.tick)
	scenes.WoodScene(screen, windowWidth, windowHeight)
	scenes.CurtainsScene(screen, windowWidth, windowHeight)
	scenes.HudTextScene(screen, windowWidth, windowHeight, *g.currentScore)

	utils.PrintMemUsage(screen)
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
