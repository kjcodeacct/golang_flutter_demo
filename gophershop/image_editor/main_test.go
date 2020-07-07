package image_editor_test

import (
	"testing"

	image_editor "github.com/kjcodeacct/golang_flutter_demo/gophershop/image_editor"
)

func TestEdit(t *testing.T) {
	testEditInstance := &image_editor.EditorInstance{
		Hue:        -50,
		Saturation: 1.6,
		Invert:     true,
		Filepath:   "./test.png",
	}

	testEditInstance.EditImage(nil)
}
