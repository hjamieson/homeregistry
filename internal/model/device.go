package model

import (
	"time"
)

type Nic struct {
	IP  string `json:",omitempty"`
	MAC string `json:",omitempty"`
}

type Device struct {
	Id       string
	Hostname string
	NIC      []Nic     `json:",omitempty"`
	Manuf    string    `json:",omitempty"`
	OS       string    `json:",omitempty"`
	Class    string    `json:",omitempty"`
	Contact  string    `json:",omitempty"`
	LastInfo time.Time `json:",omitzero"`
}
