package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {

}

func Curtain() *ebiten.Image {
	img, _ := Stall.GetSprite("curtain", true)
	return img
}

func CurtainStraight() *ebiten.Image {
	img, _ := Stall.GetSprite("curtain_straight", true)
	return img
}

func BackgroundGreen() *ebiten.Image {
	img, _ := Stall.GetSprite("bg_green", true)
	return img
}

func Wood() *ebiten.Image {
	img, _ := Stall.GetSprite("bg_wood", true)
	return img
}

func Wave() *ebiten.Image {
	img, _ := Stall.GetSprite("water1", true)
	return img
}

func Duck() *ebiten.Image {
	img, _ := Objects.GetSprite("duck_outline_target_brown", false)
	return img
}

func Stick() *ebiten.Image {
	img, _ := Objects.GetSprite("stick_woodFixed_outline", true)
	return img
}

func BulletHole() *ebiten.Image {
	img, _ := Objects.GetSprite("shot_grey_small", true)
	return img
}

func Crosshairs() *ebiten.Image {
	img, _ := Hud.GetSprite("crosshair_red_small", true)
	return img
}