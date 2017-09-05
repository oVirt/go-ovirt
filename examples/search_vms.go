package main

import (
	"fmt"
	"time"

	"github.com/imjoey/go-ovirt"
)

func main() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("qwer1234").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %s\n", err.Error())
	}
	defer conn.Close()

	// Get the reference to the "vms" service:
	vmsService := conn.SystemService().VmsService()

	// Use the "list" method of the "vms" service to list all the virtual machines of the system:
	vmsResponse, err := vmsService.List().Search("name=RHEL7.3").CaseSensitive(false).Send()

	if err != nil {
		fmt.Printf("Failed to get vm list, reason: %v\n", err)
		return
	}
	if vms, ok := vmsResponse.Vms(); ok {
		// Print the virtual machine names and identifiers:
		for _, vm := range vms.Slice() {
			fmt.Print("VM: (")
			if vmName, ok := vm.Name(); ok {
				fmt.Printf(" name: %v", vmName)
			}
			if vmID, ok := vm.Id(); ok {
				fmt.Printf(" id: %v", vmID)
			}
			fmt.Println(")")
		}
	}
}
