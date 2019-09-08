package config

import (
	"io/ioutil"
	"os"
)

const (
	fullOwnerPermission os.FileMode = 0700
)

func safeCreateSymLink(oldname, newname string) (err error) {
	if err := os.Remove(newname); !os.IsNotExist(err) {
		return err
	}

	return os.Symlink(oldname, newname)
}

func createSymLinkIfNotExists(oldname, newname string) (err error) {
	if _, err = os.Stat(newname); os.IsNotExist(err) {
		return os.Symlink(oldname, newname)
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
