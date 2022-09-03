package vm_meido

import (
	"io/ioutil"
	"path"

	"git.sr.ht/~spc/go-ini"
	"github.com/mitchellh/go-homedir"
	"zaun.moe/vm-meido/providers"
)

type Config struct {
	Region      Region                 `ini:"region"`
	Provider    providers.ProviderName `ini:"provider"`
	Digtalocean struct {
		ApiKey string `ini:"api_key"`
	} `ini:"digitalocean"`
}

func LoadConfig() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(home, ".config/vm-meido.ini")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = ini.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
