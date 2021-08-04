package main

import (
	"embed"
	"github.com/atkinsonbg/ebiten-tutorial/scenes"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var windowWidth = 900
var windowHeight = 520

//go:embed assets/images/*
var content embed.FS

type Game struct {
	tick float64
	stallSprite []byte
	stallJson []byte
	duckSprite []byte
	duckJson []byte
}

func init() {
}

func (g *Game) loadAllImageData() {
	if len(g.stallJson) == 0 {
		g.stallJson, _ = content.ReadFile("assets/images/stall.json")
		log.Print("Loaded stall JSON data...")
	}

	if len(g.stallSprite) == 0 {
		g.stallSprite, _ = content.ReadFile("assets/images/stall_sprite.png")
		log.Print("Loaded stall sprite data...")
	}

	if len(g.duckJson) == 0 {
		g.duckJson, _ = content.ReadFile("assets/images/objects.json")
		log.Print("Loaded duck JSON data...")
	}

	if len(g.duckSprite) == 0 {
		g.duckSprite, _ = content.ReadFile("assets/images/objects_sprite.png")
		log.Print("Loaded duck sprite data...")
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.loadAllImageData()
	g.updateTick()

	scenes.BackgroundScene(screen, windowWidth, windowHeight, g.stallJson, g.stallSprite)
	scenes.DuckScene(screen, windowWidth, windowHeight, g.tick, g.duckJson, g.duckSprite)
	scenes.WaveScene(screen, windowWidth, windowHeight, g.tick, g.stallJson, g.stallSprite)
	scenes.WoodScene(screen, windowWidth, windowHeight, g.stallJson, g.stallSprite)
	scenes.CurtainsScene(screen, windowWidth, windowHeight, g.stallJson, g.stallSprite)
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
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
