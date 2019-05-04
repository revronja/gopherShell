package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		log.Fatal("Error %s", e)
	}
}

// func takes in net conn and writes it back
func handleNetConn(conn net.Conn) {
	// read incoming connection data
	fmt.Printf("Receiving connection from %s\n", conn.RemoteAddr().String())

	for {
		connData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			CheckErr(err)
		}

		dataStrings := strings.TrimSpace(string(connData))
		if dataStrings == "q" {
			break
		}
		fmt.Printf("data received from client: %s", dataStrings)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		cmd, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "%s\n", cmd)

		conn.Write([]byte(dataStrings + "\n"))
	}
	conn.Close()

}

func checkPort() (bool, error) {
	return false, nil
}

var helpMsg = `
   **  GoTCP Server Help Display **
   --help 


`

// main
func main() {
	// read args
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
		return
	}
	// and read flags
	//var helpStr string
	//flag.String("help")
	LPORT := flag.String("p", "", "port to listen on")
	fmt.Printf("lport is %s", *LPORT)
	flag.Parse()

	l, err := net.Listen("tcp4", fmt.Sprintf("127.0.0.1:%s", *LPORT))
	CheckErr(err)

	fmt.Printf("Listening on %s for incoming connections\n", *LPORT)
	defer l.Close()
	for {
		c, err := l.Accept()
		CheckErr(err)
		go handleNetConn(c)
	}
}
