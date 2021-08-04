package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func WoodScene(screen *ebiten.Image, windowWidth int, windowHeight int, json []byte, sprite []byte) {
	imgWoodX := 0
	imgWood := assets.Wood(json, sprite)
	imgWoodWidth := imgWood.Bounds().Dx()
	imgWoodHeight := imgWood.Bounds().Dy() / 2
	for imgWoodX < windowWidth {
		imgWoodOpts := &ebiten.DrawImageOptions{}
		imgWoodOpts.GeoM.Translate(float64(imgWoodX), float64(windowHeight - imgWoodHeight))
		screen.DrawImage(imgWood, imgWoodOpts)
		imgWoodX = imgWoodWidth + imgWoodX
	}
}
