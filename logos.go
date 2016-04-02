package main

import (
	"bufio"
	"fmt"
	"image"
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

	saveImage(rgba, "out.png")
}

// saveImage save image in PNG format
func saveImage(rgba *image.RGBA, outputFilename string) {
	// Save that RGBA image to disk.
	outFile, err := os.Create(outputFilename)
	if err != nil {
		Logger.Fatal(err)
	}

	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		Logger.Fatal(err)
	}

	err = b.Flush()
	if err != nil {
		Logger.Fatal(err)
	}
}
