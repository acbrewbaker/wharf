package tlc_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func must(t *testing.T, err error) {
	if err != nil {
		t.Error("must failed: ", err.Error())
		t.FailNow()
	}
}

type regEntry struct {
	Path string
	Size int
	Byte byte
}

type symlinkEntry struct {
	Oldname string
	Newname string
}

var regulars = []regEntry{
	{"foo/file_f", 50, 0xd},
	{"foo/dir_a/baz", 10, 0xa},
	{"foo/dir_b/zoom", 30, 0xc},
	{"foo/file_z", 40, 0xe},
	{"foo/dir_a/bazzz", 20, 0xb},
}

var symlinks = []symlinkEntry{
	{"file_z", "foo/file_m"},
	{"dir_a/baz", "foo/file_o"},
}

func mktestdir(t *testing.T, name string) string {
	tmpPath, err := ioutil.TempDir(".", "tmp_"+name)
	must(t, err)

	must(t, os.RemoveAll(tmpPath))

	for _, entry := range regulars {
		fullPath := filepath.Join(tmpPath, entry.Path)
		must(t, os.MkdirAll(filepath.Dir(fullPath), os.FileMode(0777)))
		file, err := os.Create(fullPath)
		must(t, err)

		filler := []byte{entry.Byte}
		for i := 0; i < entry.Size; i++ {
			_, err := file.Write(filler)
			must(t, err)
		}
		must(t, file.Close())
	}

	for _, entry := range symlinks {
		new := filepath.Join(tmpPath, entry.Newname)
		must(t, os.Symlink(entry.Oldname, new))
	}

	return tmpPath
}
