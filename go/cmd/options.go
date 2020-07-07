package main

import (
	"github.com/go-flutter-desktop/go-flutter"

	image_editor "github.com/kjcodeacct/golang_flutter_demo/gophershop/image_editor"
	file_picker "github.com/miguelpruivo/flutter_file_picker/go"
)

var options = []flutter.Option{
	// set default dimensions of desktop app
	flutter.WindowInitialDimensions(1280, 720),
	// load file picker plugin
	flutter.AddPlugin(&file_picker.FilePickerPlugin{}),
	// load image_editor package created above
	flutter.AddPlugin(&image_editor.EditorInstance{}),
}
