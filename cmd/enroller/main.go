package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
	"uuid"

	"github.com/hjamieson/homeregistry/internal/model"
	"github.com/hjamieson/homeregistry/internal/platform"
)

func main() {
	fmt.Println("the enroller!")
	hostname, _ := os.Hostname()
	os := runtime.GOOS
	machineID, _ := platform.MachineID()
	manuf , _ := platform.Manufacturer()

	device := model.Device{
		Id : uuid.NewV7().String(),
		OS: os,
		Hostname: hostname,
		MachineID: machineID,
		Manuf: manuf,
		LastInfo: time.Now(),
	}
	out, _ := json.Marshal(device)

	fmt.Println(string(out))
}
