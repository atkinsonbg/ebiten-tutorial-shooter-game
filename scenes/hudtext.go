package scenes

import (
	"fmt"
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/atkinsonbg/ebiten-tutorial/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"strconv"
)

func HudTextScene(screen *ebiten.Image, windowWidth int, windowHeight int, currentScore utils.Score) {
	s := strconv.Itoa(currentScore.GetScore())
	t := fmt.Sprintf("SCORE: %s", s)
	bounds := text.BoundString(assets.Font, t)
	text.Draw(screen, t, assets.Font, 10, windowHeight - bounds.Dy(), color.White)
}