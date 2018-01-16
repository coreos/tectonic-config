package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// ToFile writes the configuration object to a file at path. Supported extensions are `json` and
// `yaml`.
func ToFile(path string, obj interface{}) error {
	var data []byte
	var err error
	switch filepath.Ext(path) {
	case ".yaml":
		data, err = yaml.Marshal(obj)
	case ".json":
		data, err = json.Marshal(obj)
	default:
		err = fmt.Errorf("unrecognized config file format: %q", filepath.Ext(path))
	}
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

// FromFile reads the configuration object from the file at path. Supported extensions are `json` and
// `yaml`.
func FromFile(path string, obj interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config from %q: %v", path, err)
	}

	switch filepath.Ext(path) {
	case ".yaml":
		return yaml.Unmarshal(b, obj)
	case ".json":
		return json.Unmarshal(b, obj)
	default:
		return fmt.Errorf("unrecognized config file format: %q", filepath.Ext(path))
	}
}
