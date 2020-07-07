package image_editor_test

import (
	image_editor "gohershop/image"
	"testing"
)

func TestEdit(t *testing.T) {
	testEditInstance := &image_editor.EditorInstance{
		Hue:        -50,
		Saturation: 1.6,
		Invert:     true,
		Filepath:   "./test.png",
	}

	testEditInstance.EditImage()
}
