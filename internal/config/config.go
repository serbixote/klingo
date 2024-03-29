package config

import (
	"fmt"
	"github.com/marco2704/klingo/internal/utils"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"
)

const (
	defaultKlingoDir   string = ".klingo"
	defaultContextsDir string = "contexts"
	defaultContextFile string = "default"
	defaultConfigFile  string = "config"
)

var (
	config *klingoConfig
)

type klingoConfig struct {
	dirPath        string
	contexts       []string
	currentContext string
}

func GetKlingoConfig() *klingoConfig {
	return config
}

func (c *klingoConfig) CreateContext(context string) error {
	if c.contextExists(context) {
		return errors.Errorf("context %s already exists", context)
	}

	contextFilePath := c.contextFilePath(context)
	if err := utils.CreateEmptyFile(contextFilePath); err != nil {
		return errors.Wrap(err, "failed creating context")
	}

	fmt.Printf("Created context %s\n", context)

	return nil
}

func (c *klingoConfig) UseContext(context string) error {
	if context == "" {
		context = defaultContextFile
	}

	if c.contextExists(context) {
		contextFilePath := c.contextFilePath(context)
		configFilePath := c.configFilePath()

		if err := utils.SafeCreateSymLink(contextFilePath, configFilePath); err != nil {
			return errors.Wrap(err, "failed switching context")
		}

		c.currentContext = context

		fmt.Printf("Switched to context %s\n", c.currentContext)

		return nil
	}

	return errors.Errorf("context %s does not exist", context)
}

func (c *klingoConfig) DeleteContext(context string) error {
	if context == defaultContextFile {
		return errors.Errorf("context %s can't be deleted", defaultContextFile)
	}

	if !c.contextExists(context) {
		return errors.Errorf("context %s doesn't exist", context)
	}

	contextFilePath := c.contextFilePath(context)
	if err := os.Remove(contextFilePath); err != nil {
		return errors.Wrap(err, "failed deleting context")
	}

	fmt.Printf("Deleted context %s\n", context)

	return nil
}

func (c *klingoConfig) RenameContext(oldContext, newContext string) error {
	if c.contextExists(oldContext) {
		return errors.Errorf("context %s doesn't exist", oldContext)
	} else if c.contextExists(newContext) {
		return errors.Errorf("context %s already exists", newContext)
	}

	oldContextPath := c.contextFilePath(oldContext)
	newContextPath := c.contextFilePath(newContext)

	if err := os.Rename(oldContextPath, newContextPath); err != nil {
		return errors.Wrap(err, "failed renaming context")
	}

	fmt.Printf("Renamed context %s to %s\n", oldContext, newContext)

	return nil
}

func (c *klingoConfig) CurrentContext() string {
	return c.currentContext
}

func (c *klingoConfig) Contexts() []string {
	return c.contexts
}

func (c *klingoConfig) contextExists(context string) bool {
	for _, i := range c.contexts {
		if i == context {
			return true
		}
	}
	return false
}

func (c *klingoConfig) configFilePath() string {
	return path.Join(c.dirPath, defaultConfigFile)
}

func (c *klingoConfig) contextsDirPath() string {
	return path.Join(c.dirPath, defaultContextsDir)
}

func (c *klingoConfig) contextFilePath(context string) string {
	return path.Join(c.dirPath, defaultContextsDir, context)
}

func init() {
	klingoDir := defaultKlingoDir

	if KLINGO_DIR := os.Getenv("KLINGO_DIR"); strings.TrimSpace(KLINGO_DIR) != "" {
		klingoDir = KLINGO_DIR
	} else {
		klingoDir = path.Join(os.Getenv("HOME"), klingoDir)
	}

	config = &klingoConfig{
		dirPath: klingoDir,
	}

	if err := config.initConfigFileStructure(); err != nil {
		panic(errors.Wrap(err, "failed initializing file structure"))
	}

	if err := config.loadKlingoConfig(); err != nil {
		panic(errors.Wrap(err, "failed loading configuration files"))
	}
}

// initConfigFileStructure creates the config file structure in the
// in dirPath (if needed):
//
// .klingo/
// |__ config (Symbolic link to any contexts/*)
// |__ contexts/
// |   |__ default
func (c *klingoConfig) initConfigFileStructure() error {
	if err := utils.CreateDirIfNotExists(c.dirPath); err != nil {
		return err
	}

	contextsDirPath := c.contextsDirPath()
	if err := utils.CreateDirIfNotExists(contextsDirPath); err != nil {
		return err
	}

	contextFilePath := c.contextFilePath(defaultContextFile)
	if err := utils.CreateEmptyFileIfNotExists(contextFilePath); err != nil {
		return err
	}

	configFilePath := c.configFilePath()
	if err := utils.CreateSymLinkIfNotExists(contextFilePath, configFilePath); err != nil {
		return err
	}

	return nil
}

// loadKlingoConfig loads the configuration from file structure
// at dirPath. Expected directory structure:
//
// .klingo/
// |__ config (Symbolic link to any contexts/*)
// |__ contexts/
// |   |__ default
// |   |__ home
// |   |__ project1
// |   |__ ...
func (c *klingoConfig) loadKlingoConfig() error {
	contextsDirPath := c.contextsDirPath()
	files, err := ioutil.ReadDir(contextsDirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			c.contexts = append(
				c.contexts,
				file.Name(),
			)
		}
	}

	sort.Strings(c.contexts)

	currentContextPath, err := os.Readlink(c.configFilePath())
	if err != nil {
		return err
	}

	c.currentContext = path.Base(currentContextPath)
	return nil
}
