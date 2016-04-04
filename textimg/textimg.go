package textimg

import (
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	fontPath = "textimg/Roboto-Regular.ttf"
	// font size in points
	fontSize float64 = 100
	// line spacing (e.g. 2 means double spaced)
	lineSpacing = 1.2
	fontToUse   *truetype.Font

	dpi  float64 = 72
	face font.Face
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

// SetLineSpacing allow to configure the line spacing
func SetLineSpacing(spacing float64) {
	lineSpacing = spacing
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

	face = truetype.NewFace(fontToUse, &truetype.Options{Size: fontSize})
}

// Generate convert the text into an image
func Generate(text string, r image.Rectangle) *image.RGBA {
	// Initialize the context.
	fg, bg := image.Black, image.White
	rgba := image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	drawer := &font.Drawer{
		Dst:  rgba,
		Src:  fg,
		Face: face,
	}

	x := (fixed.I(r.Dx()) - drawer.MeasureString(text)) / 2
	asc, desc := getHeights(text)

	y := ((fixed.I(r.Dy()) - (asc + desc)) / 2.0) + asc

	drawer.Dot = fixed.Point26_6{
		X: x,
		Y: y,
	}

	drawer.DrawString(text)
	return rgba
}

func getHeights(text string) (fixed.Int26_6, fixed.Int26_6) {
	asc := fixed.I(0)
	desc := fixed.I(0)

	for _, x := range text {
		bounds, _, ok := face.GlyphBounds(rune(x))
		if ok {
			minY := -bounds.Min.Y
			maxY := bounds.Max.Y

			if minY > asc {
				asc = minY
			}

			if maxY > desc {
				desc = maxY
			}
		}
	}

	return asc, desc
}

func init() {
	loadFont()
}
