// Package checkport is a small tool to check if a port is available or not.
//
// Examples/readme can be found on the GitHub page at https://github.com/ewilan-riviere/checkport
//
// If you want to use it as CLI, you can install it with:
//
//	go install github.com/ewilan-riviere/checkport
//
// Then you can use it like this:
//
//	checkport 3000
//
// If you use `sudo`, you can have more info about process if port is used:
//
//	sudo checkport 3000
package main

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <PORT>\n", os.Args[0])
		return
	}

	port, _ := strconv.Atoi(os.Args[1])
	currentUser, _ := user.Current()
	isRoot := currentUser.Uid == "0"

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	connections, err := net.ConnectionsWithContext(ctx, "all")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	found := false
	for _, conn := range connections {
		if conn.Laddr.Port == uint32(port) {
			found = true

			transport := "TCP"
			if conn.Type == 2 { // `2` generally corresponds to SOCK_DGRAM (UDP)
				transport = "UDP"
			}

			fmt.Printf("❌ Port %d [%s] is used.\n", port, transport)

			if isRoot && conn.Pid != 0 {
				p, err := process.NewProcessWithContext(ctx, conn.Pid)
				if err == nil {
					name, _ := p.NameWithContext(ctx)
					fmt.Printf("🔍 Process: %s (PID: %d)\n", name, conn.Pid)
				}
			} else if conn.Pid == 0 {
				fmt.Println("ℹ️  Unable to determine the PID (often due to the kernel or TIME_WAIT).")
			} else {
				fmt.Println("ℹ️  Execute with `sudo` to identify the process.")
			}
			break
		}
	}

	if !found {
		fmt.Printf("✅ Port %d is available.\n", port)
	}
}
