package connectors

import (
	"bytes"
	"encoding/json"
	"github.com/reef-pi/reef-pi/controller/utils"
	"testing"
)

func TestInletsAPI(t *testing.T) {
	store, err := utils.TestDB()
	if err != nil {
		t.Fatal(err)
	}
	tr := utils.NewTestRouter()
	i := Inlet{Name: "Foo", Pin: 21}
	inlets := NewInlets(store)
	inlets.DevMode = true
	if err := inlets.Setup(); err != nil {
		t.Fatal(err)
	}
	inlets.LoadAPI(tr.Router)
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(i)
	if err := tr.Do("PUT", "/api/inlets", body, nil); err != nil {
		t.Error(err)
	}
	body.Reset()
	json.NewEncoder(body).Encode(i)
	if err := tr.Do("POST", "/api/inlets/1", body, nil); err != nil {
		t.Error(err)
	}
	if err := tr.Do("GET", "/api/inlets", new(bytes.Buffer), nil); err != nil {
		t.Error(err)
	}
	if err := tr.Do("POST", "/api/inlets/1/read", new(bytes.Buffer), nil); err != nil {
		t.Error(err)
	}
	if err := tr.Do("DELETE", "/api/inlets/1", new(bytes.Buffer), nil); err != nil {
		t.Error(err)
	}
}
