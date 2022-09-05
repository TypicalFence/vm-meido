package digitalocean

import (
	"context"
	"errors"

	"github.com/digitalocean/godo"
	. "zaun.moe/vm-meido"
	"zaun.moe/vm-meido/distro"
	vm_size "zaun.moe/vm-meido/vm-size"
)

func getDropletByName(ctx context.Context, client *godo.Client, name string) (*godo.Droplet, error) {
	// create options. initially, these will be blank
	opt := &godo.ListOptions{}
	droplets, _, err := client.Droplets.ListByName(ctx, name, opt)
	if err != nil {
		return nil, err
	}

	if len(droplets) > 0 {
		return &droplets[0], nil
	}

	return nil, nil
}

func createDroplet(ctx context.Context, client *godo.Client, name string, settings VmSettings) (*godo.Droplet, error) {
	slug, err := convertDistro(settings.Distro)

	if err != nil {
		return nil, err
	}

	region, err := convertRegion(settings.Region)

	if err != nil {
		return nil, err
	}

	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   convertSize(settings.Size),
		Image: godo.DropletCreateImage{
			Slug: slug,
		},
		SSHKeys: []godo.DropletCreateSSHKey{{Fingerprint: settings.SSHkey}},
		Tags:    []string{"vm-meido"},
	}

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		return nil, err
	}

	return newDroplet, nil
}

func destroyDroplet(ctx context.Context, client *godo.Client, id int) (bool, error) {
	resp, err := client.Droplets.Delete(ctx, id)

	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

func convertDistro(name distro.Distro) (string, error) {
	switch name {
	case distro.DEBIAN_11:
		return "debian-11-x64", nil

	case distro.UBUNTU_20_04:
		return "ubuntu-22-04-x64", nil

	default:
		println("distro: " + name)
		return "", errors.New("image does not exist at digitalocean")
	}
}

func convertSize(name vm_size.VmSize) string {
	switch name {
	case vm_size.XS:
		return "s-1vcpu-512mb-10gb"

	case vm_size.SM:
		return "s-1vcpu-1gb"

	case vm_size.LG:
		return "s-2vcpu-2gb-amd"

	default:
		return convertSize(vm_size.XS)
	}
}

func convertRegion(name Region) (string, error) {
	switch name {
	case REGION_EU_WEST:
		return "fra1", nil

	case REGION_UK:
		return "lon1", nil

	case REGION_US_WEST:
		return "sfo3", nil

	case REGION_US_EAST:
		return "nyc3", nil

	default:
		println("Region: " + name)
		return "", errors.New("unsupported region")
	}
}

func getStatusFromDroplet(droplet *godo.Droplet) VmStatus {
	if droplet != nil {
		if droplet.Status == "active" {
			return VM_STATUS_RUNNING
		} else {
			return VM_STAUS_IDLE
		}
	}

	return VM_STATUS_GONE
}
