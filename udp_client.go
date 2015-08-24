package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

var packet_num int = 0

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	serv_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)

	for {
		actions(conn)
	}

	defer conn.Close()

}

func actions(conn *net.UDPConn) {
	msg := strconv.Itoa(packet_num)
	buffer := []byte(msg)
	fmt.Println("Sending", packet_num)
	_, err := conn.Write(buffer)
	i := make(chan int, 20)
	go wait_for_ack(packet_num)
	select {
	case <-i:
		return
	case <-time.After(5 * time.Second):
		fmt.Println("timed out for", packet_num)
		go actions(conn)
	}
	packet_num := <-i
	packet_num++
	fmt.Println("Next to send:", i)
	if err != nil {
		fmt.Println(msg, err)
	}
	time.Sleep(time.Second * 1)

}

