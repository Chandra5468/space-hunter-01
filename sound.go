package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

const sampleRate = 44100

func loadSound(ctx *audio.Context, path string) *audio.Player {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	d, err := vorbis.DecodeWithSampleRate(sampleRate, f)
	if err != nil {
		log.Fatal(err)
	}

	p, err := ctx.NewPlayer(d)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

func playSound(p *audio.Player) {
	// if p.IsPlaying() {
	p.Rewind()
	// }
	p.Play()
}
