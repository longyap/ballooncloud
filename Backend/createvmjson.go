package main

import (
	"encoding/json"
	"fmt"

	"github.com/libvirt/libvirt-go"
)

type VM struct {
	Name       string `json:"name"`
	MemorySize uint   `json:"memory_size"`
	NumCPUs    uint   `json:"num_cpus"`
	OS         struct {
		Type struct {
			Arch    string `json:"arch"`
			Machine string `json:"machine"`
			Name    string `json:"name"`
		} `json:"type"`
		Boot struct {
			Device string `json:"device"`
		} `json:"boot"`
	} `json:"os"`
	Disks []struct {
		Type   string `json:"type"`
		Source struct {
			File string `json:"file"`
		} `json:"source"`
		Target struct {
			Device string `json:"device"`
			Bus    string `json:"bus"`
		} `json:"target"`
		Driver struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"driver"`
	} `json:"disks"`
	Networks []struct {
		Type    string `json:"type"`
		MAC     string `json:"mac"`
		Source  string `json:"source"`
		Model   string `json:"model"`
		Address string `json:"address"`
	} `json:"networks"`
}

func main() {
	// Connect to the local libvirt daemon
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Printf("Failed to connect to libvirt: %v", err)
		return
	}
	defer conn.Close()

	// Define the virtual machine from JSON
	jsonString := `
	{
		"name": "testvm",
		"memory_size": 1048576,
		"num_cpus": 1,
		"os": {
			"type": {
				"arch": "x86_64",
				"machine": "pc-i440fx-2.9",
				"name": "hvm"
			},
			"boot": {
				"device": "hd"
			}
		},
		"disks": [
			{
				"type": "file",
				"source": {
					"file": "/path/to/image.qcow2"
				},
				"target": {
					"device": "vda",
					"bus": "virtio"
				},
				"driver": {
					"name": "qemu",
					"type": "qcow2"
				}
			}
		],
		"networks": [
			{
				"type": "network",
				"mac": "52:54:00:00:00:01",
				"source": "default",
				"model": "virtio",
				"address": ""
			}
		]
	}`

	var vm VM
	err = json.Unmarshal([]byte(jsonString), &vm)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
		return
	}

	// Convert the VM struct to XML
	xml, err := jsonToXML(vm)
	if err != nil {
		fmt.Printf("Failed to convert VM to XML: %v", err)
		return
	}

	// Create the virtual machine
	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		fmt.Printf("Failed to define virtual machine: %v", err)
		return
	}
	defer dom.Free()

	// Start the virtual machine
	err = dom.Create()
	if err != nil {
		fmt.Printf("Failed to start virtual machine: %v", err)
		return
	}

	fmt.Println("Virtual machine created and started successfully.")
}
