package config

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	defaultKlingoDir    string      = ".klingo"
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

func (c *klingoConfig) UseContext(context string) error {
	for _, item := range c.contexts {
		if item == context {
			contextsDir := path.Join(c.dir, "contexts")
			contextPath := path.Join(contextsDir, item)
			configFileSymLink := path.Join(c.dir, "config")
			if err := os.Remove(configFileSymLink); err != nil {
				return err
			}
			if err := os.Symlink(contextPath, configFileSymLink); err != nil {
				return err
			}

			c.currentContext = item

			fmt.Printf("Switched to context %s\n", c.currentContext)

			return nil
		}
	}

	return errors.Errorf("no context exists with the name: %s", context)
}

func GetKlingoConfig() *klingoConfig {
	return config
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
		exit(errors.Wrap(
			err, "Failed initializing file structure",
		))
	}

	err = loadKlingoConfig(klingoDir)
	if err != nil {
		exit(errors.Wrap(
			err, "Failed loading configuration files",
		))
	}
}

// initConfigFileStructure creates the config file structure in the
// given dir param if needed:
// .klingo/
// |__ config (Symbolic link to any contexts/*)
// |__ contexts/
// |   |__ default
func initConfigFileStructure(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, FullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	} // TODO: Check permissions of existing directory and warn user if needed

	contextsDir := path.Join(dir, "contexts")
	if _, err := os.Stat(contextsDir); os.IsNotExist(err) {
		if err := os.Mkdir(contextsDir, FullOwnerPermission); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	defaultContext := path.Join(contextsDir, "default")
	configFileSymLink := path.Join(dir, "config")

	if _, err := os.Stat(defaultContext); os.IsNotExist(err) {
		comment := []byte("# config file used by default\n")
		err := ioutil.WriteFile(defaultContext, comment, FullOwnerPermission)
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

	config.currentContext = path.Base(destination)
	return nil
}

// exit prints the error after exiting with 1 as exit code.
func exit(e error) {
	fmt.Println(e)
	os.Exit(1)
}
