package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Img struct {
	X float64
	Y float64
}

type Duck struct {
	Duck Img
	Stick Img
}

var ducks []Duck

func init() {
	addDuck()
}

func DuckScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64) {
	if len(ducks) < 4 {
		lastDuck := ducks[len(ducks)-1]
		if lastDuck.Duck.X > 200 {
			addDuck()
		}
	}

	for i, d := range ducks {
		ducks[i].Duck.X += 2
		ducks[i].Stick.X += 2

		d.Duck.Y = d.Duck.Y - math.Cos(d.Duck.X / 100) * -100
		d.Stick.Y = d.Duck.Y + 118

		imgDuck := assets.Duck()
		imgDuckOpts := &ebiten.DrawImageOptions{}
		imgDuckOpts.GeoM.Translate(d.Duck.X, d.Duck.Y)
		screen.DrawImage(imgDuck, imgDuckOpts)

		imgStick := assets.Stick()
		imgStickOpts := &ebiten.DrawImageOptions{}
		imgStickOpts.GeoM.Translate(d.Stick.X, d.Stick.Y)
		screen.DrawImage(imgStick, imgStickOpts)
		
		if ducks[i].Duck.X > float64(windowWidth) {
			ducks[i].Duck.X = -336
			ducks[i].Duck.Y = 200
			ducks[i].Stick.X = -298
			ducks[i].Stick.Y = 318
		}
	}
}

func addDuck() {
	// Initial duck that starts off screen to left
	startingDuck := Img{-136, 200}
	startingStick := Img{-98,318}

	d := Duck{}
	d.Duck = startingDuck
	d.Stick = startingStick
	ducks = append(ducks, d)
}