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
		Name        string
		Body        string
		AccessNames []string
	}{
		{
			Name: "readme.txt",
			Body: "This archive contains some text files.",
			AccessNames: []string{
				"readme.txt",
				"/readme.txt",
				"./readme.txt",
				"././readme.txt",
				"../readme.txt",
			},
		},
		{
			Name: "/gopher.txt",
			Body: "Gopher names:\nGeorge\nGeoffrey\nGonzo",
			AccessNames: []string{
				"gopher.txt",
				"/gopher.txt",
				"./gopher.txt",
				"././gopher.txt",
				"../gopher.txt",
			},
		},
		{
			Name: "./todo.txt",
			Body: "Get animal handling licence.",
			AccessNames: []string{
				"todo.txt",
				"/todo.txt",
				"./todo.txt",
				"././todo.txt",
				"../todo.txt",
			},
		},
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
		for _, path := range file.AccessNames {
			f, err := fs.Open(path)
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
}
