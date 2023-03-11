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
	alldoms, err := dom.ListAllDomains(0)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	getlist(alldoms)
	c.IndentedJSON(http.StatusOK, vmlists)
	vmlists = vmlists[:0]
	defer dom.Close()
}

func queryvm(c *gin.Context) {
	uuid := c.Query("uuid")
	filterlist(uuid)
	c.IndentedJSON(http.StatusOK, vmlists)
	vmlists = vmlists[:0]
}

func unlistvm(c *gin.Context) {
	conn := getdomain()
	uuid := c.Query("uuid")
	dom, err := conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	if err := dom.Undefine(); err != nil {
		println("Failed to destroy domain:", err)
		return
	}
	sucess := "destroy domain success"
	c.IndentedJSON(http.StatusOK, sucess)

}
func shutdownvm(c *gin.Context) {
	conn := getdomain()
	uuid := c.Query("uuid")
	dom, err := conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	if err := dom.Destroy(); err != nil {
		println("Failed to destroy domain:", err)
		return
	}
	sucess := "destroy domain success"
	c.IndentedJSON(http.StatusOK, sucess)

}
func rebootvm(c *gin.Context) {
	conn := getdomain()
	uuid := c.Query("uuid")
	dom, err := conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	if err := dom.Reboot(libvirt.DOMAIN_REBOOT_DEFAULT); err != nil {
		println("Failed to destroy domain:", err)
		return
	}
	sucess := "restart domain success"
	c.IndentedJSON(http.StatusOK, sucess)

}
func startvm(c *gin.Context) {
	conn := getdomain()
	uuid := c.Query("uuid")
	dom, err := conn.LookupDomainByUUIDString(uuid)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	if err := dom.Create(); err != nil {
		println("Failed to start domain:", err.Error())
		return
	}
	defer conn.Close()
	sucess := "start domain success"
	c.IndentedJSON(http.StatusOK, sucess)

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

func getdomain() *libvirt.Connect {
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)

	}
	//defer conn.Close()
	//alldoms, err := conn.ListAllDomains(0)
	//if err != nil {
	//	log.Fatalf("failed to dial libvirt: %v", err)
	//}
	return conn
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
func filterlist(queryuuid string) {
	getdom := getdomain()
	dom, err := getdom.LookupDomainByUUIDString(queryuuid)
	if err != nil {
		println("Failed to find domain:", err)
		return
	}
	uuid, err := dom.GetUUIDString()
	info, err := dom.GetInfo()
	name, err := dom.GetName()
	cmd := exec.Command("date")
	date, err := cmd.Output()

	state := stateswitches(info.State)
	vms := vm{uuid, name, state, info.NrVirtCpu, math.Round(float64(info.Memory) / 1000000), string(date)}
	vmlists = append(vmlists, vms)
	dom.Free()

}

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/vm", getvm)
	router.GET("/api/vm/instance", queryvm)
	router.GET("/api/vm/instance/start", startvm)
	router.GET("/api/vm/instance/stop", shutdownvm)
	router.GET("/api/vm/instance/reboot", rebootvm)
	router.GET("/api/vm/instance/unlist", unlistvm)
	router.Run("localhost:8080")

}
