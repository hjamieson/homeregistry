package platform

import (
	"strings"
	"testing"
)

func TestRegexUUID(t *testing.T) {
	sample := "      \"IOPlatformUUID\" = \"8DE9B81F-54F4-5479-8029-1571CE7B0265\"\n"
	found := uuidPattern.FindStringSubmatch(sample)
	if found == nil {
		t.Errorf("expected something, got %v", found)
	}
	t.Logf("found: %v", found[1])
}

func TestOsDarwin(t *testing.T) {
	os := Os()
	if !strings.Contains(os, "macOS") {
		t.Errorf("got %q, want %q", os, "macOS")
	}
}
func TestManufacturer(t *testing.T) {
	manuf, err := Manufacturer()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(manuf, "Apple") {
		t.Errorf("got %q, want %q", manuf, "Apple")
	}
}
