package utils

import (
	"io/ioutil"
	"os"
)

const (
	fullOwnerPermission os.FileMode = 0700
)

func SafeCreateSymLink(oldName, newName string) error {
	if err := os.Remove(newName); !os.IsNotExist(err) {
		return err
	}
	return os.Symlink(oldName, newName)
}

func CreateSymLinkIfNotExists(oldName, newName string) (err error) {
	if _, err = os.Stat(newName); os.IsNotExist(err) {
		return os.Symlink(oldName, newName)
	}
	return err
}

func CreateDirIfNotExists(dirPath string) (err error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, fullOwnerPermission)
	}
	return err
}

func CreateEmptyFileIfNotExists(filePath string) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		return CreateEmptyFile(filePath)
	}
	return err
}

func CreateEmptyFile(filePath string) error {
	return ioutil.WriteFile(filePath, nil, fullOwnerPermission)
}
