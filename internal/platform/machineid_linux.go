package platform

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
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
