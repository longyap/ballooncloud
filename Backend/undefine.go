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

	dom, err := conn.LookupDomainByUUIDString("242a70fb-9876-463a-9343-46342b8da766")
	if err != nil {
		fmt.Println("Failed to find domain:", err)
		return
	}
	defer dom.Free()

	if err := dom.Undefine(); err != nil {
		fmt.Println("Failed to destroy domain:", err)
		return
	}

	fmt.Println("Domain deleted successfully")
}
