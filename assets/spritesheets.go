package assets

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/png"
	"log"
)

//go:embed images/spritesheet_hud.png
var hudSpriteSheetBytes []byte

//go:embed images/spritesheet_objects.png
var objectsSpriteSheetBytes []byte

//go:embed images/spritesheet_stall.png
var stallSpriteSheetBytes []byte

//go:embed metadata/metadata_hud.json
var hudMetadataBytes []byte

//go:embed metadata/metadata_objects.json
var objectsMetadataBytes []byte

//go:embed metadata/metadata_stall.json
var stallMetadataBytes []byte

var Hud *SpriteSheet
var Objects *SpriteSheet
var Stall *SpriteSheet

func init() {
	var err error

	Hud, err = NewSpriteSheet(hudSpriteSheetBytes, hudMetadataBytes)
	if err != nil {
		log.Panicln(err)
	}

	Objects, err = NewSpriteSheet(objectsSpriteSheetBytes, objectsMetadataBytes)
	if err != nil {
		log.Panicln(err)
	}

	Stall, err = NewSpriteSheet(stallSpriteSheetBytes, stallMetadataBytes)
	if err != nil {
		log.Panicln(err)
	}
}

type SpriteSheet struct {
	cache    map[string]*ebiten.Image
	image    *ebiten.Image
	metadata map[string]map[string]int
}

func NewSpriteSheet(spriteSheetBytes []byte, metadataBytes []byte) (*SpriteSheet, error) {
	cache := make(map[string]*ebiten.Image)

	pngImage, err := png.Decode(bytes.NewReader(spriteSheetBytes))
	if err != nil {
		return nil, err
	}

	ebitenImage := ebiten.NewImageFromImage(pngImage)

	var metadata map[string]map[string]int
	err = json.Unmarshal(metadataBytes, &metadata)
	if err != nil {
		return nil, err
	}

	spriteSheet := &SpriteSheet{
		cache:    cache,
		image:    ebitenImage,
		metadata: metadata,
	}

	return spriteSheet, nil
}

func (spriteSheet *SpriteSheet) GetSprite(spriteName string) (*ebiten.Image, error) {
	if sprite, ok := spriteSheet.cache[spriteName]; ok {
		return sprite, nil
	}

	x0 := spriteSheet.metadata[spriteName]["x"]
	y0 := spriteSheet.metadata[spriteName]["y"]
	x1 := x0 + spriteSheet.metadata[spriteName]["width"]
	y1 := y0 + spriteSheet.metadata[spriteName]["height"]
	r := image.Rect(x0, y0, x1, y1)

	img := spriteSheet.image.SubImage(r)

	sprite := ebiten.NewImageFromImage(img)

	spriteSheet.cache[spriteName] = sprite

	return sprite, nil
}