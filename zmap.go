package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

func connect(path, address string) {
	cmd := exec.Command(path, address)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}

func message(message string) {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("MessageBoxW")
	title, _ := syscall.UTF16PtrFromString("Zmap")
	text, _ := syscall.UTF16PtrFromString(message)

	ret, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0x00000030)
	fmt.Printf("Return: %d\n", ret)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No parameters")
		message("No parameters")
		return
	}
	if strings.Index(args[0], "zmap://") == -1 || strings.Index(args[0], "@") == -1 {
		fmt.Println("Not correct address")
		message("Not correct address")
		return
	}
	uri := strings.Replace(args[0], "zmap://", "", 1)
	uri = strings.Replace(uri, "/", "", 1)
	address := strings.Split(uri, "@")

	switch address[0] {
	case "rdp":
		connect("mstsc", "/v:"+address[1])
	case "winbox":
		connect("C:\\zmap\\winbox.exe", address[1])
	case "ssh":
		connect("C:\\zmap\\putty.exe", address[1])
	case "vnc":
		connect("C:\\zmap\\vnc.exe", address[1])
	}
}
