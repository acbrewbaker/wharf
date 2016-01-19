package megafile_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/itchio/wharf.proto/megafile"
	"github.com/stretchr/testify/assert"
)

func Test_Write(t *testing.T) {
	tmpPath := mktestdir(t, "write_read")
	defer os.RemoveAll(tmpPath)

	info, err := megafile.Walk(tmpPath, 16)
	must(t, err)

	wmpPath, err := ioutil.TempDir(".", "tmp_write_write")
	must(t, err)
	defer os.RemoveAll(wmpPath)

	w, err := info.NewWriter(wmpPath)
	must(t, err)

	info2, err := megafile.Walk(wmpPath, 16)
	must(t, err)
	assert.Equal(t, info, info2, "Created same directory structure")

	must(t, w.Close())
}