package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/digitalocean/go-libvirt"
	"github.com/gin-gonic/gin"
)

func getvm(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vmlists)
}

var vmlists []string

func main() {
	// This dials libvirt on the local machine, but you can substitute the first
	// two parameters with "tcp", "<ip address>:<port>" to connect to libvirt on
	// a remote machine.
	c, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	v, err := l.Version()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	domains, err := l.Domains()
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	fmt.Println("ID\tName\t\tUUID")
	fmt.Printf("--------------------------------------------------------\n")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)

		vmlists = append(vmlists, string(d.ID), string(d.Name), string(d.Network))

	}

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}
	router := gin.Default()
	router.GET("/vm", getvm)

	router.Run("localhost:8080")
}
