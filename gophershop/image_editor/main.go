package image_editor

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Brightness float64 `json:"brightness"`
	Saturation float64 `json:"saturation"`
	Hue        int     `json:"hue"`
	Blur       float64 `json:"blur"`
	Invert     bool    `json:"invert"`
	Grayscale  bool    `json:"grayscale"`
	Filepath   string  `json:"filepath"`
}

type Editor struct {
	channel *plugin.MethodChannel
}

func (this *Editor) InitPlugin(messenger plugin.BinaryMessenger) error {

	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("editImage", this.parseEdit)
	return nil
}

func (this *Editor) parseEdit(args interface{}) (interface{}, error) {

	dartMsg := args.(string)
	fmt.Println("asdfasdfasdf")
	editorInstance := &EditorInstance{}
	err := json.Unmarshal([]byte(dartMsg), &editorInstance)
	if err != nil {
		return nil, err
	}

	if editorInstance.Filepath != "BOB" {
		return nil, errors.New("WRONG WRONG WRONG")
	}

	outputFileName, err := editorInstance.EditImage()
	if err != nil {
		return nil, err
	}

	return outputFileName, nil
}

func (this *EditorInstance) EditImage() (string, error) {

	fileDir := filepath.Dir(this.Filepath)
	fileName := filepath.Base(this.Filepath)

	fileName = "output_" + fileName
	outputFile := filepath.Join(fileDir, fileName)
	inputFileExt := filepath.Ext(this.Filepath)

	outputFile = strings.Replace(outputFile, inputFileExt, ".png",
		strings.LastIndex(fileName, inputFileExt))

	img, err := imgio.Open(this.Filepath)
	if err != nil {
		return "", err
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
		return "", err
	}

	return outputFile, nil
}
