package vm_meido

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"git.sr.ht/~spc/go-ini"
	distro "zaun.moe/vm-meido/distro"
	vm_size "zaun.moe/vm-meido/vm-size"
)

type MeidoFile struct {
	Distro distro.Distro  `ini:"distro"`
	Size   vm_size.VmSize `ini:"size"`
}

func LoadMeidoFile() (*MeidoFile, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configPath := path.Join(cwd, "MEIDOFILE")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, errors.New("No MEIDOFILE in current direcotry")
	}

	var meidoFile MeidoFile
	err = ini.UnmarshalWithOptions(data, &meidoFile, ini.Options{AllowNumberSignComments: true, AllowEmptyValues: false})
	if err != nil {
		return nil, err
	}

	return &meidoFile, nil
}
