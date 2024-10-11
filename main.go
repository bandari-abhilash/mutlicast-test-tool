package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var options string
	fmt.Println("----Select from the following options-----")
	fmt.Println("1. Multicast Sender(Will initiate multicast data)")
	fmt.Println("2. Multicast Receiver(Will receive Multicast Data)")
	fmt.Println("3. Exit")

	fmt.Scan(&options)

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
	if err != nil {
		fmt.Println("Error while listing Interfaces")
	}

	fmt.Printf("The interfaces on this system are %v", interfaceList)

}
