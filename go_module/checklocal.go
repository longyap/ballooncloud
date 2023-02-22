package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"libvirt.org/go/libvirt"
)

type vm struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	State string `json:"state"`
	Vcpus uint   `json:"vcpus"`
	Ram   uint64 `json:"ram"`
}

var vmlists []vm

func getvm(c *gin.Context) {
	getlist()
	c.IndentedJSON(http.StatusOK, vmlists)
}
func stateswitches(no libvirt.DomainState) string {
	switch no {
	case libvirt.DOMAIN_NOSTATE:
		return "No State"
	case libvirt.DOMAIN_RUNNING:
		return "Running"
	case libvirt.DOMAIN_BLOCKED:
		return "Blocked"
	case libvirt.DOMAIN_PAUSED:
		return "Paused"
	case libvirt.DOMAIN_SHUTDOWN:
		return "Shutting Down"
	case libvirt.DOMAIN_SHUTOFF:
		return "Shut Off"
	case libvirt.DOMAIN_CRASHED:
		return "Crashed"
	default:
		return "Unknown"
	}

}

func getlist() {
	conn, err := libvirt.NewConnectReadOnly("qemu:///system")
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)

	}
	defer conn.Close()
	alldoms, err := conn.ListAllDomains(0)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	for _, dom := range alldoms {
		uuid, err := dom.GetUUIDString()
		info, err := dom.GetInfo()
		name, err := dom.GetName()
		if err != nil {
			log.Fatalf("", err)

		}

		state := stateswitches(info.State)
		vms := vm{uuid, name, state, info.NrVirtCpu, info.Memory}
		vmlists = append(vmlists, vms)
		dom.Free()
	}

}
func main() {
	router := gin.Default()
	router.GET("/vm", getvm)

	router.Run("localhost:8080")
}
