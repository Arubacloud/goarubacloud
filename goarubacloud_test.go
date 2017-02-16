package goarubacloud

import (
    "testing"
)

const (
	endpoint = "dc1"
	username = "ARU-xxxx"
	password = "xxxx"
)

func TestGetTemplate(t *testing.T) {
	client, err := NewAPI(endpoint, username, password)
	_ , err = client.GetTemplate("WS12-002_W2K12R2_1_0")
	if err != nil {
		t.Errorf("Template not found: " + err.Error())
	} 
}

func TestGetPackage(t *testing.T) {
	client, err := NewAPI(endpoint, username, password)
	_ , err = client.GetPreconfiguredPackage("Large")
	if err != nil {
		t.Errorf("Package not found: " + err.Error())
	} 
}
