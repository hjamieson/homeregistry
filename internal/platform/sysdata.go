package platform

import (
	"fmt"
	"net"

	"github.com/hjamieson/homeregistry/internal/model"
)

func NicData() ([]model.Nic, error) {
	intfc, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var nics []model.Nic

	for _, nic := range intfc {
		if nic.Flags&net.FlagLoopback != 0 {
			continue
		}
		fmt.Printf("%#v ", nic)
		fmt.Printf("| MAC(%v)", nic.HardwareAddr.String())
		if nic.Flags&net.FlagUp != 0 {
			fmt.Printf("| UP")
		}
		addrs, _ := nic.Addrs()
		if len(addrs) == 0 {
			continue
		}
		device := model.Nic{}
		device.MAC = nic.HardwareAddr.String()
		device.IP = addrs[0].String()
		nics = append(nics, device)
		fmt.Println()
		fmt.Printf("NIC: %v\n", device)
	}

	return nics, nil
}
