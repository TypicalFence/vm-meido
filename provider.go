package vm_meido

import (
	"zaun.moe/vm-meido/distro"
	vm_size "zaun.moe/vm-meido/vm-size"
)

type Provider interface {
	GetVmStatus(vmName string) (VmStatus, error)
	GetVmState(vmName string) (*VmState, error)
	ProvisionVm(settings VmSettings) (string, error)
	DestroyVm(vmName string) (bool, error)
}

type VmSettings struct {
	Size   vm_size.VmSize
	Distro distro.Distro
	Region Region
	SSHkey string
}
