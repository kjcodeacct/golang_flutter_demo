package image_editor

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFileCleanup(t *testing.T) {

	editFileDir := filepath.Dir("./")
	err := filepath.Walk(editFileDir, cleanupEditorFiles)
	require.NoError(t, err)

}
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
