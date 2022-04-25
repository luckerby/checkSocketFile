package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"syscall"
	"time"
)

func main() {
	// Direct path on the Kubernetes node
	//socketFile := `\var\lib\kubelet\plugins_registry\secrets-store.csi.k8s.io-reg.sock`
	// Path on the mapped directory inside the hostPath volume attached to the container running this code
	socketFile := `\plugins_registry\secrets-store.csi.k8s.io-reg.sock`

	// Run indefinitely
	for {
		// === Linux code only ===
		if runtime.GOOS == "linux" {
			fmt.Printf("os.ModeSocket= %v\n", os.ModeSocket)
			fmt.Printf("os.ModeSocket [binary]= %b\n", os.ModeSocket)
			// os.Stat against the socket file
			fmt.Printf("Calling os.Stat against %v...", socketFile)
			stat, err := os.Stat(socketFile)
			if err != nil {
				fmt.Printf("error\n")
				fmt.Println(err)
			} else {
				fmt.Printf("ok\n")
				fmt.Printf("Calling Mode() against the FileInfo object returned previously...")
				fileMode := stat.Mode()
				fmt.Printf("done\n")
				// Print data retrieved about the file
				fmt.Printf("[os.Stat]  Socket file mode result=  %v\n", fileMode)
				fmt.Printf("[os.Stat]  Socket file mode result (binary)= %b\n", fileMode)
				fmt.Printf("Bitwise AND between file socket mode (from os.Stat) and os.ModeSocket [binary]= %b\n", fileMode&os.ModeSocket)
				// Print misc data about the file
				fmt.Printf("syscall Win32FileAttributeData= %v\n", stat.Sys().(*syscall.Win32FileAttributeData).FileAttributes)
				fmt.Printf("syscall Win32FileAttributeData [binary]= %b\n", stat.Sys().(*syscall.Win32FileAttributeData).FileAttributes)
			}
			// os.Lstat against the socket file
			fmt.Printf("Calling os.Lstat against %v...", socketFile)
			statWithLstat, err := os.Lstat(socketFile)
			if err != nil {
				fmt.Printf("error\n")
				fmt.Println(err)
			} else {
				fmt.Printf("ok\n")
				fmt.Printf("Calling Mode() against the FileInfo object returned previously...")
				fileMode := statWithLstat.Mode()
				fmt.Printf("done\n")
				// Print data retrieved about the file
				fmt.Printf("[os.Lstat] Socket file mode result= %v\n", fileMode)
				fmt.Printf("[os.Lstat] Socket file mode result (binary)= %b\n", fileMode)
				fmt.Printf("Bitwise AND between socket file mode (from os.Lstat) and os.ModeSocket (binary)= %b\n", fileMode&os.ModeSocket)
			}
		}

		// === Windows code only ===
		if runtime.GOOS == "windows" {
			fmt.Printf("Check if the registration socket file %v exists...", socketFile)
			// Despite os.Stat not detecting if a file is used as a Unix Domain Socket, we
			//  don't really care, as we only want to see if the file exists
			_, err := os.Stat(socketFile)
			if err != nil {
				fmt.Printf("not found\n")
				fmt.Println(err)
			} else {
				fmt.Printf("ok\n")
				fmt.Printf("Dialing against socket file %v...", socketFile)
				c, errDial := net.Dial("unix", socketFile)
				if errDial == nil {
					c.Close()
					fmt.Printf("ok\n")
				} else {
					fmt.Printf("error\n")
					fmt.Println(errDial)
				}
			}
		}

		// Wait a while until printing the stats again
		fmt.Println("Sleeping 5s...")
		time.Sleep(5 * time.Second)
	}
}
