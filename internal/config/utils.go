package config

import (
	"io/ioutil"
	"os"
)

const (
	fullOwnerPermission os.FileMode = 0700
)

// safeCreateSymLink creates a symlink after deleting the
// newname if already exists.
func safeCreateSymLink(oldName, newName string) error {
	if err := os.Remove(newName); !os.IsNotExist(err) {
		return err
	}

	return os.Symlink(oldName, newName)
}

// createSymLinkIfNotExists creates a symlink if newname
// does not already exists.
func createSymLinkIfNotExists(oldName, newName string) (err error) {
	if _, err = os.Stat(newName); os.IsNotExist(err) {
		return os.Symlink(oldName, newName)
	}

	return err
}

// createDirIfNotExists creates a directory if this does
// not exists yet.
func createDirIfNotExists(dirPath string) (err error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, fullOwnerPermission)
	}

	return err
}

// createFileIfNotExists is creates a file if this does
// not exists yet.
func createFileIfNotExists(filePath string) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return createEmptyFile(filePath)
	}

	return err
}

// createEmptyFile creates an empty file with 0700 permissions.
func createEmptyFile(filePath string) error {
	return ioutil.WriteFile(filePath, nil, fullOwnerPermission)
}
