package platform

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"time"
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

func Manufacturer() (string, error){
	return "Apple Inc.", nil
}