package assets

import (
	"bytes"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/png"
	"log"
)

func Curtain(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"curtain", sprite)
}

func CurtainStraight(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"curtain_straight", sprite)
}

func BackgroundGreen(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"bg_green", sprite)
}

func Wood(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"bg_wood", sprite)
}

func Wave(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"wave", sprite)
}

func Duck(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"duck", sprite)
}

func Stick(json []byte, sprite []byte) *ebiten.Image {
	return getSubimageByName(json,"stick", sprite)
}

func getSubimageByName(jsonBytes []byte, jsonName string, spriteBytes []byte) *ebiten.Image {
	jsonMap := map[string]interface{}{}
	_ = json.Unmarshal(jsonBytes, &jsonMap)
	coordinates := jsonMap[jsonName].(map[string]interface{})

	sheet := getEbitenImage(spriteBytes)
	xCoor := int(coordinates["x"].(float64))
	yCoor := int(coordinates["y"].(float64))
	wCoor := int(coordinates["x"].(float64)) + int(coordinates["w"].(float64))
	hCoor := int(coordinates["y"].(float64)) + int(coordinates["h"].(float64))
	rect := image.Rect(xCoor, yCoor, wCoor, hCoor)
	return sheet.SubImage(rect).(*ebiten.Image)
}

func getEbitenImage(spriteBytes []byte) *ebiten.Image {
	pngImage, err := png.Decode(bytes.NewReader(spriteBytes))
	if err != nil {
		log.Panic()
	}

	return ebiten.NewImageFromImage(pngImage)
}