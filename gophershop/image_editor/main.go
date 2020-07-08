package image_editor

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pborman/uuid"
)

// channelName is a reference to the go package used for this plugin
const channelName = "gophershop/image_editor"
const editorFilePrefix = "gophershop_output_"

// this could be handled alot nicer, but is convenient
var currentFile = ""

type EditorInstance struct {
	Brightness float64 `json:"brightness"`
	Saturation float64 `json:"saturation"`
	Hue        float64 `json:"hue"`
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
	log.Println("golang:", "received", dartMsg)
	editorInstance := &EditorInstance{}
	err := json.Unmarshal([]byte(dartMsg), &editorInstance)
	if err != nil {
		log.Println("golang:", err.Error())
		return nil, err
	}

	// cleanup previously edited image, since it is cached in the flutter instance
	editFileDir := filepath.Dir(editorInstance.Filepath)
	err = filepath.Walk(editFileDir, cleanupEditorFiles)
	if err != nil {
		return nil, err
	}

	outputFilePath, err := editorInstance.EditImage()
	if err != nil {
		log.Println("golang:", err.Error())
		return nil, err
	}

	return outputFilePath, nil
}

func (this *EditorInstance) EditImage() (string, error) {

	fileDir := filepath.Dir(this.Filepath)
	fileName := filepath.Base(this.Filepath)

	fileUUID := uuid.New()
	fileName = editorFilePrefix + fileUUID + "_" + fileName
	outputFilePath := filepath.Join(fileDir, fileName)
	inputFileExt := filepath.Ext(this.Filepath)

	outputFilePath = strings.Replace(outputFilePath, inputFileExt, ".png",
		strings.LastIndex(fileName, inputFileExt))

	if _, err := os.Stat(outputFilePath); err == nil {
		err := os.Remove(outputFilePath)

		if err != nil {
			return "", err
		}
	}

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
		img = adjust.Hue(img, int(this.Hue))
	}

	if this.Saturation > 0 {
		img = adjust.Saturation(img, this.Saturation)
	}

	err = imgio.Save(outputFilePath, img, imgio.PNGEncoder())
	if err != nil {
		return "", err
	}

	return outputFilePath, nil
}

func cleanupEditorFiles(path string, f os.FileInfo, err error) error {

	if strings.HasPrefix(f.Name(), editorFilePrefix) {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return nil
}
