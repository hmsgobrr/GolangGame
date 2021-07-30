package main

import "github.com/veandco/go-sdl2/sdl"

type animator struct {
	container *element
	sequences map[string]*sequence
	current   string
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string) *animator {
	return &animator{
		container: container,
		sequences: sequences,
		current:   defaultSequence,
	}
}

// func (an *animator) onDraw(renderer *sdl.Renderer) error {

// }

type sequence struct {
	texture     []*sdl.Texture
	frame       int
	samplerRate float64
}
