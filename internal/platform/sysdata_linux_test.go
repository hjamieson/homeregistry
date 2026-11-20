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

func TestGetNic(t *testing.T) {
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
