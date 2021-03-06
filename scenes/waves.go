package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

func WaveScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64) {
	imgWaveX := -30
	imgWaveY := 320.00

	if currentTick <= 30 {
		imgWaveX = imgWaveX + int(currentTick)
		imgWaveY = imgWaveY + currentTick
	} else {
		imgWaveX = (imgWaveX + int(currentTick)) * -1
		imgWaveY = (imgWaveY + 30) - (currentTick - 30)
	}

	imgWave := assets.Wave()
	imgWaveWidth := imgWave.Bounds().Dx()
	for imgWaveX < windowWidth {
		imgWoodOpts := &ebiten.DrawImageOptions{}
		imgWoodOpts.GeoM.Translate(float64(imgWaveX), imgWaveY)
		screen.DrawImage(imgWave, imgWoodOpts)
		imgWaveX = imgWaveWidth + imgWaveX
	}
}