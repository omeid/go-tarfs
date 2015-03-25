[![GoDoc](https://godoc.org/github.com/omeid/go-tarfs?status.svg)](https://godoc.org/github.com/omeid/go-tarfs) [![Build Status](https://drone.io/github.com/omeid/go-tarfs/status.png)](https://drone.io/github.com/omeid/go-tarfs/latest)
# tarfs
In-memory http.FileSystem from tar archives.

### Why?
If you have multiple assets for your program that you don't want to [embed](https://github.com/omeid/go-resources) them in your binary but still want an easier way to ship them along your binary, tarfs is your friend.

### Usage
See the [GoDoc](https://godoc.org/github.com/omeid/go-tarfs)


### Contributing
Please consider opening an issue first, or just send a pull request. :)

### Credits
See [Contributors](https://github.com/omeid/go-tarfs/graphs/contributors).

### LICENSE
  [MIT](LICENSE).

### TODO
  - Add more tests

### SEE ALSO
  - [github.com/omeid/go-resources](http://godoc.org/github.com/omeid/go-resources)
  - [x/tools/godoc/vfs/zipfs](http://godoc.org/golang.org/x/tools/godoc/vfs/zipfs)
  - [github.com/tsuru/tusru/fs](http://godoc.org/github.com/tsuru/tsuru/fs)
