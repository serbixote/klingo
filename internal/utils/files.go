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

func CreateDirIfNotExists(absoluteDirPath string) (err error) {
	if _, err := os.Stat(absoluteDirPath); os.IsNotExist(err) {
		return os.Mkdir(absoluteDirPath, fullOwnerPermission)
	}
	return err
}

func CreateFileIfNotExists(absoluteFilePath string) (err error) {
	if _, err = os.Stat(absoluteFilePath); os.IsNotExist(err) {
		return CreateEmptyFile(absoluteFilePath)
	}
	return err
}

func CreateEmptyFile(absoluteFilePath string) error {
	return ioutil.WriteFile(absoluteFilePath, nil, fullOwnerPermission)
}
