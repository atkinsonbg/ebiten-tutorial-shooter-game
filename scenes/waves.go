package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func WaveScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64) {
	imgWaveX := -30

	if currentTick <= 30 {
		imgWaveX = imgWaveX + int(currentTick)
	} else {
		imgWaveX = (imgWaveX + int(currentTick)) * -1
	}

	imgWave := assets.Wave()
	imgWaveWidth := imgWave.Bounds().Dx()
	imgWaveHeight := 280.00
	for imgWaveX < windowWidth {
		imgWoodOpts := &ebiten.DrawImageOptions{}
		imgWoodOpts.GeoM.Translate(float64(imgWaveX), imgWaveHeight)
		screen.DrawImage(imgWave, imgWoodOpts)
		imgWaveX = imgWaveWidth + imgWaveX
	}
}