package digitalocean

import (
	"context"

	"github.com/cip8/autoname"
	"github.com/digitalocean/godo"
	. "zaun.moe/vm-meido"
	"zaun.moe/vm-meido/providers"
)

type DigitaloceanProvider struct {
	client godo.Client
}

func CreateDigitaloceanProvider(config Config) *DigitaloceanProvider {
	p := new(DigitaloceanProvider)
	// TODO validate
	apikey := config.Digtalocean.ApiKey
	p.client = *godo.NewFromToken(apikey)
	return p
}

func (p *DigitaloceanProvider) GetVmStatus(vmName string) (VmStatus, error) {
	ctx := context.TODO()
	droplet, err := getDropletByName(ctx, &p.client, vmName)

	if err != nil {
		return "", err

	}
	return getStatusFromDroplet(droplet), nil

}

func (p *DigitaloceanProvider) ProvisionVm(settings VmSettings) (string, error) {
	ctx := context.TODO()

	name := autoname.Generate("-")
	droplet, err := createDroplet(ctx, &p.client, name, settings)

	if err != nil {
		return "", err
	}

	return droplet.Name, nil
}

func (p *DigitaloceanProvider) GetVmState(vmName string) (*VmState, error) {
	ctx := context.TODO()

	droplet, err := getDropletByName(ctx, &p.client, vmName)

	if err != nil {
		return nil, err
	}

	ip, err := droplet.PublicIPv4()

	if err != nil {
		return nil, err
	}

	return &VmState{
		Name:      vmName,
		Provider:  string(providers.DIGITALOCEAN),
		IpAddress: ip,
		VmStatus:  getStatusFromDroplet(droplet),
	}, nil
}

func (p *DigitaloceanProvider) DestroyVm(vmName string) (bool, error) {
	ctx := context.TODO()

	droplet, err := getDropletByName(ctx, &p.client, vmName)

	if err != nil {
		return false, err
	}

	success, err := destroyDroplet(ctx, &p.client, droplet.ID)

	if err != nil {
		return false, err
	}

	return success, nil
}
