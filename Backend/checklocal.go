package main

import (
	"log"
	"math"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"libvirt.org/go/libvirt"
)

type vm struct {
	UUID  string  `json:"uuid"`
	Name  string  `json:"name"`
	State string  `json:"state"`
	Vcpus uint    `json:"vcpus"`
	Ram   float64 `json:"ram"`
	Date  string  `json:"date"`
}

var vmlists []vm

func getvm(c *gin.Context) {
	dom := getdomain()
	getlist(dom)
	c.IndentedJSON(http.StatusOK, vmlists)
	vmlists = vmlists[:0]
}

func queryvm(c *gin.Context) {
	dom := getdomain()
	uuid := c.Query("uuid")
	filterlist(uuid, dom)
	c.IndentedJSON(http.StatusOK, vmlists)
	vmlists = vmlists[:0]
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

func getdomain() []libvirt.Domain {
	conn, err := libvirt.NewConnectReadOnly("qemu:///system")
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)

	}
	defer conn.Close()
	if err != nil {
		log.Fatalf("", err)

	}
	alldoms, err := conn.ListAllDomains(0)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}
	return alldoms
}
func getlist(alldoms []libvirt.Domain) {
	for _, dom := range alldoms {
		uuid, err := dom.GetUUIDString()
		info, err := dom.GetInfo()
		name, err := dom.GetName()
		cmd := exec.Command("date")
		date, err := cmd.Output()

		if err != nil {
			log.Fatalf("", err)

		}
		state := stateswitches(info.State)
		vms := vm{uuid, name, state, info.NrVirtCpu, math.Round(float64(info.Memory) / 1000000), string(date)}
		vmlists = append(vmlists, vms)
		dom.Free()

	}

}
func filterlist(queryuuid string, alldoms []libvirt.Domain) {
	for _, dom := range alldoms {
		uuid, err := dom.GetUUIDString()
		info, err := dom.GetInfo()
		name, err := dom.GetName()
		cmd := exec.Command("date")
		date, err := cmd.Output()

		if err != nil {
			log.Fatalf("", err)

		}
		if queryuuid == uuid {
			state := stateswitches(info.State)
			vms := vm{uuid, name, state, info.NrVirtCpu, math.Round(float64(info.Memory) / 1000000), string(date)}
			vmlists = append(vmlists, vms)
			dom.Free()
		}
	}
}
func shutdown(queryuuid string, alldoms []libvirt.Domain) {
	for _, dom := range alldoms {
		uuid, err := dom.GetUUIDString()
		if err != nil {
			log.Fatalf("", err)
		}
		if queryuuid == uuid {
			dom.Shutdown()
			dom.Free()
		}
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/vm", getvm)
	router.GET("/api/vm/1", queryvm)
	router.Run("localhost:8080")

}
