package utils

import (
	"github.com/atkinsonbg/ebiten-tutorial/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const debouncer = 600 * time.Millisecond

func init() {

}

func CheckMouseClick(lastClickedAt time.Time) (time.Time, int, int) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && time.Now().Sub(lastClickedAt) > debouncer {
		assets.GunshotPlayer.Play()
		assets.GunshotPlayer.Rewind()
		x, y := ebiten.CursorPosition()
		return time.Now(), x, y
	}

	return lastClickedAt, 0, 0
}