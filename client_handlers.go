package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

var g_clients []net.Conn

func ClientHandler(port int) {
	ipCo, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	
	for {
		conn, _ := ipCo.Accept()
		g_clients = append(g_clients, conn)
		go MessageHandler(conn)
	}
}

func MessageHandler(co net.Conn) {
	for {
		message, _ := bufio.NewReader(co).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		
		for i := 0; i < len(g_clients); i++ {
			g_clients[i].Write([]byte(newmessage + "\n"))
		}
	}
}
