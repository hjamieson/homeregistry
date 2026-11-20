package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"uuid"

	"github.com/hjamieson/homeregistry/internal/model"
	"github.com/hjamieson/homeregistry/internal/platform"
)

func main() {
	hostname, _ := os.Hostname()
	os := platform.Os()
	machineID, _ := platform.MachineID()
	manuf, _ := platform.Manufacturer()
	nic, _ := platform.NetworkInterface()

	device := model.Device{
		Id:        uuid.NewV7().String(),
		OS:        os,
		NIC:       []model.Nic{nic},
		Hostname:  hostname,
		MachineID: machineID,
		Manuf:     manuf,
		LastInfo:  time.Now(),
	}
	out, _ := json.Marshal(device)

	fmt.Println(string(out))
}
