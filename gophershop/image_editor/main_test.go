package image_editor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEdit(t *testing.T) {
	testEditInstance := &EditorInstance{
		Hue:        -50,
		Saturation: 1.6,
		Invert:     true,
		Filepath:   "./test.png",
	}

	testFilePath, err := testEditInstance.EditImage()
	require.NoError(t, err)

	t.Log("created file:", testFilePath)
}
