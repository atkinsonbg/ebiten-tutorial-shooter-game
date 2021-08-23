package utils

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"runtime"
)

func PrintMemUsage(screen *ebiten.Image) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	cpx, cpy := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%v\nAlloc = %v MiB\nTotalAlloc = %v MiB\nSys = %v MiB\nNumGC = %v\nCursor, X= %v, Y= %v", ebiten.CurrentFPS(), bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC, cpx, cpy))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
