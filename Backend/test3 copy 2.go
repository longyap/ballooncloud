package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/libvirt/libvirt-go"
)

type VMConfig struct {
	Name    string `json:"name"`
	Memory  int    `json:"memory"`
	VCPUs   int    `json:"vcpus"`
	Image   string `json:"image"`
	Network string `json:"network"`
}

func main() {
	// Read configuration from JSON file
	jsonFile := "vmconfig.json"
	configFile, err := os.Open(jsonFile)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	var config VMConfig
	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		panic(err)
	}

	// Connect to Libvirt
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Generate XML configuration
	xml := fmt.Sprintf(`<domain type='kvm'>
              <name>%s</name>
              <memory unit='KiB'>%d</memory>
              <vcpu placement='static'>%d</vcpu>
              <os>
                <type arch='x86_64' machine='pc-i440fx-2.12'>hvm</type>
                <boot dev='hd'/>
              </os>
              <devices>
                <disk type='file' device='disk'>
                  <driver name='qemu' type='qcow2'/>
                  <source file='%s'/>
                  <target dev='vda' bus='virtio'/>
                </disk>
                <interface type='network'>
                  <mac address='52:54:00:00:00:01'/>
                  <source network='%s'/>
                  <model type='virtio'/>
                </interface>
              </devices>
            </domain>`, config.Name, config.Memory, config.VCPUs, config.Image, config.Network)

	// Define and start the domain
	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		panic(err)
	}
	if err := dom.Create(); err != nil {
		panic(err)
	}

	fmt.Println("VM created successfully!")
}
