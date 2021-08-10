package assets

import (
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"log"
)

//go:embed sounds/assets_ragtime.ogg
var ragtimeMusic []byte

//go:embed sounds/assets_gunshot.wav
var gunshotSound []byte

//go:embed sounds/assets_hit.wav
var hitSound []byte

//go:embed sounds/assets_miss.wav
var missSound []byte

var audioContext *audio.Context
var RagtimeMusicPlayer *audio.Player
var GunshotPlayer *audio.Player
var HitPlayer *audio.Player
var MissPlayer *audio.Player
var sampleRate = 44100

func init() {
	audioContext = audio.NewContext(sampleRate)
	RagtimeMusicPlayer = BackgroundMusicPlayer()
	GunshotPlayer = getWavPlayer(gunshotSound, 0.5)
	HitPlayer = getWavPlayer(hitSound, 0.2)
	MissPlayer = getWavPlayer(missSound, 0.2)
}

func getWavPlayer(sound []byte, volume float64) *audio.Player {
	b := bytes.NewReader(sound)
	s, err := wav.DecodeWithSampleRate(sampleRate, b)
	if err != nil {
		log.Panicln(err)
	}
	p, err := audio.NewPlayer(audioContext, s)
	if err != nil {
		log.Panicln(err)
	}
	p.SetVolume(volume)
	p.Rewind()
	return p
}

func BackgroundMusicPlayer() *audio.Player {
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
	return p
}
