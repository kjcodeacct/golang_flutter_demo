package image_editor

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

// channelName is a reference to the go package used for this plugin
const channelName = "gophershop/image_editor"

type EditorInstance struct {
	Brightness float64
	Saturation float64
	Hue        int
	Blur       float64
	Invert     bool
	Grayscale  bool
	Filepath   string
}

func (this *EditorInstance) InitPlugin(messenger plugin.BinaryMessenger) error {

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("editImage", this.EditImage)
	return nil
}

func (this *EditorInstance) EditImage(arguments interface{}) (reply interface{}, err error) {

	fileDir := filepath.Dir(this.Filepath)
	fileName := filepath.Base(this.Filepath)

	fileName = "output_" + fileName
	outputFile := filepath.Join(fileDir, fileName)
	inputFileExt := filepath.Ext(this.Filepath)

	outputFile = strings.Replace(outputFile, inputFileExt, ".png",
		strings.LastIndex(fileName, inputFileExt))

	img, err := imgio.Open(this.Filepath)
	if err != nil {
		log.Println(err)
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

	if this.Saturation > 0 {
		img = adjust.Saturation(img, this.Saturation)
	}

	err = imgio.Save(outputFile, img, imgio.PNGEncoder())
	if err != nil {
		log.Println(err)
		return
	}

	return nil, nil
}
