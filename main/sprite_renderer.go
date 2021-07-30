package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container *element
	tex       *sdl.Texture

	width, height int32
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string, width, height int32) *spriteRenderer {
	return &spriteRenderer{
		container: container,
		tex:       textureFromPNG(renderer, filename),
		width:     width,
		height:    height,
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	_, _, texWidth, texHeight, err := sr.tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	// Converting coordinates to top left of sprite
	x := sr.container.position.x - float64(sr.width)/2.0
	y := sr.container.position.y - float64(sr.height)/2.0

	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: texWidth, H: texHeight},
		&sdl.Rect{X: int32(x), Y: int32(y), W: sr.width, H: sr.height},
		sr.container.rotation,
		&sdl.Point{X: sr.width / 2, Y: sr.height / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func textureFromPNG(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := img.Load(filename)
	if err != nil {
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}

	return tex
}

func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
