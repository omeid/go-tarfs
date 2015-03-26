// In memory http.FileSystem from tar archives
package tarfs

import (
	"archive/tar"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


// New returns an http.FileSystem that holds all the files in the tar,
// It reads the whole archive from the Reader. It is the caller's responsibility to call Close on the Reader when done.
func New(tarstream io.Reader) (http.FileSystem, error) {
	tr := tar.NewReader(tarstream)

	tarfs := tarfs{make(map[string]file)}
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			return nil, err
		}
		data, err := ioutil.ReadAll(tr)
		if err != nil {
			return nil, err
		}

		tarfs.files[hdr.Name] = file{data: data, fi: hdr.FileInfo()}
	}
	return &tarfs, nil
}

type file struct {
	*bytes.Reader
	data []byte
	fi   os.FileInfo

	files []os.FileInfo
}

type tarfs struct {
	files map[string]file
}

func (tf *tarfs) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.IndexRune(name, filepath.Separator) >= 0 ||
		strings.Contains(name, "\x00") {
		return nil, errors.New("http: invalid character in file path")
	}
	f, ok := tf.files[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	if f.fi.IsDir() {
		f.files = []os.FileInfo{}
		for path, file := range tf.files {
			if strings.HasPrefix(path, name) {
				s, _ := file.Stat()
				f.files = append(f.files, s)
			}
		}

	}
	f.Reader = bytes.NewReader(f.data)
	return &f, nil
}

// A noop-closer.
func (f *file) Close() error {
	return nil
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	if f.fi.IsDir() && f.files != nil {
		return f.files, nil
	}
	return nil, os.ErrNotExist
}

func (f *file) Stat() (os.FileInfo, error) {
	return f.fi, nil
}
