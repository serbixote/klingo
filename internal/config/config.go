package config

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	defaultKlingoDir    string      = "~/.klingo"
	FullOwnerPermission os.FileMode = 0700
)

var (
	config *klingoConfig
)

type klingoConfig struct {
	dir            string
	contexts       []string
	currentContext string
}

func (*klingoConfig) UseContext(context string) error {
	return nil
}

func GetKlingoConfig() (*klingoConfig, error) {
	if config != nil {
		return config, nil
	}

	klingoDir := defaultKlingoDir

	if KLINGO_HOME := os.Getenv("KLINGO_HOME"); strings.TrimSpace(KLINGO_HOME) != "" {
		klingoDir = KLINGO_HOME
	}

	err := initConfigFiles(klingoDir)
	if err != nil {
		return nil, err
	}

	// Load config to config var
	err = loadConfig(klingoDir)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func initConfigFiles(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, FullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	contextsDir := path.Join(dir, "contexts")
	if _, err := os.Stat(contextsDir); os.IsNotExist(err) {
		if err := os.Mkdir(contextsDir, FullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	defaultContext := path.Join(contextsDir, "default")
	configFileSymLink := path.Join(contextsDir, "config")

	if _, err := os.Stat(defaultContext); os.IsNotExist(err) {
		comment := []byte("# config file used by default")
		err := ioutil.WriteFile(defaultContext, comment, FullOwnerPermission)
		if err != nil {
			return err
		}

		if _, err = os.Stat(configFileSymLink); os.IsExist(err) {
			if err = os.Remove(configFileSymLink); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if _, err := os.Stat(configFileSymLink); os.IsNotExist(err) {
		if err = os.Symlink(defaultContext, configFileSymLink); err != nil {
			return err
		}
	}

	return nil
}

func loadConfig(dir string) error {
	config := &klingoConfig{
		dir: dir,
	}

	contextsDir := path.Join(dir, "contexts")
	files, err := ioutil.ReadDir(contextsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			config.contexts = append(
				config.contexts,
				file.Name(),
			)
		}
	}

	destination, err := os.Readlink(path.Join(dir, "config"))
	if err != nil {
		return err
	}

	config.currentContext = destination

	return nil
}
