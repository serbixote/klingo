package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

func workDir(t *testing.T) string {
	t.Helper()

	workDirPath, err := os.Getwd()
	assert.Nil(t, err)
	return workDirPath
}

func testDataDir(t *testing.T) string {
	workDirPath := workDir(t)
	return path.Join(workDirPath, "testdata")
}

func composeTestDataFilePath(t *testing.T, fileName string) string {
	testDataDirPath := testDataDir(t)
	return path.Join(testDataDirPath, fileName)
}

func removeFileOrEmptyDir(t *testing.T, filePath string) {
	err := os.Remove(filePath)
	assert.Nil(t, err)
}

func TestCreateDirIfNotExists(t *testing.T) {
	existingDirPath := composeTestDataFilePath(t, "existing-dir")
	err := CreateDirIfNotExists(existingDirPath)
	assert.Nil(t, err)

	newDirPath := composeTestDataFilePath(t, "new-dir")
	err = CreateDirIfNotExists(newDirPath)
	assert.Nil(t, err)

	if _, err = os.Stat(newDirPath); os.IsNotExist(err) {
		t.Errorf("%s dir wasn't created", newDirPath)
	} else {
		assert.Nil(t, err)
	}

	removeFileOrEmptyDir(t, newDirPath)
}

func TestCreateEmptyFileIfNotExist(t *testing.T) {
	emptyFilePath := composeTestDataFilePath(t, "empty-file")
	err := CreateEmptyFileIfNotExists(emptyFilePath)
	assert.Nil(t, err)

	newEmptyFilePath := composeTestDataFilePath(t, "new-empty-file")
	err = CreateEmptyFileIfNotExists(newEmptyFilePath)
	assert.Nil(t, err)

	if fileInfo, err := os.Stat(newEmptyFilePath); os.IsNotExist(err) {
		t.Errorf("%s file wasn't created", newEmptyFilePath)
	} else if err != nil {
		t.Errorf("Expected nil, but got: %#v", err)
	} else if fileInfo.Size() > 0 {
		t.Errorf("%s file isn't empty", newEmptyFilePath)
	}

	removeFileOrEmptyDir(t, newEmptyFilePath)
}

func TestCreateEmptyFile(t *testing.T) {
	newEmptyFilePath := composeTestDataFilePath(t, "new-empty-file")
	err := CreateEmptyFile(newEmptyFilePath)
	assert.Nil(t, err)

	if fileInfo, err := os.Stat(newEmptyFilePath); os.IsNotExist(err) {
		t.Errorf("%s file wasn't created", newEmptyFilePath)
	} else if err != nil {
		t.Errorf("Expected nil, but got: %#v", err)
	} else if fileInfo.Size() > 0 {
		t.Errorf("%s file isn't empty", newEmptyFilePath)
	}

	removeFileOrEmptyDir(t, newEmptyFilePath)
}
