package configloader

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

func Load(path string, cfg interface{}) error {
	err := k.Load(file.Provider(path), yaml.Parser())
	if err != nil {
		return err
	}

	err = k.Unmarshal("", &cfg)
	if err != nil {
		return err
	}

	return nil
}
