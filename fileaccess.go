package main

import (
	"image"
	_ "image/png"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

func GetImageFromFilePath(filePath string) *ebiten.Image {
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}

	defer f.Close()
	image, _, err := image.Decode(f)

	if err != nil {
		return nil
	}

	return ebiten.NewImageFromImage(image)
}