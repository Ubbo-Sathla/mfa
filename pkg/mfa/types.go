package mfa

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/user"
	"path"
)

var (
	conf []*Mfa
)

type Mfa struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

func GetConfig() []*Mfa {
	return conf
}

func LoadConfig() error {
	b, err := LoadConfigBytes(".mfa", ".mfa.yml", ".mfa.yaml")
	if err != nil {
		return err
	}
	var c []*Mfa
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return err
	}

	conf = c

	return nil
}
func LoadConfigBytes(names ...string) ([]byte, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	// homedir
	for i := range names {
		content, err := ioutil.ReadFile(path.Join(u.HomeDir, names[i]))
		if err == nil {
			return content, nil
		}
	}
	// relative
	for i := range names {
		content, err := ioutil.ReadFile(names[i])
		if err == nil {
			return content, nil
		}
	}
	return nil, err
}
