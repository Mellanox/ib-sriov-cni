// Copyright 2025 ib-sriov-cni authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	"net"

	"github.com/containernetworking/cni/pkg/types"
	"github.com/containernetworking/plugins/pkg/ns"
	"github.com/vishvananda/netlink"

	rdmatypes "github.com/k8snetworkplumbingwg/rdma-cni/pkg/types"
)

// NetConf extends types.NetConf for ib-sriov-cni
type NetConf struct {
	types.NetConf
	Master              string
	DeviceID            string `json:"deviceID"` // PCI address of a VF in valid sysfs format
	VFID                int
	HostIFNames         string // VF netdevice name(s)
	HostIFGUID          string // VF netdevice GUID
	ContIFNames         string // VF names after in the container; used during deletion
	GUID                string `json:"-"` // Taken from either CNI_ARGS "guid" attribute or from RuntimeConfig
	PKey                string `json:"pkey"`
	LinkState           string `json:"link_state,omitempty"` // auto|enable|disable
	RdmaIso             bool   `json:"rdmaIsolation,omitempty"`
	IBKubernetesEnabled bool   `json:"ibKubernetesEnabled,omitempty"`
	RdmaNetState        rdmatypes.RdmaNetState
	RuntimeConfig       RuntimeConf `json:"runtimeConfig,omitempty"`
	Args                struct {
		CNI map[string]string `json:"cni"`
	} `json:"args"`
}

// RuntimeConf represents the plugin's runtime configurations
type RuntimeConf struct {
	InfinibandGUID string `json:"infinibandGUID"`
}

// Manager provides interface invoke sriov nic related operations
type Manager interface {
	SetupVF(conf *NetConf, podifName string, cid string, netns ns.NetNS) error
	ReleaseVF(conf *NetConf, podifName string, cid string, netns ns.NetNS) error
	ResetVFConfig(conf *NetConf) error
	ApplyVFConfig(conf *NetConf) error
}

// mocked netlink interface
// required for unit tests

// NetlinkManager is an interface to mock nelink library
type NetlinkManager interface {
	LinkByName(string) (netlink.Link, error)
	LinkSetUp(netlink.Link) error
	LinkSetDown(netlink.Link) error
	LinkSetNsFd(netlink.Link, int) error
	LinkSetName(netlink.Link, string) error
	LinkSetVfState(netlink.Link, int, uint32) error
	LinkSetVfPortGUID(netlink.Link, int, net.HardwareAddr) error
	LinkSetVfNodeGUID(netlink.Link, int, net.HardwareAddr) error
	LinkDelAltName(netlink.Link, string) error
}

// PciUtils is interface to help in SR-IOV functions
type PciUtils interface {
	GetSriovNumVfs(ifName string) (int, error)
	GetVFLinkNamesFromVFID(pfName string, vfID int) ([]string, error)
	GetPciAddress(ifName string, vf int) (string, error)
	RebindVf(pfName, vfPciAddress string) error
}
