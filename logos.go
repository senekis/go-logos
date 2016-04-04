package main

import (
	"bufio"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/senekis/logos/textimg"
)

var (
	// Logger set the logger for debugging
	Logger = log.New(os.Stderr, "[logos] ", log.LstdFlags)
)

func main() {
	fmt.Print("Enter text to convert: ")

	var text string
	n, err := fmt.Scanf("%s\n", &text)
	if err != nil {
		fmt.Println(n, err)
	}

	text = strings.ToUpper(text)

	textimg.Logger = Logger
	rgba := textimg.Generate(text)

	saveImage(rgba, "basic_image.png")

	firstLetter := string(text[0])
	textimg.SetImageHeight(81)
	textimg.SetFontSize(60)
	rgba = textimg.Generate(firstLetter)

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
