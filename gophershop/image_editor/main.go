package image_editor

import (
	"fmt"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
)

type EditorInstance struct {
	Brightness float64
	Hue        int
	Blur       float64
	Invert     bool
	Grayscale  bool
	Filepath   string
}

func (this *EditorInstance) EditImage() {

	img, err := imgio.Open(this.Filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if this.Invert {
		img = effect.Invert(img)
	}

	if this.Grayscale {
		img = effect.Grayscale(img)
	}

	if this.Blur > 0 {
		img = blur.Gaussian(img, this.Blur)
	}

	if this.Brightness > 0 {
		img = adjust.Brightness(img, this.Brightness)
	}

	if this.Hue != 0 {
		img = adjust.Hue(img, this.Hue)
	}

	err = imgio.Save("output.png", img, imgio.PNGEncoder())
	if err != nil {
		fmt.Println(err)
		return
	}

}
