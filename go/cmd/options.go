package main

import (
	"github.com/go-flutter-desktop/go-flutter"

	file_picker "github.com/miguelpruivo/flutter_file_picker/go"
)

var options = []flutter.Option{
	// set default dimensions of desktop app
	flutter.WindowInitialDimensions(1280, 720),
	// load file picker plugin
	flutter.AddPlugin(&file_picker.FilePickerPlugin{}),
}
