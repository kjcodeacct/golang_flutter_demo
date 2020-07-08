import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:file_picker/file_picker.dart';
import 'package:unicorndial/unicorndial.dart';
import 'dart:io';
import 'dart:developer' as developer;

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        // This is the theme of your application.
        //
        // Try running your application with "flutter run". You'll see the
        // application has a blue toolbar. Then, without quitting the app, try
        // changing the primarySwatch below to Colors.green and then invoke
        // "hot reload" (press "r" in the console where you ran "flutter run",
        // or simply save your changes to "hot reload" in a Flutter IDE).
        // Notice that the counter didn't reset back to zero; the application
        // is not restarted.
        primarySwatch: Colors.blue,
      ),
      home: MyHomePage(title: 'Gophershop'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  // This widget is the home page of your application. It is stateful, meaning
  // that it has a State object (defined below) that contains fields that affect
  // how it looks.

  // This class is the configuration for the state. It holds the values (in this
  // case the title) provided by the parent (in this case the App widget) and
  // used by the build method of the State. Fields in a Widget subclass are
  // always marked "final".

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  // file handling
  // TODO - break file handling out into mixin or other state class
  final GlobalKey<ScaffoldState> _scaffoldKey = GlobalKey<ScaffoldState>();
  String loadedImagePath = 'assets/default.png';
  String originalImagePath;
  File loadedImage;

  String workingImagePath;
  String _path;
  Map<String, String> _paths;
  String _extension;
  bool _loadingPath = false;
  bool _multiPick = false;
  FileType _pickingType = FileType.any;

  void _openFileExplorer() async {
    setState(() => _loadingPath = true);
    try {
      if (_multiPick) {
        _path = null;
        _paths = await FilePicker.getMultiFilePath(
            type: _pickingType,
            allowedExtensions: (_extension?.isNotEmpty ?? false)
                ? _extension?.replaceAll(' ', '')?.split(',')
                : null);
      } else {
        _paths = null;
        _path = await FilePicker.getFilePath(
            type: _pickingType,
            allowedExtensions: (_extension?.isNotEmpty ?? false)
                ? _extension?.replaceAll(' ', '')?.split(',')
                : null);
      }
    } on PlatformException catch (e) {
      print("Unsupported operation" + e.toString());
    }

    if (!mounted) return;
    setState(() {
      _loadingPath = false;

      if (_path != null) {
        loadedImagePath = _path;
      }

      originalImagePath = loadedImagePath;
      print("loading image $loadedImagePath");
    });
  }

  // image processing value placekeepers
  double brightnessVal = 0;
  double saturationVal = 0;
  double hueVal = 0;
  double blurVal = 0;
  bool invertImg = false;
  bool grayScale = false;

  static const image_editor_lib = const MethodChannel('gophershop/image_editor');

  void _editImage() async {
    print("opening and editing image...");
    // this maps directly to the editorInstance found in gophershop/image_editor/main.go:19
    var editorInstance = {
      "brightness": brightnessVal,
      "saturation": saturationVal,
      "hue": hueVal,
      "blur": blurVal,
      "invert": invertImg,
      "grayscale": grayScale,
      "filepath": originalImagePath
    };

    var jsonText = jsonEncode(editorInstance);
    try {
      loadedImagePath = await image_editor_lib.invokeMethod("editImage", jsonText);
    } on PlatformException catch (e) {
      print("error:" + e.toString());
    } finally {
      print("$loadedImagePath edited...");
      loadedImage = new File(loadedImagePath);
    }
  }

  var globalScrollDirection = Axis.horizontal;

  @override
  Widget build(BuildContext context) {
    // This method is rerun every time setState is called, for instance as done on _openFileExplorer
    //
    // The Flutter framework has been optimized to make rerunning build methods
    // fast, so that you can just rebuild anything that needs updating rather
    // than having to individually change instances of widgets.

    var appSize = MediaQuery.of(context).size.width;
    if (appSize >= 720) {
      globalScrollDirection = Axis.horizontal;
    } else {
      globalScrollDirection = Axis.vertical;
    }

    var menuButton = List<UnicornButton>();

    menuButton.add(UnicornButton(
        hasLabel: true,
        labelText: "Save",
        currentButton: FloatingActionButton(
          heroTag: "save",
          backgroundColor: Colors.greenAccent,
          mini: true,
          child: Icon(Icons.save_alt),
          onPressed: () => _openFileExplorer(),
        )));

    menuButton.add(UnicornButton(
        hasLabel: true,
        labelText: "Open",
        currentButton: FloatingActionButton(
          heroTag: "open",
          backgroundColor: Colors.grey,
          mini: true,
          child: Icon(Icons.folder_open),
          onPressed: () => _openFileExplorer(),
        )));

    // call image editing plugin
    if (originalImagePath != null) {
      _editImage();
    }
    // TODO reload image and not re-open
    // pause for a moment to let the edited image write
    sleep(const Duration(milliseconds: 500));
    loadedImage = new File(loadedImagePath);

    setState(() {
      print("refreshing image...");
      imageCache.clear();
      imageCache.clearLiveImages();
    });

    // rebuild ui
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),
      ),
      body: Center(
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              Flexible(
                flex: 1,
                child: Text('File open $loadedImagePath',
                    style: TextStyle(fontWeight: FontWeight.bold)),
              ),
              Flexible(
                  flex: 12,
                  child: Container(
                    margin: const EdgeInsets.all(30.0),
                    padding: const EdgeInsets.all(10.0),
                    decoration: BoxDecoration(
                        border: Border.all(
                          color: Theme.of(context).dividerColor,
                          // width: 5.0,
                        ),
                        borderRadius: BorderRadius.all(Radius.circular(5.0))),
                    child: Image.file(loadedImage),
                  )),
              Flexible(
                flex: 4,
                child: ListView(
                  scrollDirection: globalScrollDirection,
                  shrinkWrap: true,
                  children: [
                    Text(
                      'Brightness',
                    ),
                    Slider(
                      value: brightnessVal,
                      onChanged: (value) {
                        setState(() => brightnessVal = value);
                      },
                      min: 0,
                      max: 10,
                    ),
                    Text(
                      'Hue',
                    ),
                    Slider(
                      value: hueVal,
                      onChanged: (value) {
                        setState(() => hueVal = value);
                      },
                      min: -100,
                      max: 100,
                    ),
                    Text(
                      'Blur',
                    ),
                    Slider(
                      value: blurVal,
                      onChanged: (value) {
                        setState(() => blurVal = value);
                      },
                      min: 0,
                      max: 10,
                    ),
                    Text(
                      'Saturation',
                    ),
                    Slider(
                      value: saturationVal,
                      onChanged: (value) {
                        setState(() => saturationVal = value);
                      },
                      min: 0,
                      max: 10,
                    ),
                  ],
                ),
              ),
              Flexible(
                flex: 4,
                child: ListView(
                  scrollDirection: Axis.vertical,
                  children: [
                    CheckboxListTile(
                      title: Text("Invert Image"),
                      value: invertImg,
                      onChanged: (newValue) {
                        setState(() {
                          invertImg = newValue;
                        });
                      },
                      controlAffinity: ListTileControlAffinity.leading,
                    ),
                    CheckboxListTile(
                      title: Text("Grayscale Image"),
                      value: grayScale,
                      onChanged: (newValue) {
                        setState(() {
                          grayScale = newValue;
                        });
                      },
                      controlAffinity: ListTileControlAffinity.leading,
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
      floatingActionButton: UnicornDialer(
        orientation: UnicornOrientation.VERTICAL,
        parentButton: Icon(Icons.menu),
        childButtons: menuButton,
        animationDuration: 0,
      ),
    );
  }
}
