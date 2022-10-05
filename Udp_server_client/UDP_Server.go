package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, startTime time.Time) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	duration := time.Since(startTime)
	fmt.Println("The duration between receiving message and sending message is ", duration)
	if duration.Seconds() >= 5 {
		fmt.Println("Error: The duration is more than 5 seconds.")
	}
	_, err := conn.WriteToUDP([]byte("I am fine thank you!"), addr)// add WriteToUDP function to reply to the client
	if err != nil {
		fmt.Printf("Couldn’t send response %v", err)
	}
}
func main() {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		IP: net.ParseIP("127.0.0.1"),
		Port: 1234,
	}
	for {
		//net.ListenUDP is to listen to a UDP network on a certain IP address, if the IP address sends a message, it will receive the message.
		ser, err := net.ListenUDP("udp", &addr)
		if err != nil { fmt.Printf("Some error %v\n", err)
			return
		}
		for {
			_, remoteaddr, err := ser.ReadFromUDP(p)
			start := time.Now()
			fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
			if err != nil {
				continue
			}
			go sendResponse(ser, remoteaddr ,start)
		}
	}
}
