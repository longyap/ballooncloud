package main

import (
	"fmt"

	"github.com/libvirt/libvirt-go"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	xml := `<domain type='kvm'>
              <name>myvm</name>
              <memory unit='KiB'>1048576</memory>
              <vcpu placement='static'>1</vcpu>
              <os>
                <type arch='x86_64' machine='pc-i440fx-2.12'>hvm</type>
                <boot dev='hd'/>
              </os>
              <devices>
                <disk type='file' device='disk'>
                  <driver name='qemu' type='qcow2'/>
                  <source file='/path/to/disk/image.qcow2'/>
                  <target dev='vda' bus='virtio'/>
                </disk>
                <interface type='network'>
                  <mac address='52:54:00:00:00:01'/>
                  <source network='default'/>
                  <model type='virtio'/>
                </interface>
              </devices>
            </domain>`

	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		panic(err)
	}

	if err := dom.Create(); err != nil {
		panic(err)
	}

	fmt.Println("VM created successfully!")

}
