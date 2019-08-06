package config

import (
	"errors"
	"regexp"
	"strings"
)

// section is struct for representing a section in a ssh config file.
type section struct {
	Host string // Argument of the Host keyword
}

// Section exported struct for section.
type Section struct {
	section
}

// NewSection returns a Section instance created using the given name.
// The name string must:
// - Contain only alphanumeric characters and optional '/' characters
//   for representing its "parent folders" at the UI.
// - Have an alphanumeric character as the last one, not '/'
// - Have more than 2 characters
// TODO: add maximum
func (h *Section) NewSection(name string) (*Section, error) {
	err := validName(name)
	if err != nil {
		return nil, err
	}

	name = strings.ReplaceAll(name, "/", ".")

	return &Section{
		section: section{
			Host: name,
		},
	}, nil
}

// Name returns the full Name as it should to be shown at the UI.
func (h *Section) Name() string {
	return strings.ReplaceAll(h.Host, ".", "/")
}

// validName checks that the given string is a "valid" name, according
// to the requirements listed at the NewSection method description.
func validName(name string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9/]*$")
	if !re.MatchString(name) {
		return errors.New("Name can not contain other than alphanumeric and '/' characters")
	}

	return nil
}
