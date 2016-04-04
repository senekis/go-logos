package main

import (
	"bufio"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/senekis/logos/textimg"
)

var (
	// Logger set the logger for debugging
	Logger = log.New(os.Stderr, "[logos] ", log.LstdFlags)
)

func main() {
	text := os.Args[1]

	textimg.Logger = Logger
	textimg.SetFontSize(60)

	// text = strings.ToUpper(text)
	rgba := textimg.Generate(text, image.Rect(0, 0, 200, 200))

	saveImage(rgba, "basic_image.png")

	firstLetter := string(text[0])
	rgba = textimg.Generate(firstLetter, image.Rect(0, 0, 80, 80))

	file, err := os.Open("mask.png")
	if err != nil {
		Logger.Fatal(err)
	}

	mask, err := png.Decode(bufio.NewReader(file))
	if err != nil {
		Logger.Fatal(err)
	}

	maskRgba := mask.(*image.NRGBA)
	point := image.Point{
		X: (rgba.Bounds().Dx() / 2) - (maskRgba.Bounds().Dx() / 2),
		Y: 0,
	}

	draw.Draw(maskRgba, mask.Bounds(), rgba, point, draw.Over)
	saveImage(maskRgba, "first_letter.png")
}

// saveImage save image in PNG format
func saveImage(img image.Image, outputFilename string) {
	// Save that RGBA image to disk.
	outFile, err := os.Create(outputFilename)
	if err != nil {
		Logger.Fatal(err)
	}

	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		Logger.Fatal(err)
	}

	err = b.Flush()
	if err != nil {
		Logger.Fatal(err)
	}
}
