# textimg

Library to generate letter icons

## Installation

You just need type:

```shell
go get github.com/senekis/textimg
```

## Usage

```go
import(
  "github.com/senekis/textimg"
  "image/png"
)

func main() {
  textimg.SetFontSize(60)
	rgba := textimg.Generate("Ahoy", image.Rect(0, 0, 300, 200))

	outFile, err := os.Create("ahoy.png")
	if err != nil {
		panic(err)
	}

	defer outFile.Close()
	err = png.Encode(outFile, rgba)
  if err != nil {
		panic(err)
	}
}
```
