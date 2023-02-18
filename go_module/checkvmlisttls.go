package main

import (
	"crypto/tls"
	"crypto/x509"

	"fmt"
	"io/ioutil"
	"log"

	"github.com/digitalocean/go-libvirt"
)

func main() {
	// This dials libvirt on the local machine
	// It connects to libvirt via TLS over TCP
	// To connect to a remote machine, you need to have the ca/cert/key of it.
	keyFileXML, err := ioutil.ReadFile("/etc/pki/libvirt/private/clientkey.pem")
	if err != nil {
		log.Fatalf("%v", err)
	}

	certFileXML, err := ioutil.ReadFile("/etc/pki/libvirt/clientcert.pem")
	if err != nil {
		log.Fatalf("%v", err)
	}

	caFileXML, err := ioutil.ReadFile("/etc/pki/CA/cacert.pem")
	if err != nil {
		log.Fatalf("%v", err)
	}
	cert, err := tls.X509KeyPair([]byte(certFileXML), []byte(keyFileXML))
	if err != nil {
		log.Fatalf("%v", err)
	}

	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM([]byte(caFileXML))

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      roots,
	}

	// Use host name or IP which is valid in certificate
	addr := "10.10.10.10"
	port := "16514"
	c, err := tls.Dial("tcp", addr+":"+port, config)
	if err != nil {
		log.Fatalf("failed to dial libvirt: %v", err)
	}

	// Drop a byte before libvirt.New(c)
	// More details at https://github.com/digitalocean/go-libvirt/issues/89
	// Remove this line if the issue does not exist any more
	c.Read(make([]byte, 1))

	l := libvirt.New(c)
	if err := l.Connect(); err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	v, err := l.Version()
	if err != nil {
		log.Fatalf("failed to retrieve libvirt version: %v", err)
	}
	fmt.Println("Version:", v)

	// Return both running and stopped VMs
	flags := libvirt.ConnectListDomainsActive | libvirt.ConnectListDomainsInactive
	domains, _, err := l.ConnectListAllDomains(1, flags)
	if err != nil {
		log.Fatalf("failed to retrieve domains: %v", err)
	}

	fmt.Println("ID\tName\t\tUUID")
	fmt.Println("--------------------------------------------------------")
	for _, d := range domains {
		fmt.Printf("%d\t%s\t%x\n", d.ID, d.Name, d.UUID)
	}

	if err := l.Disconnect(); err != nil {
		log.Fatalf("failed to disconnect: %v", err)
	}
}
