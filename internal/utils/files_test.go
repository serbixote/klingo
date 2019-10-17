package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func workDirAbsolutePath(t *testing.T) string {
	t.Helper()

	wd, err := os.Getwd()
	assert.Nil(t, err)
	return wd
}

func removeFile(t *testing.T, fileAbsolutePath string) {
	err := os.Remove(fileAbsolutePath)
	assert.Nil(t, err)
}

func TestCreateEmptyFile(t *testing.T) {
	wd := workDirAbsolutePath(t)
	newEmptyFileAbsolutePath := path.Join(wd, "testdata", "new-empty-file")

	err := CreateEmptyFile(newEmptyFileAbsolutePath)
	assert.Nil(t, err)

	if fileInfo, err := os.Stat(newEmptyFileAbsolutePath); os.IsNotExist(err) {
		t.Errorf("%s file wasn't created", newEmptyFileAbsolutePath)
	} else if err != nil {
		t.Errorf("Expected nil, but got: %#v", err)
	} else if fileInfo.Size() > 0 {
		t.Errorf("%s file isn't empty", newEmptyFileAbsolutePath)
	}

	removeFile(t, newEmptyFileAbsolutePath)
}
