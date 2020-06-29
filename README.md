
---
# Getting Started

* Install go 1.13+
    * instructions can be found at <https://golang.org/doc/install>
* Install flutter
    * instructions can be found at <https://flutter.dev/docs/get-started/install>
* Install go-flutter
    * install hover (flutters unofficial go client)
        * ```GO111MODULE=on go get -u -a github.com/go-flutter-desktop/hover```
* Install dependencies
    * mac - have xcode installed
    * linux - install the following packages 'libgl1-mesa-dev xorg-dev'
    * windows - install gflw for windows <https://www.glfw.org/docs/latest/compile.html#compile_deps>
* Enable desktop builds
    ```
    flutter config --enable-macos-desktop
    flutter config --enable-linux-desktop
    flutter config --enable-windows-desktop
    ```

---