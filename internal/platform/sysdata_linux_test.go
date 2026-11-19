package platform

import "testing"

func TestRegex(t *testing.T) {
	for _, s := range []string{"!!!!", "hello world", "<script>"} {
		if uuidPattern.MatchString(s) {
			t.Errorf("%q matched, but shouldn't", s)
		}
	}
}

func TestHostnameCtl(t *testing.T) {
	props, err := HostnameCtl()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(props)
	if len(props) == 0 {
		t.Errorf("0 properties returned")
	}
}
