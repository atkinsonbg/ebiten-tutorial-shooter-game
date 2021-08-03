package assets

import (
	"bytes"
	"embed"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/png"
	"log"
)

//go:embed images/*
var content embed.FS

func Curtain() *ebiten.Image {
	return getSubimageByName("images/stall.json","curtain", "stall_sprite.png")
}

func CurtainStraight() *ebiten.Image {
	return getSubimageByName("images/stall.json","curtain_straight", "stall_sprite.png")
}

func BackgroundGreen() *ebiten.Image {
	return getSubimageByName("images/stall.json","bg_green", "stall_sprite.png")
}

func Wood() *ebiten.Image {
	return getSubimageByName("images/stall.json","bg_wood", "stall_sprite.png")
}

func Wave() *ebiten.Image {
	return getSubimageByName("images/stall.json","wave", "stall_sprite.png")
}

func Duck() *ebiten.Image {
	return getSubimageByName("images/objects.json","duck", "objects_sprite.png")
}

func getSubimageByName(jsonFile string, jsonName string, spriteSheet string) *ebiten.Image {
	jsonBytes, _ := content.ReadFile(jsonFile)
	jsonMap := map[string]interface{}{}
	_ = json.Unmarshal(jsonBytes, &jsonMap)
	coordinates := jsonMap[jsonName].(map[string]interface{})

	sheet := getEbitenImage(spriteSheet)
	xCoor := int(coordinates["x"].(float64))
	yCoor := int(coordinates["y"].(float64))
	wCoor := int(coordinates["x"].(float64)) + int(coordinates["w"].(float64))
	hCoor := int(coordinates["y"].(float64)) + int(coordinates["h"].(float64))
	rect := image.Rect(xCoor, yCoor, wCoor, hCoor)
	return sheet.SubImage(rect).(*ebiten.Image)
}

func getEbitenImage(name string) *ebiten.Image {
	imgBytes, err := content.ReadFile("images/" + name)
	if err != nil {
		log.Panic()
	}

	pngImage, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		log.Panic()
	}

	return ebiten.NewImageFromImage(pngImage)
}