package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
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
var portOption string
var multicastAddress string
var interfacesLists sync.Map

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from panic : ", r)
		}
		main()
	}()

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

	selectOptions(transmitterOptions, interfaceOption, portOption)

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

func selectOptions(transmitterOptions string, interfaceOption string, portOption string) {

	fmt.Println("----Select from the following options-----")
	fmt.Println("1. Multicast Sender(Will initiate multicast data)")
	fmt.Println("2. Multicast Receiver(Will receive Multicast Data)")
	fmt.Println("3. Exit")

	fmt.Scan(&transmitterOptions)
	switch transmitterOptions {
	case "1":
		fmt.Println("Enter multicast address  between (224.0.0.0 to 239.255.255.255) to use")
		fmt.Scan(&multicastAddress)
		fmt.Println("Please enter the port number between 1 and 6445")
		fmt.Scan(&portOption)
		mutlicastSender(interfaceOption, portOption, multicastAddress)
	case "2":
		fmt.Println("Enter multicast address  between (224.0.0.0 to 239.255.255.255) to use")
		fmt.Scan(&multicastAddress)
		fmt.Println("Please enter the port number between 1 and 6445")
		fmt.Scan(&portOption)
		muliticastListener(interfaceOption, portOption, multicastAddress)
	case "3":
		os.Exit(0)
	default:
		fmt.Println("Please select from the above options")
	}
}

func mutlicastSender(interfaceOption string, portOption string, multicastAddress string) {
	testMessage := "This is a multicast test message"
	interfaceDetails, ok := interfacesLists.Load(interfaceOption)
	if !ok {
		fmt.Println("Please enter the proper option")
		os.Exit(0)
	}
	if interfaceDetailsTypeCasted, typeCasted := interfaceDetails.(interfaces); typeCasted {
		port, convErr := strconv.Atoi(portOption)
		if convErr != nil {
			fmt.Println("Error while converting port from string to int")
		}
		fmt.Println("Printing IP-----------", net.IP(strings.TrimSpace(interfaceDetailsTypeCasted.interfaceIp)))
		laddrObj := net.UDPAddr{
			IP:   net.IP(interfaceDetailsTypeCasted.interfaceIp),
			Port: port,
		}
		raddObj := net.UDPAddr{
			IP:   net.IP(multicastAddress),
			Port: port,
		}
		fmt.Println("------------------", net.IP.IsMulticast(net.IP(multicastAddress)))
		udpConn, err := net.DialUDP("udp", &laddrObj, &raddObj)
		if err != nil {
			fmt.Println("Failed to create a udp connection", err.Error())
		}

		buf := []byte(testMessage)
		for {
			fmt.Printf("Sending packets on interface %s and port %d \n", interfaceDetailsTypeCasted.interfaceIp, port)
			if _, err := udpConn.Write(buf); err != nil {
				fmt.Println("Failed to send udp packets", err)
			}
		}
	}

}

func muliticastListener(interfaceOption string, portOption string, multicastAddress string) {
	interfaceDetails, ok := interfacesLists.Load(interfaceOption)
	if !ok {
		fmt.Println("Please enter the proper option")
		os.Exit(0)
	}
	if _, typeCasted := interfaceDetails.(interfaces); typeCasted {
		port, convErr := strconv.Atoi(portOption)
		if convErr != nil {
			fmt.Println("Error while converting port from string to int")
		}
		addrObj := net.UDPAddr{
			IP:   net.IP(multicastAddress),
			Port: port,
		}

		listener, err := net.ListenMulticastUDP("udp", nil, &addrObj)
		if err != nil {
			fmt.Println("Error while listening for multicast", err.Error())
		}
		buffer := make([]byte, 1024)
		for {
			n, _, errO := listener.ReadFromUDP(buffer)
			if errO != nil {
				fmt.Println("Unable to read the packets", errO.Error())
			}
			message := string(buffer[:n])
			fmt.Println(message)
		}
	}
}
