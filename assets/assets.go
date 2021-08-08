package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func init() {

}

func Curtain() *ebiten.Image {
	img, _ := Stall.GetSprite("curtain")
	return img
}

func CurtainStraight() *ebiten.Image {
	img, _ := Stall.GetSprite("curtain_straight")
	return img
}

func BackgroundGreen() *ebiten.Image {
	img, _ := Stall.GetSprite("bg_green")
	return img
}

func Wood() *ebiten.Image {
	img, _ := Stall.GetSprite("bg_wood")
	return img
}

func Wave() *ebiten.Image {
	img, _ := Stall.GetSprite("water1")
	return img
}

func Duck() *ebiten.Image {
	img, _ := Objects.GetSprite("duck_outline_target_brown")
	return img
}

func Stick() *ebiten.Image {
	img, _ := Objects.GetSprite("stick_woodFixed_outline")
	return img
}