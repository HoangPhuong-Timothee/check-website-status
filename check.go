package main

import (
	"fmt"
	"net"
	"time"
)

func Check(d, p string) string {
	address := d + ":" + p
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable.\n Error: %v", d, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable.\n From: %v\n To: %v", d, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
