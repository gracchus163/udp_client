package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

var packet_num int = 1

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
	}
}

func main() {
	serv_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	listen_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10002")

	CheckError(err)

	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	conn, err := net.DialUDP("udp", local_addr, serv_addr)
	CheckError(err)

	serv_conn, err := net.ListenUDP("udp", listen_addr)

	CheckError(err)
	for {
		actions(conn, serv_conn)
	}

	defer conn.Close()

}

func actions(conn *net.UDPConn, serv_conn *net.UDPConn) {
	msg := strconv.Itoa(packet_num)
	buffer := []byte(msg)
	fmt.Println("Sending", packet_num)
	_, err := conn.Write(buffer)
	i := make(chan int, 20)
	wait_for_ack(serv_conn, packet_num, i)
	fmt.Println("after waiting for ack", <-i)
	select {
	case <-i:
	fmt.Println(<-i, "I'm in a select!")
		break
	case <-time.After(5 * time.Second):
		fmt.Println("timed out for", packet_num)
		actions(conn, serv_conn)
	}
	packet_num := <-i
	packet_num++
	fmt.Println("Next to send:", i)
	if err != nil {
		fmt.Println(msg, err)
	}
	time.Sleep(time.Second * 1)

}
