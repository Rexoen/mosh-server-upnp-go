package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"

	upnp "github.com/rexoen/mosh-server-upnp-go/utils"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	args := os.Args[1:]
	out, err := exec.Command("mosh-server", args...).Output()
	if err != nil {
		log.Fatal("Error happend ", "\n", err)
	}
	serverCmdOutput := string(out)
	r := regexp.MustCompile(`MOSH\W+CONNECT\W+(\d+)`)
	port := r.FindStringSubmatch(serverCmdOutput)[1]
	localIP := GetLocalIP()
	log.Println(port)
	log.Println(localIP)
	port_int, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("strconv failed: ", err)
	}
	upnp.ForwardPort(localIP, port_int)
	if err != nil {
		log.Fatal("Error discovering UPnP servers: ", err)
	}

	fmt.Println(serverCmdOutput)
}
