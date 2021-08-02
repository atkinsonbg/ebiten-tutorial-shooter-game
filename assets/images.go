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
	return getSubimageByName("curtain")
}

func CurtainStraight() *ebiten.Image {
	return getSubimageByName("curtain_straight")
}

func BackgroundGreen() *ebiten.Image {
	return getSubimageByName("bg_green")
}

func Wood() *ebiten.Image {
	return getSubimageByName("bg_wood")
}

func Wave() *ebiten.Image {
	return getSubimageByName("wave")
}

func getSubimageByName(name string) *ebiten.Image {
	jsonBytes, _ := content.ReadFile("images/stall.json")
	jsonMap := map[string]interface{}{}
	_ = json.Unmarshal(jsonBytes, &jsonMap)
	coordinates := jsonMap[name].(map[string]interface{})

	spriteSheet := getEbitenImage("stall_sprite.png")
	xCoor := int(coordinates["x"].(float64))
	yCoor := int(coordinates["y"].(float64))
	wCoor := int(coordinates["x"].(float64)) + int(coordinates["w"].(float64))
	hCoor := int(coordinates["y"].(float64)) + int(coordinates["h"].(float64))
	rect := image.Rect(xCoor, yCoor, wCoor, hCoor)
	return spriteSheet.SubImage(rect).(*ebiten.Image)
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