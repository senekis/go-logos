package main

import (
	"bufio"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/senekis/textimg"
)

var (
	logger = log.New(os.Stderr, "[lettericons] ", log.LstdFlags)
)

func main() {
	text := os.Args[1]
	text = strings.ToUpper(text)

	textimg.Logger = logger
	textimg.SetFontSize(60)
	rgba := textimg.Generate(text, image.Rect(0, 0, 200, 200))

	saveImage(rgba, "basic_image.png")

	rounded(rgba, 50)
	saveImage(rgba, "basic_image_round.png")

	firstLetter := string(text[0])
	rgba = textimg.Generate(firstLetter, image.Rect(0, 0, 80, 80))
	applyMaskFile(rgba, "example/mask.png")
}

func rounded(img *image.RGBA, radius int) {
	mask := image.NewRGBA(img.Bounds())
	gc := draw2dimg.NewGraphicContext(mask)

	gc.SetFillColor(image.Black)
	gc.SetStrokeColor(image.Transparent)
	draw2dkit.RoundedRectangle(gc, float64(mask.Rect.Min.X), float64(mask.Rect.Min.Y), float64(mask.Rect.Max.X), float64(mask.Rect.Max.Y), float64(radius), float64(radius))
	gc.FillStroke()

	draw.DrawMask(img, img.Bounds(), img, image.ZP, mask, mask.Rect.Min, draw.Src)
}

func applyMaskFile(rgba image.Image, maskPath string) {
	file, err := os.Open(maskPath)
	if err != nil {
		logger.Fatal(err)
	}

	mask, err := png.Decode(bufio.NewReader(file))
	if err != nil {
		logger.Fatal(err)
	}

	maskRgba := mask.(*image.NRGBA)
	point := image.Point{
		X: (rgba.Bounds().Dx() / 2) - (maskRgba.Bounds().Dx() / 2),
		Y: 0,
	}

	draw.Draw(maskRgba, mask.Bounds(), rgba, point, draw.Over)
	saveImage(maskRgba, "first_letter.png")
}

// saveImage save image to disk
func saveImage(img image.Image, outputFilename string) {
	outFile, err := os.Create(outputFilename)
	if err != nil {
		logger.Fatal(err)
	}

	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, img)
	if err != nil {
		logger.Fatal(err)
	}

	err = b.Flush()
	if err != nil {
		logger.Fatal(err)
	}
}
