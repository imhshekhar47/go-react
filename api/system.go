package api

import (
	"errors"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// PodInfo : data model
type PodInfo struct {
	ID   string `json:"id"`
	IP   string `json:"ip"`
	Name string `json:"name"`
	Node string `json:"node"`
}

// HostInfo : data model
type HostInfo struct {
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ipaddress"`
}

func getHostIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err == nil {
			for _, addr := range addrs {
				var ip net.IP

				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				return ip.String(), nil
			}
		}
	}
	return "", errors.New("Could not find ip")
}

// GetHostInfo : returns host information
func GetHostInfo(ctx *gin.Context) {
	hostname := os.Getenv("APP_NODENAME")
	ipaddr := os.Getenv("APP_NODEIP")

	ctx.JSON(http.StatusOK, HostInfo{
		Hostname:  hostname,
		IPAddress: ipaddr,
	})
}

// GetPodInfo : returns pod information
func GetPodInfo(ctx *gin.Context) {

	pID := os.Getenv("APP_PODID")
	pIP := os.Getenv("APP_PODIP")
	pName := os.Getenv("APP_PODNAME")
	pNode := os.Getenv("APP_NODENAME")

	podInfo := PodInfo{
		ID:   pID,
		IP:   pIP,
		Name: pName,
		Node: pNode,
	}

	ctx.JSON(http.StatusOK, podInfo)
}
