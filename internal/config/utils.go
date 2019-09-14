package config

import (
	"io/ioutil"
	"os"
)

const (
	fullOwnerPermission os.FileMode = 0700
)

func safeCreateSymLink(oldName, newName string) error {
	if err := os.Remove(newName); !os.IsNotExist(err) {
		return err
	}

	return os.Symlink(oldName, newName)
}

func createSymLinkIfNotExists(oldName, newName string) (err error) {
	if _, err = os.Stat(newName); os.IsNotExist(err) {
		return os.Symlink(oldName, newName)
	}

	return err
}

func createDirIfNotExists(dirPath string) (err error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, fullOwnerPermission)
	}

	return err
}

func createFileIfNotExists(filePath string) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return createEmptyFile(filePath)
	}

	return err
}

func createEmptyFile(filePath string) error {
	return ioutil.WriteFile(filePath, nil, fullOwnerPermission)
}
