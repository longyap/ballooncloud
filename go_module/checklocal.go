package main

import (
	"fmt"
	"log"

	"libvirt.org/go/libvirt"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)

	}
	defer conn.Close()
	//	indoms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	//	if err != nil {
	//		log.Fatalf("failed to dial libvirt: %v", err)
	//	}
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	for _, dom := range doms {
		uuid, err := dom.GetUUIDString()
		vcpu, err := dom.GetVcpus()
		vcpuinfo, err := dom.GetOSType()
		name, err := dom.GetName()

		if err == nil {
			fmt.Printf(vcpuinfo, "\n", uuid, "\n", name, "  %s\n", vcpu)
		}
		dom.Free()
	}

	fmt.Printf("\n%d running domains:\n", len(doms))
	for _, dom := range doms {
		name, err := dom.GetName()

		if err == nil {
			fmt.Printf("  %s\n", name)
		}
		dom.Free()
	}
}
