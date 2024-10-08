// Getting network interfaces for Linux

//go:build linux

package network

import (
	"github.com/cloud-barista/cm-honeybee/agent/lib/routes"
	network2 "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/network"
	"github.com/shirou/gopsutil/v3/net"
)

func GetNICs() ([]network2.NIC, error) {
	var nics []network2.NIC

	interfaces, err := net.Interfaces()
	if err != nil {
		return nics, err
	}

	var defaultRoutes []routes.RouteStruct

	defaultRoutes, err = routes.GetRoutes(true)
	if err != nil {
		return nics, err
	}

	for _, i := range interfaces {
		var addresses []string
		var gateway string

		for _, a := range i.Addrs {
			addresses = append(addresses, a.Addr)
		}

		for _, route := range defaultRoutes {
			if route.Interface == i.Name {
				gateway = route.NextHop
				break
			}
		}

		nics = append(nics, network2.NIC{
			Interface:  i.Name,
			Address:    addresses,
			Gateway:    gateway,
			MACAddress: i.HardwareAddr,
			MTU:        i.MTU,
		})
	}

	return nics, nil
}
