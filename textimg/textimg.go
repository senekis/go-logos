package textimg

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var (
	fontPath = "textimg/Roboto-Regular.ttf"
	// font size in points
	fontSize float64 = 100
	// line spacing (e.g. 2 means double spaced)
	fontSpacing = 1.2
	fontToUse   *truetype.Font

	imageHeight         = 120
	dpi         float64 = 72
	// Logger set the logger
	Logger = log.New(os.Stderr, "[textimg] ", log.LstdFlags)
)

// SetFontPath allow to configure the font file to use
func SetFontPath(path string) {
	fontPath = path
}

// SetFontSize allow to configure the font size in points
func SetFontSize(size float64) {
	fontSize = size
}

// SetFontSpacing allow to configure the line spacing
func SetFontSpacing(spacing float64) {
	fontSpacing = spacing
}

// SetImageHeight configure the image height
func SetImageHeight(height int) {
	imageHeight = height
}

// SetDPI configure screen resolution in Dots Per Inch
func SetDPI(dpiValue float64) {
	dpi = dpiValue
}

func loadFont() {
	fontBytes, err := ioutil.ReadFile(fontPath)

	if err != nil {
		Logger.Fatal(err)
	}

	fontToUse, err = freetype.ParseFont(fontBytes)
	if err != nil {
		Logger.Fatal(err)
	}
}

// Generate convert the text into an image
func Generate(text string) *image.RGBA {
	imageWidth := calculateWidth(text)

	// Initialize the context.
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(fontToUse)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	pt := freetype.Pt(0, int(c.PointToFixed(fontSize)>>6))
	_, err := c.DrawString(text, pt)
	if err != nil {
		Logger.Fatal(err)
	}

	return rgba
}

func calculateWidth(text string) int {
	// Truetype stuff
	opts := truetype.Options{}
	opts.Size = fontSize
	face := truetype.NewFace(fontToUse, &opts)

	spacing64 := fontSpacing * 64.0
	width64 := 0.0

	for _, x := range text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			Logger.Print("Error getting text width")
		}
		iwidthf := float64(awidth) + spacing64
		width64 = width64 + iwidthf
	}

	return int(width64 / 64)
}

func init() {
	loadFont()
}
