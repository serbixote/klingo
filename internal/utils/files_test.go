package utils

import (
	"gotest.tools/assert"
	"os"
	"path"
	"testing"
)

func workDirAbsolutePath(t *testing.T) string {
	t.Helper()

	wd, err := os.Getwd()
	assert.NilError(t, err)
	return wd
}

func TestCreateEmptyFile(t *testing.T) {
	wd := workDirAbsolutePath(t)
	emptyFileAbsolutePath := path.Join(wd, "testdata", "empty-file")

	err := CreateEmptyFile(emptyFileAbsolutePath)
	assert.NilError(t, err)

	if fileInfo, err := os.Stat(emptyFileAbsolutePath); os.IsNotExist(err) {
		t.Errorf("%s file wasn't created", emptyFileAbsolutePath)
	} else if fileInfo.Size() > 0 {
		t.Errorf("%s file isn't empty", emptyFileAbsolutePath)
	}

	assert.NilError(t, err)
}
