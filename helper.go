package ovirtsdk

import (
	"fmt"
	"time"
)

const DefaultInterval = 10 * time.Second

const DefaultVMTimeout = 120 * time.Second

// WaitForVM waits for VM to given status
func (c *Connection) WaitForVM(vmID string, status VmStatus, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = DefaultVMTimeout
	}
	if vmID == "" {
		return fmt.Errorf("Invalid VM ID")
	}
	vmService := c.SystemService().VmsService().VmService(vmID)
	for {
		resp, err := vmService.Get().Send()
		if err != nil {
			return err
		}
		if timeout <= 0 {
			return fmt.Errorf("Timeout for waiting for VM to %v", status)
		}

		vm, ok := resp.Vm()
		if !ok {
			continue
		}
		if vm.MustStatus() == status {
			break
		}

		timeout = timeout - DefaultInterval
		time.Sleep(DefaultInterval)
	}

	return nil
}
