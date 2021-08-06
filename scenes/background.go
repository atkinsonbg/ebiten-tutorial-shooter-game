package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func BackgroundScene(screen *ebiten.Image, windowWidth int, windowHeight int) {
	imgBackground := assets.BackgroundGreen()
	imgBackgroundX := 0
	imgBackgroundY := imgBackground.Bounds().Dy()
	imgBackgroundWidth := imgBackground.Bounds().Dx()

	for imgBackgroundX < windowWidth {
		backgroundOpts := &ebiten.DrawImageOptions{}
		backgroundOpts2 := &ebiten.DrawImageOptions{}
		// top row
		backgroundOpts.GeoM.Translate(float64(imgBackgroundX), 0)
		screen.DrawImage(imgBackground, backgroundOpts)
		// bottom row
		backgroundOpts2.GeoM.Translate(float64(imgBackgroundX), float64(imgBackgroundY))
		screen.DrawImage(imgBackground, backgroundOpts2)

		imgBackgroundX += imgBackgroundWidth
	}
}
