package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	// "net/http"
)

type interfaces struct {
	index         int
	interfaceName string
	interfaceIp   string
}

var interfaceOption string
var transmitterOptions string

func main() {
	var interfacesLists sync.Map

	interfaceOption = listInterfaces(&interfacesLists, interfaceOption)
	value, ok := interfacesLists.Load(interfaceOption)
	if data, noValue := value.(interfaces); noValue && data.interfaceName == "EXIT" {
		os.Exit(0)
	}
	if !ok {
		fmt.Printf("Please select a valid option, The selected option %s is not listed \n\n", interfaceOption)
		interfaceOption = listInterfaces(&interfacesLists, interfaceOption)
		value, ok := interfacesLists.Load(interfaceOption)
		if data, noValue := value.(interfaces); noValue && data.interfaceName == "EXIT" {
			os.Exit(0)
		}
		if !ok {
			fmt.Printf("Please select a valid option, The selected option %s is not listed \n\n", interfaceOption)
			main()
		}
	}

	selectOptions(transmitterOptions)

}

func listInterfaces(interfacesLists *sync.Map, interfaceOption string) (intValue string) {
	interfaceList, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error while listing Interfaces")
	}

	fmt.Printf("Select any of the following interfaces--------------- \n\n")

	var count int = 0
	for _, iface := range interfaceList {
		if iface.Flags&net.FlagUp != 0&net.FlagBroadcast&net.FlagMulticast {
			addrs, errA := iface.Addrs()
			if errA != nil {
				log.Println("Unable to get Interface address", errA.Error())
			}
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok {
					count += 1
					tempObj := interfaces{
						index:         count,
						interfaceName: iface.Name,
						interfaceIp:   fmt.Sprintf("%-40s", ipNet.IP.String()),
					}
					interfacesLists.Store(strconv.Itoa(count), tempObj)
					fmt.Printf("%d   %-40s        %s   \n", count, ipNet.IP.String(), iface.Name)
				}
			}
		}
	}
	fmt.Printf("%d   %-40s        %s   \n", count+1, "EXIT", "")
	interfacesLists.Store(strconv.Itoa(count+1), interfaces{index: count + 1, interfaceName: "EXIT"})

	fmt.Printf("\n\n")
	fmt.Scan(&interfaceOption)
	return interfaceOption
}

func selectOptions(transmitterOptions string) {

	fmt.Println("----Select from the following options-----")
	fmt.Println("1. Multicast Sender(Will initiate multicast data)")
	fmt.Println("2. Multicast Receiver(Will receive Multicast Data)")
	fmt.Println("3. Exit")

	fmt.Scan(&transmitterOptions)
	switch transmitterOptions {
	case "1":
		fmt.Println("You have selected option 1")
	case "2":
		fmt.Println("You have selected option 2")
	case "3":
		os.Exit(0)
	default:
		fmt.Println("Please select from the above options")
	}
}

func mutlicastSender() {
	fmt.Println("Sending Mutlicast packets")
}
