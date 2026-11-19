package platform

import (
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

