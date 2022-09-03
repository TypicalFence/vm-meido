package vm_meido

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

type VmStatus string

const (
	VM_STATUS_RUNNING VmStatus = "RUNNING"
	VM_STAUS_IDLE     VmStatus = "IDLE"
	VM_STATUS_GONE    VmStatus = "GONE"
)

type VmState struct {
	Name      string   `json:"name"`
	Provider  string   `json:"provider"`
	IpAddress string   `json:"ip"`
	VmStatus  VmStatus `josn:"vm_status"`
}

func SaveState(state VmState) error {
	bytes, err := json.Marshal(state)
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	meidoDir := path.Join(cwd, ".vm-meido")

	err = os.MkdirAll(meidoDir, os.ModePerm)
	if err != nil {
		return err
	}

	stateFile := path.Join(meidoDir, "state.json")

	f, err := os.Create(stateFile)
	f.Write(bytes)

	return nil
}

func LoadState() (*VmState, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	meidoDir := path.Join(cwd, ".vm-meido")

	err = os.MkdirAll(meidoDir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	stateFile := path.Join(meidoDir, "state.json")

	bytes, err := ioutil.ReadFile(stateFile)
	if err != nil {
		return nil, err
	}

	var state VmState

	err = json.Unmarshal(bytes, &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}
