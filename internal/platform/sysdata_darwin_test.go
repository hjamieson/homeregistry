package platform

import (
	"regexp"
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

func TestAppleInc(t *testing.T){
		pat := regexp.MustCompile(`<([\w\s \.]+)>`)
	sample := "here is <Apple Inc.> the end"

	toks := pat.FindStringSubmatch(sample)
	if toks[1] != "Apple Inc." {
		t.Errorf("want %q, got %q", "Apple Inc.", toks)
	}

}

func TestGetNic(t *testing.T){
	nics, err := NetworkInterface()
	if err != nil {
		t.Fatal(err)
	}
	if nics[0].IP == "" {
		t.Errorf("IP is nil")
	}
	if nics[0].MAC == "" {
		t.Error("MAC is nil")
	}
	t.Log(nics[0])
}