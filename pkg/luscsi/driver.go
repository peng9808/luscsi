package luscsi

import (
	csicommon "sigs.k8s.io/azurelustre-csi-driver/pkg/csi-common"
)

// DriverOptions defines driver parameters specified in driver deployment
type DriverOptions struct {
	NodeID                     string
	DriverName                 string
	EnableAzureLustreMockMount bool
	WorkingMountDir            string
}

// Driver implements all interfaces of CSI drivers
type Driver struct {
	csicommon.CSIDriver
	csicommon.DefaultIdentityServer
	csicommon.DefaultControllerServer
	csicommon.DefaultNodeServer
}

func NewDriver(options *DriverOptions) *Driver {
	d := Driver{}
	d.Name = options.DriverName
	d.Version = driverVersion
	d.NodeID = options.NodeID

	d.DefaultControllerServer.Driver = &d.CSIDriver
	d.DefaultIdentityServer.Driver = &d.CSIDriver
	d.DefaultNodeServer.Driver = &d.CSIDriver

	return &d
}
