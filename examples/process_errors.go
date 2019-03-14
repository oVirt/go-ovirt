package examples

import (
	"fmt"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func processAuthError() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("wrong-password"). // Incorrect password
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Get the reference to the "datacenters" service:
	_, err = conn.SystemService().DataCentersService().List().Send()

	if err != nil {
		if _, ok := err.(*ovirtsdk4.AuthError); ok {
			fmt.Printf("Failed to authenticate user\n")
			return
		}
		fmt.Printf("Failed to get datacenter list, reason: %v\n", err)
		return
	}
}

func processNotFoundError() {
	inputRawURL := "https://10.1.111.229/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("correct-password").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Get the reference to the "datacenters" service:
	_, err = conn.SystemService().DataCentersService().
		DataCenterService("not-exists").
		Get().
		Send()

	if err != nil {
		if _, ok := err.(*ovirtsdk4.NotFoundError); ok {
			fmt.Printf("DataCenter not found\n")
			return
		}
		fmt.Printf("Failed to get datacenter, reason: %v\n", err)
		return
	}
}
