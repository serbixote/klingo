package config

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

const (
	defaultKlingoDir       string      = ".klingo"
	defaultContextsDir     string      = "contexts"
	DefaultContextFileName string      = "default"
	defaultConfigFileName  string      = "config"
	fullOwnerPermission    os.FileMode = 0700
)

var (
	config *klingoConfig
)

func GetKlingoConfig() *klingoConfig {
	return config
}

type klingoConfig struct {
	dirPath        string
	contexts       []string
	currentContext string
}

func (c *klingoConfig) CurrentContext() string {
	return c.currentContext
}

func (c *klingoConfig) UseContext(context string) error {
	for _, item := range c.contexts {
		if item == context {
			contextsDirPath := path.Join(c.dirPath, defaultContextsDir)
			contextFilePath := path.Join(contextsDirPath, item)
			configFileSymLink := path.Join(c.dirPath, defaultConfigFileName)
			if err := os.Remove(configFileSymLink); err != nil {
				return err
			}
			if err := os.Symlink(contextFilePath, configFileSymLink); err != nil {
				return err
			}

			c.currentContext = item

			fmt.Printf("Switched to context %s\n", c.currentContext)

			return nil
		}
	}

	return errors.Errorf("no context exists with the name: %s", context)
}

func init() {
	klingoDir := defaultKlingoDir

	if KLINGO_DIR := os.Getenv("KLINGO_DIR"); strings.TrimSpace(KLINGO_DIR) != "" {
		klingoDir = KLINGO_DIR
	} else {
		klingoDir = path.Join(os.Getenv("HOME"), klingoDir)
	}

	err := initConfigFileStructure(klingoDir)
	if err != nil {
		fmt.Printf("error: %v", errors.Wrap(
			err, "Failed initializing file structure",
		))
		os.Exit(1)
	}

	err = loadKlingoConfig(klingoDir)
	if err != nil {
		fmt.Printf("error: %v", errors.Wrap(
			err, "Failed loading configuration files",
		))
		os.Exit(1)
	}
}

// initConfigFileStructure creates the config file structure in the
// given dir param if needed:
//
// .klingo/
// |__ config (Symbolic link to any contexts/*)
// |__ contexts/
// |   |__ default
func initConfigFileStructure(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, fullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} // TODO: Check permissions of existing directory and warn user if needed

	contextsDir := path.Join(dir, defaultContextsDir)
	if _, err := os.Stat(contextsDir); os.IsNotExist(err) {
		if err := os.Mkdir(contextsDir, fullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	defaultContext := path.Join(contextsDir, DefaultContextFileName)
	configFileSymLink := path.Join(dir, defaultConfigFileName)

	if _, err := os.Stat(defaultContext); os.IsNotExist(err) {
		comment := []byte("# default config file\n")
		err := ioutil.WriteFile(defaultContext, comment, fullOwnerPermission)
		if err != nil {
			return err
		}

		if _, err = os.Stat(configFileSymLink); os.IsExist(err) {
			if err = os.Remove(configFileSymLink); err != nil {
				return err
			}
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

// loadKlingoConfig loads the configuration from the give dir param
// to the config variable (implicitly). Expected directory structure:
//
// .klingo/
// |__ config (Symbolic link to any contexts/*)
// |__ contexts/
// |   |__ default
// |   |__ home
// |   |__ project1
// |   |__ ...
func loadKlingoConfig(dir string) error {
	config = &klingoConfig{
		dirPath: dir,
	}

	contextsDirPath := path.Join(dir, defaultContextsDir)
	files, err := ioutil.ReadDir(contextsDirPath)
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
	sort.Strings(config.contexts)

	currentContextFileName, err := os.Readlink(path.Join(dir, defaultConfigFileName))
	if err != nil {
		return err
	}

	config.currentContext = path.Base(currentContextFileName)
	return nil
}
