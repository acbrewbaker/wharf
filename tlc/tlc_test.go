package tlc

import (
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/itchio/arkive/zip"

	"github.com/itchio/wharf/archiver"
	"github.com/itchio/wharf/state"
	"github.com/stretchr/testify/assert"
)

func Test_NonDirWalk(t *testing.T) {
	tmpPath, err := ioutil.TempDir("", "nondirwalk")
	must(t, err)
	defer os.RemoveAll(tmpPath)

	foobarPath := path.Join(tmpPath, "foobar")
	f, err := os.Create(foobarPath)
	must(t, err)
	must(t, f.Close())

	_, err = WalkDir(f.Name(), &WalkOpts{})
	assert.NotNil(t, err, "should refuse to walk non-directory")
}

func Test_WalkZip(t *testing.T) {
	tmpPath := mktestdir(t, "walkzip")
	defer os.RemoveAll(tmpPath)

	tmpPath2, err := ioutil.TempDir("", "walkzip2")
	must(t, err)
	defer os.RemoveAll(tmpPath2)

	container, err := WalkDir(tmpPath, &WalkOpts{})
	must(t, err)

	zipPath := path.Join(tmpPath2, "container.zip")
	zipWriter, err := os.Create(zipPath)
	must(t, err)
	defer zipWriter.Close()

	_, err = archiver.CompressZip(zipWriter, tmpPath, &state.Consumer{})
	must(t, err)

	zipSize, err := zipWriter.Seek(0, io.SeekCurrent)
	must(t, err)

	zipReader, err := zip.NewReader(zipWriter, zipSize)
	must(t, err)

	zipContainer, err := WalkZip(zipReader, &WalkOpts{})
	must(t, err)

	if testSymlinks {
		assert.Equal(t, "5 files, 3 dirs, 2 symlinks", container.Stats(), "should report correct stats")
	} else {
		assert.Equal(t, "5 files, 3 dirs, 0 symlinks", container.Stats(), "should report correct stats")
	}

	totalSize := int64(0)
	for _, regular := range regulars {
		totalSize += int64(regular.Size)
	}
	assert.Equal(t, totalSize, container.Size, "should report correct size")

	must(t, container.EnsureEqual(zipContainer))
}

func Test_Walk(t *testing.T) {
	tmpPath := mktestdir(t, "walk")
	defer os.RemoveAll(tmpPath)

	container, err := WalkDir(tmpPath, &WalkOpts{})
	must(t, err)

	dirs := []string{
		"foo",
		"foo/dir_a",
		"foo/dir_b",
	}
	for i, dir := range dirs {
		assert.Equal(t, dir, container.Dirs[i].Path, "dirs should be all listed")
	}

	files := []string{
		"foo/dir_a/baz",
		"foo/dir_a/bazzz",
		"foo/dir_b/zoom",
		"foo/file_f",
		"foo/file_z",
	}
	for i, file := range files {
		assert.Equal(t, file, container.Files[i].Path, "files should be all listed")
	}

	if testSymlinks {
		for i, symlink := range symlinks {
			assert.Equal(t, symlink.Newname, container.Symlinks[i].Path, "symlink should be at correct path")
			assert.Equal(t, symlink.Oldname, container.Symlinks[i].Dest, "symlink should point to correct path")
		}
	}

	if testSymlinks {
		assert.Equal(t, "5 files, 3 dirs, 2 symlinks", container.Stats(), "should report correct stats")
	} else {
		assert.Equal(t, "5 files, 3 dirs, 0 symlinks", container.Stats(), "should report correct stats")
	}

	totalSize := int64(0)
	for _, regular := range regulars {
		totalSize += int64(regular.Size)
	}
	assert.Equal(t, totalSize, container.Size, "should report correct size")

	if testSymlinks {
		container, err := WalkDir(tmpPath, &WalkOpts{Dereference: true})
		must(t, err)

		assert.EqualValues(t, 0, len(container.Symlinks), "when dereferencing, no symlinks should be listed")

		files := []string{
			"foo/dir_a/baz",
			"foo/dir_a/bazzz",
			"foo/dir_b/zoom",
			"foo/file_f",
			"foo/file_m",
			"foo/file_o",
			"foo/file_z",
		}
		for i, file := range files {
			assert.Equal(t, file, container.Files[i].Path, "when dereferencing, symlinks should appear as files")
		}

		// add both dereferenced symlinks to total size
		totalSize += int64(regulars[3].Size) // foo/file_z
		totalSize += int64(regulars[1].Size) // foo/dir_a/baz
		assert.Equal(t, totalSize, container.Size, "when dereferencing, should report correct size")
	}
}

func Test_Prepare(t *testing.T) {
	tmpPath := mktestdir(t, "prepare")
	defer os.RemoveAll(tmpPath)

	container, err := WalkDir(tmpPath, &WalkOpts{})
	must(t, err)

	tmpPath2, err := ioutil.TempDir("", "prepare")
	defer os.RemoveAll(tmpPath2)
	must(t, err)

	err = container.Prepare(tmpPath2)
	must(t, err)

	container2, err := WalkDir(tmpPath2, &WalkOpts{})
	must(t, err)

	must(t, container.EnsureEqual(container2))
}

// Support code

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

var testSymlinks = runtime.GOOS != "windows"

func mktestdir(t *testing.T, name string) string {
	tmpPath, err := ioutil.TempDir("", "tmp_"+name)
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

	if testSymlinks {
		for _, entry := range symlinks {
			new := filepath.Join(tmpPath, entry.Newname)
			must(t, os.Symlink(entry.Oldname, new))
		}
	}

	return tmpPath
}
