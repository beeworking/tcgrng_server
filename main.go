package main

import (
	"log"
	"fmt"
	"flag"
)

func main() {
	portPtr := flag.Int("port", 8080, "Port of the new Game-Server")

	flag.Parse()
	
	if (*portPtr == 8080) {
		log.Println(fmt.Sprintf("ERROR - Launch a server with option -port"))
		return
	}
		
	
	port := *portPtr
	
	go ClientHandler(port)
	
	for {}
}
