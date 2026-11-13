package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hjamieson/homeregistry/internal/model"
)

func Demo1() {
	dev := model.Device{
		Id:       "123456789",
		Hostname: "klaatu",
		OS:       "linux",
		NIC: []model.Nic{
			{IP: "192.168.1.13", MAC: "ba:7d:ed:b0:51:03"},
			{IP: "192.168.1.14", MAC: "fe80::400:2c3e:bd0b:cdc2"},
			{IP: "fe80::400:2c3e:bd0b:cdc2", },
		},
		Manuf: "Gigabyte",
		Class: "laptop",
		Contact: "billy bob",
		LastInfo: time.Now(),
	}

	check, err := json.Marshal(dev)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("result: %v \n", string(check))
}

func Demo2() {
	dev := model.Device{
		Id:       "123456789",
		Hostname: "klaatu",
	}

	check, err := json.Marshal(dev)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("result: %v \n", string(check))
}
