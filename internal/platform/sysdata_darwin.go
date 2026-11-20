package platform

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/hjamieson/homeregistry/internal/model"
)

const ioregCommand = "/usr/sbin/ioreg"

// example: "      \"IOPlatformUUID\" = \"8DE9B81F-54F4-5479-8029-1571CE7B0265\"\n"
var uuidPattern = regexp.MustCompile(`"IOPlatformUUID" = "([0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12})"`)

func MachineID() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, ioregCommand, "-rd1", "-c", "IOPlatformExpertDevice")
	out, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("ioreg failed: %w", err)
	}

	found := uuidPattern.FindStringSubmatch(string(out))
	if found == nil {
		return "", errors.New("not found")
	}

	return found[1], nil
}

func Manufacturer() (string, error) {
	return "Apple Inc.", nil
}

func Os() string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "sw_vers")
	data, err := cmd.Output()
	if err != nil {
		return runtime.GOOS
	}
	var buf []string
	for line := range strings.Lines(string(data)) {
		if t := strings.Split(line, ":"); len(t) > 1 {
			buf = append(buf, strings.TrimSpace(t[1]))
		}
	}
	os := strings.Join(buf, ".")
	if len(os) == 0 {
		return runtime.GOOS
	}

	return os
}

func NetworkInterface() ([]model.Nic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	cmd := exec.CommandContext(ctx, "ifconfig", "en0")
	defer cancel()

	var result []model.Nic
	out, err := cmd.Output()
	if err != nil {
		return result, err
	}
	// fmt.Println(string(out))

	//% ifconfig en0
	//	en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
	//    options=6460<TSO4,TSO6,CHANNEL_IO,PARTIAL_CSUM,ZEROINVERT_CSUM>
	//    ether ba:7d:ed:b0:51:03
	//    inet6 fe80::400:2c3e:bd0b:cdc2%en0 prefixlen 64 secured scopeid 0xf
	//    inet 192.168.1.13 netmask 0xffffff00 broadcast 192.168.1.255
	//    nd6 options=201<PERFORMNUD,DAD>
	//    media: autoselect
	//    status: active
	macPattern := regexp.MustCompile(`ether ([\w]{2}:[\w]{2}:[\w]{2}:[\w]{2}:[\w]{2}:[\w]{2})`)
	addrPattern := regexp.MustCompile(`inet ([\d]+\.[\d]+\.[\d]+\.[\d]+)`)
	macMatch := macPattern.FindStringSubmatch(string(out))
	nic := model.Nic{}
	if macMatch != nil {
		nic.MAC = macMatch[1]
	}
	addrMatch := addrPattern.FindStringSubmatch(string(out))
	if addrMatch != nil {
		nic.IP = addrMatch[1]
	}
	result = append(result, nic)
	return result, nil
}
