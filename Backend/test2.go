package main

import (
	"fmt"

	"libvirt.org/go/libvirt"
)

func main() {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Failed to connect to libvirt:", err)
		return
	}
	defer conn.Close()

	dom, err := conn.LookupDomainByUUIDString("84b00e85-b74e-4376-bc6f-d70d109db94a")
	if err != nil {
		fmt.Println("Failed to find domain:", err)
		return
	}
	defer dom.Free()

	if err := dom.Create(); err != nil {
		fmt.Println("Failed to destroy domain:", err)
		return
	}

	fmt.Println("Domain startup successfully")
}
