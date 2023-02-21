package main

import (
	"fmt"
	"log"

	"libvirt.org/go/libvirt"
)

func main() {
	conn, err := libvirt.NewConnectReadOnly("qemu:///system")
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)

	}
	hname, err := conn.GetHostname()
	defer conn.Close()

	indoms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_INACTIVE)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	alldoms, err := conn.ListAllDomains(0)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	for _, dom := range alldoms {
		name, err := dom.GetName()
		uuid, err := dom.GetUUIDString()
		status, reason, err := dom.GetState()
		info, err := dom.GetInfo()

		if err != nil {
		}
		fmt.Println(name, "no1 \n", uuid, "no2 \n", status, "%no3 \n", reason, "\n", info.NrVirtCpu, "\n", info.Memory)
	}

	for _, dom := range indoms {

		dom.Free()
	}

	fmt.Printf("\n%d running domains:\n", len(indoms))
	for _, dom := range doms {
		name, err := dom.GetName()

		if err == nil {
			fmt.Printf(hname, "  %s\n", name)
		}
		dom.Free()
	}
}
