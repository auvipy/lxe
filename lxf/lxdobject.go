package lxf

import "github.com/lxc/lxe/lxf/device"

// LXDObject contains common properties of containers and sandboxes without CRI influence
type LXDObject struct {
	// client holds the lxf.Client representing as a lxd client
	client *Client
	// ID is a unique generated ID and is read-only
	ID string
	// ETag uniquely identifies user modifiable content of this resource, prevents race conditions when saving
	// see: https://lxd.readthedocs.io/en/latest/api-extensions/#etag
	ETag string
	// ETag uniquely identifies user modifiable content of the state of this resource, prevents race conditions when
	// trying to modify the state
	//StateETag string
	// Devices
	Proxies device.Proxies
	Disks   device.Disks
	Blocks  device.Blocks
	Nics    device.Nics
	// Config contains options not provided by a own property
	Config map[string]string
}
