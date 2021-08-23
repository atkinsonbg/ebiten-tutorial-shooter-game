package assets

import (
	_ "embed"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//go:embed fonts/Retrophile_Regular.ttf
var customFont []byte

var Font font.Face

func init() {
	tt, _ := truetype.Parse(customFont)
	Font = truetype.NewFace(tt, &truetype.Options{
		Size: 36,
		DPI: 72,
		Hinting: font.HintingFull,
	})
}


