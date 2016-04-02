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

// TextConf set font conf, dpi and image height
type TextConf struct {
	fontPath    string
	fontSize    float64
	fontSpacing float64
	imageHeight int
	dpi         float64
}

var (
	config    TextConf
	fontToUse *truetype.Font

	// Logger set the logger
	Logger = log.New(os.Stderr, "[textimg] ", log.LstdFlags)
)

// Setup configuration for Font, imageHeight and dpi
func Setup(cfg TextConf) {
	fontBytes, err := ioutil.ReadFile(cfg.fontPath)

	if err != nil {
		Logger.Fatal(err)
	}

	fontToUse, err = freetype.ParseFont(fontBytes)
	if err != nil {
		Logger.Fatal(err)
	}

	config = cfg
}

// Generate convert the text into an image
func Generate(text string) *image.RGBA {
	imageWidth := calculateWidth(text)

	// Initialize the context.
	fg, bg := image.Black, image.Transparent
	rgba := image.NewRGBA(image.Rect(0, 0, imageWidth, config.imageHeight))

	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(config.dpi)
	c.SetFont(fontToUse)
	c.SetFontSize(config.fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(config.fontSize)>>6))
	_, err := c.DrawString(text, pt)
	if err != nil {
		Logger.Fatal(err)
	}

	return rgba
}

func calculateWidth(text string) int {
	// Truetype stuff
	opts := truetype.Options{}
	opts.Size = config.fontSize
	face := truetype.NewFace(fontToUse, &opts)

	spacing64 := config.fontSpacing * 64.0
	width64 := spacing64 * 2.0

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
	Setup(TextConf{"textimg/Roboto-Regular.ttf", 100, 1.2, 120, 72})
}
