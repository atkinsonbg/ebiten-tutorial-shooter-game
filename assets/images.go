package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

func init() {

}

func Curtain() *ebiten.Image {
	return getSubimage(StallSpritesheet, GetSubImageCoordinates(StallSpriteJsonBytes, "curtain"))
}

func CurtainStraight() *ebiten.Image {
	return getSubimage(StallSpritesheet, GetSubImageCoordinates(StallSpriteJsonBytes, "curtain_straight"))
}

func BackgroundGreen() *ebiten.Image {
	return getSubimage(StallSpritesheet, GetSubImageCoordinates(StallSpriteJsonBytes, "bg_green"))
}

func Wood() *ebiten.Image {
	return getSubimage(StallSpritesheet, GetSubImageCoordinates(StallSpriteJsonBytes, "bg_wood"))
}

func Wave() *ebiten.Image {
	return getSubimage(StallSpritesheet, GetSubImageCoordinates(StallSpriteJsonBytes, "wave"))
}

func Duck() *ebiten.Image {
	return getSubimage(ObjectsSpritesheet, GetSubImageCoordinates(ObjectsSpriteJsonBytes, "duck"))
}

func Stick() *ebiten.Image {
	return getSubimage(ObjectsSpritesheet, GetSubImageCoordinates(ObjectsSpriteJsonBytes, "stick"))
}

func getSubimage(spritesheet *ebiten.Image, coordinates map[string]interface{}) *ebiten.Image {
	xCoor := int(coordinates["x"].(float64))
	yCoor := int(coordinates["y"].(float64))
	wCoor := int(coordinates["x"].(float64)) + int(coordinates["w"].(float64))
	hCoor := int(coordinates["y"].(float64)) + int(coordinates["h"].(float64))
	rect := image.Rect(xCoor, yCoor, wCoor, hCoor)
	return spritesheet.SubImage(rect).(*ebiten.Image)
}