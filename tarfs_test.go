package tarfs

import (
	"archive/tar"
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

func TestOpen(t *testing.T) {

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	// Open the tar archive for reading.
	fs, err := New(bytes.NewReader(buf.Bytes()))
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		f, err := fs.Open(file.Name)
		if err != nil {
			t.Fatal(err)
		}
		content, _ := ioutil.ReadAll(f)
		if string(content) != file.Body {
			t.Fatalf("For '%s'\nExpected:\n%s\nGot:\n%s\n", file.Name, file.Body, content)
		}

		var (
			s, _ = f.Stat()
			size = int64(len(file.Body))
			got  = s.Size()
		)

		if size != got {
			t.Fatalf("For '%s'\nExpected Size:\n%v\nGot:\n%v\n", file.Name, size, got)
		}
	}
}
