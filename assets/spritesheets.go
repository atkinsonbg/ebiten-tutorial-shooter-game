package assets

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
	"log"
)

//go:embed images/objects_sprite.png
var objectsSpritesheetBytes []byte
//go:embed images/objects.json
var ObjectsSpriteJsonBytes []byte
var ObjectsSpritesheet *ebiten.Image

//go:embed images/stall_sprite.png
var stallSpritesheetBytes []byte
//go:embed images/stall.json
var StallSpriteJsonBytes []byte
var StallSpritesheet *ebiten.Image

func init() {
	ObjectsSpritesheet = toImage(objectsSpritesheetBytes)
	StallSpritesheet = toImage(stallSpritesheetBytes)
}

func GetSubImageCoordinates(jsonBytes []byte, imageName string) map[string]interface{} {
	jsonMap := map[string]interface{}{}
	_ = json.Unmarshal(jsonBytes, &jsonMap)
	coordinates := jsonMap[imageName].(map[string]interface{})
	return coordinates
}

func toImage(imageBytes []byte) *ebiten.Image {
	pngImage, err := png.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		log.Panicln(err)
	}

	ebitenImage := ebiten.NewImageFromImage(pngImage)

	return ebitenImage
}