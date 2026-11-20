package platform

import (
	"fmt"
	"net"
)

func NicData() ([]string, error) {
	nics, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, nic := range nics {
		fmt.Printf("%#v ", nic)
		fmt.Printf("| MAC(%v)", nic.HardwareAddr.String())
		if nic.Flags&net.FlagLoopback != 0 {
			fmt.Printf("| loopback")
		}
		if nic.Flags&net.FlagUp != 0 {
			fmt.Printf("| UP")
		}
		addrs, _ := nic.Addrs()
		if len(addrs) != 0 {
			fmt.Printf("|addrs: %v", addrs)
		}
		fmt.Println()
	}



	return nil, nil
}
