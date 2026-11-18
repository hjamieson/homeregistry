package platform

import "testing"

func TestRegex(t *testing.T) {
	for _, s := range []string{"!!!!", "hello world", "<script>"} {
		if uuidPattern.MatchString(s) {
			t.Errorf("%q matched, but shouldn't", s)
		}
	}
}
