package platform

import (
	"errors"
	"log"
	"os"
	"strings"
)

var PLATFORM_ERR = errors.New("platform error")
var MACHINEID = "/etc/machine-id"

func MachineID() (string, error) {
	bytes, err := os.ReadFile(MACHINEID)
	if err != nil {
		log.Printf("unable to open file: %v", err)
		return "", PLATFORM_ERR
	}
	ID := strings.TrimSuffix(string(bytes), "\n")

	return ID, nil
}
