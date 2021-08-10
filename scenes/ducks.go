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
	DuckImage *ebiten.Image
	Duck Img
	Stick Img
	BulletHole Img
	HasBeenShot bool
}

var ducks []Duck

func init() {
}

func DuckScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentTick float64, currentX int, currentY int) {
	if len(ducks) < 4 {
		if (len(ducks) == 0) {
			addDuck()
		} else {
			lastDuck := ducks[len(ducks)-1]
			if lastDuck.Duck.X > 200 {
				addDuck()
			}
		}
	}

	for i, d := range ducks {
		ducks[i].Duck.X += 2
		ducks[i].Stick.X += 2

		d.Duck.Y = d.Duck.Y - math.Cos(d.Duck.X / 100) * -40
		d.Stick.Y = d.Duck.Y + 108

		if currentX > 0 && currentY > 0 {
			if (float64(currentX) > d.Duck.X && float64(currentX) < d.Duck.X+200) && (float64(currentY) > d.Duck.Y && float64(currentY) < d.Duck.Y+200) {
				if !d.HasBeenShot {
					assets.HitPlayer.Play()
					assets.HitPlayer.Rewind()
					ducks[i].BulletHole = Img{float64(currentX) - d.Duck.X, float64(currentY) - d.Duck.Y}
					ducks[i].HasBeenShot = true
				}
			}
		}
		
		imgDuckOpts := &ebiten.DrawImageOptions{}
		imgDuckOpts.GeoM.Translate(d.Duck.X, d.Duck.Y)

		if d.HasBeenShot {
			imgHole := assets.BulletHole()
			imgHoleOpts := &ebiten.DrawImageOptions{}
			imgHoleOpts.GeoM.Translate(d.BulletHole.X, d.BulletHole.Y)
			d.DuckImage.DrawImage(imgHole, imgHoleOpts)
		}

		screen.DrawImage(d.DuckImage, imgDuckOpts)

		imgStick := assets.Stick()
		imgStickOpts := &ebiten.DrawImageOptions{}
		imgStickOpts.GeoM.Translate(d.Stick.X, d.Stick.Y)
		screen.DrawImage(imgStick, imgStickOpts)
		
		if d.Duck.X > float64(windowWidth) {
			ducks[i].Duck.X = -426
			ducks[i].Duck.Y = 200
			ducks[i].Stick.X = -388
			ducks[i].Stick.Y = 288
			ducks[i].HasBeenShot = false
			ducks[i].BulletHole = Img{}
			ducks[i].DuckImage = assets.Duck()
		}

		cpx, cpy := ebiten.CursorPosition()
		imgCrosshairs := assets.Crosshairs()
		imgCrosshairsOpts := &ebiten.DrawImageOptions{}
		imgCrosshairsOpts.GeoM.Translate(float64(cpx - 20), float64(cpy - 20))
		screen.DrawImage(imgCrosshairs, imgCrosshairsOpts)
	}
}

func addDuck() {
	// Initial duck that starts off screen to left
	startingDuck := Img{-136, 200}
	startingStick := Img{-98,288}

	d := Duck{}
	d.DuckImage = assets.Duck()
	d.Duck = startingDuck
	d.Stick = startingStick
	ducks = append(ducks, d)
}