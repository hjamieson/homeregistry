package platform

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/hjamieson/homeregistry/internal/model"
)

var ErrNoData = errors.New("not found")

const notavailable = ""
const idfile = "/etc/machine-id"

var uuidPattern = regexp.MustCompile("^[0-9a-z]+$")

func MachineID() (string, error) {
	data, err := os.ReadFile(idfile)
	if err != nil {
		return notavailable, fmt.Errorf("not found: %w", err)
	}
	id := strings.TrimSpace(string(data))
	if len(id) != 32 {
		return notavailable, ErrNoData
	}
	if !uuidPattern.MatchString(id) {
		return notavailable, ErrNoData
	}

	return id, nil
}

func HostnameCtl() (map[string]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	cmd := exec.CommandContext(ctx, "hostnamectl")
	defer cancel()

	props := make(map[string]string)

	data, err := cmd.Output()
	if err != nil {
		return props, err
	}

	for line := range strings.Lines(string(data)) {
		if t := strings.Split(line, ":"); len(t) == 2 {
			key := strings.TrimSpace(t[0])
			value := strings.TrimSpace(t[1])
			props[key] = value
		}
	}

	return props, nil
}

func Manufacturer() (string, error) {
	props, err := HostnameCtl()
	if err != nil {
		return "", err
	}

	manuf, ok := props["Hardware Vendor"]
	if !ok {
		return "", errors.New("not found")
	}

	return manuf, nil
}

func Os() string {
	goos := runtime.GOOS
	props, err := HostnameCtl()
	if err != nil {
		return goos
	}

	os, ok := props["Operating System"]
	if !ok {
		return goos
	}
	return os
}

func NetworkInterface() (model.Nic, error) {
	return model.Nic{}, errors.New("no impl")
}
