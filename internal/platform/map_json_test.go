package platform

import (
	"encoding/json"
	"testing"
)

func TestMapJson(t *testing.T) {
	m := map[string]string{"os": "linux", "id": "123", "manuf": "apple"}
	d, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(d))

	min := make(map[string]string)
	err = json.Unmarshal(d, &min)
	if err != nil {
		t.Error(err)
	}
	t.Log(min)

}

func TestExtraProps(t *testing.T) {
	type Data struct {
		ID   string
		Name string
	}

	m := map[string]string{"ID": "123", "Name": "Larry", "extra": "garbage"}
	mout, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	}
	data := Data{}
	err = json.Unmarshal(mout, &data)

	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v", data)
}
