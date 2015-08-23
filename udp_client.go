package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	serv_addr,err := net.ResolveUDPAddr("udp","127.0.0.1:10001")
	CheckError(err)

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)

	defer conn.Close()
	i := 0
	for {
		msg := strconv.Itoa(i)
		buffer := []byte(msg)
		fmt.Println("Sending", i)
		_,err := conn.Write(buffer)
		i = wait_for_ack(i)
		i++
		fmt.Println("Next to send:", i)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}
}