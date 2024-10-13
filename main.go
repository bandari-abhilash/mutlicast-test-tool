package main

import (
	"fmt"
	"net"
	"os"
)

type interfaces struct {
	index         int
	interfaceName string
	interfaceIp   string
	multicast     net.Flags
}

func main() {
	var interfacesLists []interfaces
	var options string
	fmt.Println("----Select from the following options-----")
	fmt.Println("1. Multicast Sender(Will initiate multicast data)")
	fmt.Println("2. Multicast Receiver(Will receive Multicast Data)")
	fmt.Println("3. Exit")

	// fmt.Scan(&options)

	switch options {
	case "1":
		fmt.Println("You have selected option 1")
	case "2":
		fmt.Println("You have selected option 2")
	case "3":
		os.Exit(0)
	default:
		fmt.Println("Please select from the above options")
	}

	interfaceList, err := net.Interfaces()
	val := net.IPv6interfacelocalallnodes

	fmt.Println(val)
	if err != nil {
		fmt.Println("Error while listing Interfaces")
	}
	var count int = 0
	for _, iface := range interfaceList {
		if iface.Flags&net.FlagUp != 0&net.FlagBroadcast&net.FlagMulticast {
			count += 1
			tempObj := interfaces{
				index:         iface.Index,
				interfaceName: iface.Name,
				interfaceIp:   iface.HardwareAddr.String(),
				multicast:     net.FlagMulticast,
			}
			addrs, errA := iface.Addrs()
			if errA != nil {
				fmt.Println("Error:", errA)
			}
			for _, addr := range addrs {
				if _, ok := addr.(*net.IPNet); ok {
					fmt.Println("-------------", addr.String())
					// fmt.Printf("%-40s %s (%s)\n", ipNet.IP.String(), iface.Name, iface.HardwareAddr)
				}
			}
			interfacesLists = append(interfacesLists, tempObj)
			// fmt.Println("----------------", tempObj)
			// fmt.Printf("The interface name %s and the broadcast flag %d   \n", iface.Name, net.FlagBroadcast)
		}
		// fmt.Println("------------", interfacesLists)
		// fmt.Println("Interface name", iface)
		// fmt.Println("hardware address ------------------", iface.HardwareAddr.String())

	}

	fmt.Println("Count is", count)
	// fmt.Println(interfacesLists)

	// fmt.Printf("The interfaces on this system are %v", interfaceList)

}
