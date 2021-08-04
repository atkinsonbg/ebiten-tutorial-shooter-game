package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func CurtainsScene(screen *ebiten.Image, windowWidth int, windowHeight int, json []byte, sprite []byte) {
	// Curtain sides (left)
	screen.DrawImage(assets.Curtain(json, sprite), nil)

	// Curtain sides (right)
	curtainSideOpts := &ebiten.DrawImageOptions{}
	curtainSideOpts.GeoM.Scale(-1,1) // Flip it
	curtainSideOpts.GeoM.Translate(float64(windowWidth),0) // Weird, I would have thought I needed to subtract the width of the image here!
	screen.DrawImage(assets.Curtain(json, sprite), curtainSideOpts)

	// Curtain top
	curtainX := 0
	imgCurtainTop := assets.CurtainStraight(json, sprite)
	curtainTopWidth := imgCurtainTop.Bounds().Dx()
	for curtainX < windowWidth {
		curtainOpts := &ebiten.DrawImageOptions{}
		curtainOpts.GeoM.Translate(float64(curtainX), 0)
		screen.DrawImage(imgCurtainTop, curtainOpts)
		curtainX = curtainX + curtainTopWidth
	}
}

