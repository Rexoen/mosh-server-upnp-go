package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"regexp"
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
	lanIP := GetLocalIP()
	fmt.Println(GetLocalIP())
	out, err := exec.Command("mosh-server", args...).Output()
	if err != nil {
		log.Fatal("Error happend ", "\n", err)
	}
	r := regexp.MustCompile(`MOSH\W+CONNECT\W+(\d+)`)
	cmd_output := string(out)
	port := r.FindStringSubmatch(cmd_output)[1]
	fmt.Println("running upnpc with arguments:")
	upnpc_args := []string{"-e", "mosh", "-a", lanIP, port, port, "UDP"}
	fmt.Println(upnpc_args)
	out, err = exec.Command("upnpc", upnpc_args...).CombinedOutput()
	if err != nil {
		log.Fatal("Error happend ", "\n", err)
	}
	fmt.Println(string(out))
	fmt.Println(cmd_output)
}
