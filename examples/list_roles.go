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

	roleService := conn.SystemService().RolesService()
	resp, err := roleService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get role list, reason: %v\n", err)
		return
	}

	if roleSlice, ok := resp.Roles(); ok {
		for _, role := range roleSlice.Slice() {
			fmt.Printf("Role: (")
			if roleName, ok := role.Name(); ok {
				fmt.Printf(" name: %v", roleName)
			}
			if roleDesc, ok := role.Description(); ok {
				fmt.Printf(" desc: %v", roleDesc)
			}
			fmt.Println(")")
		}
	}

}
