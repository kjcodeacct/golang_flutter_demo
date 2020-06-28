
---
# Getting Started

* Install go 1.13+
    * instructions can be found at <https://golang.org/doc/install>
* Install flutter
    * instructions can be found at <https://flutter.dev/docs/get-started/install>
* Install go-flutter
    * install hover (flutters unofficial go client)
        * > GO111MODULE=on go get -u -a github.com/go-flutter-desktop/hover
    * mac - have xcode installed
        * enable flutter to build for mac
            * > flutter config --enable-macos-desktop
    * linux - install the following packages 'libgl1-mesa-dev xorg-dev'
        * enable flutter to build for linux
            * > flutter config --enable-linux-desktop
    * windows - install gflw for windows <https://www.glfw.org/docs/latest/compile.html#compile_deps>
        * enabled flutter to build for windows
            * > flutter config --enable-windows-desktop

---