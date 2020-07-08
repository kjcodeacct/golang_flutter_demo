This directory contains the go-flutter wrapping for the plugin linking to gophershop/image_editor.

If the gophershop/image_editor package has updated, it is pulling from the remote on github.com, and not locally.

To update the package please delete the entry for github.com/kjcodeacct/golang_flutter_demo/gophershop/image_editor in go.mod.
Then run:
```
$ go mod tidy
$ go get github.com/kjcodeacct/golang_flutter_demo/gophershop/image_editor@master
```