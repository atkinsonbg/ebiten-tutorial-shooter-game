package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"log"
)

//go:embed sounds/assets_ragtime.ogg
var ragtimeMusic []byte

var audioContext *audio.Context
var RagtimeMusicPlayer *audio.Player
var sampleRate = 44100

func init() {
	audioContext = audio.NewContext(sampleRate)
	RagtimeMusicPlayer, _ = BackgroundMusicPlayer()
}

func BackgroundMusicPlayer() (*audio.Player, error) {
	b := bytes.NewReader(ragtimeMusic)
	oggS, err := vorbis.DecodeWithSampleRate(sampleRate, b)
	if err != nil {
		log.Panicln(err)
	}
	l := audio.NewInfiniteLoop(oggS, oggS.Length())
	p, err := audio.NewPlayer(audioContext, l)
	if err != nil {
		log.Panicln(err)
	}
	p.SetVolume(0.1)
	return p, nil
}
