package scenes

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Duck struct {
	X float64
	Y float64
}

var ducks []Duck

func DuckScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64) {
	log.Print(len(ducks))
	if len(ducks) == 0 {
		d := Duck{-100, 200}
		ducks = append(ducks, d)
	}

	if len(ducks) < 4 {
		lastDuck := ducks[len(ducks)-1]
		if lastDuck.X > 200 {
			d := Duck{-100, 200}
			ducks = append(ducks, d)
		}
	}

	duckToDelete := 0
	for i, d := range ducks {
		ducks[i].X += 12
		imgDuck := assets.Duck()
		imgDuckOpts := &ebiten.DrawImageOptions{}
		imgDuckOpts.GeoM.Translate(d.X, d.Y)
		screen.DrawImage(imgDuck, imgDuckOpts)
		
		if d.X > float64(windowWidth) {
			duckToDelete = i
		}
	}

	if duckToDelete > 0 {
		ducks = append(ducks[:duckToDelete], ducks[duckToDelete+1:]...)
	}
}