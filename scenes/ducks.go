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

func DuckScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64, json []byte, sprite []byte) {
	if len(ducks) == 0 {
		addDuck()
	}

	if len(ducks) < 4 {
		lastDuck := ducks[len(ducks)-1]
		if lastDuck.Duck.X > 200 {
			addDuck()
		}
	}

	duckToDelete := 0
	for i, d := range ducks {
		ducks[i].Duck.X += 10
		ducks[i].Stick.X += 10
		
		d.Duck.Y = d.Duck.Y - math.Cos(d.Duck.X / 100) * -100
		d.Stick.Y = d.Duck.Y + 118

		imgDuck := assets.Duck(json, sprite)
		imgDuckOpts := &ebiten.DrawImageOptions{}
		imgDuckOpts.GeoM.Translate(d.Duck.X, d.Duck.Y)
		screen.DrawImage(imgDuck, imgDuckOpts)

		imgStick := assets.Stick(json, sprite)
		imgStickOpts := &ebiten.DrawImageOptions{}
		imgStickOpts.GeoM.Translate(d.Stick.X, d.Stick.Y)
		screen.DrawImage(imgStick, imgStickOpts)
		
		if d.Duck.X > float64(windowWidth) {
			duckToDelete = i
		}
	}

	if duckToDelete > 0 {
		ducks = append(ducks[:duckToDelete], ducks[duckToDelete+1:]...)
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