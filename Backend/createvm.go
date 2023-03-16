package main

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

func main() {
	// Connect to the local libvirt daemon
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Printf("Failed to connect to libvirt: %v", err)
		return
	}
	defer conn.Close()

	// Define the virtual machine
	xml := `<domain type='kvm'>
                <name>testvm</name>
                <memory unit='KiB'>1048576</memory>
                <vcpu placement='static'>1</vcpu>
                <os>
                    <type arch='x86_64' machine='pc-i440fx-2.9'>hvm</type>
                    <boot dev='hd'/>
                </os>
                <devices>
                    <disk type='file' device='disk'>
                        <driver name='qemu' type='qcow2'/>
                        <source file='/path/to/image.qcow2'/>
                        <target dev='vda' bus='virtio'/>
                    </disk>
                    <interface type='network'>
                        <mac address='52:54:00:00:00:01'/>
                        <source network='default'/>
                        <model type='virtio'/>
                    </interface>
                </devices>
            </domain>`

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
