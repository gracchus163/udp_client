package main

import (
	"fmt"
	"net"
	"strconv"
)

func wait_for_ack(expected int) int {
	fmt.Println("Waiting for ACK for", expected)
	serv_addr, err := net.ResolveUDPAddr("udp", ":10002")
	CheckError(err)

	serv_conn, err := net.ListenUDP("udp", serv_addr)
	CheckError(err)
	defer serv_conn.Close()

	buffer := make([]byte, 1024)
	ack_val := -1
	n, addr, err := serv_conn.ReadFromUDP(buffer)

	var curr_packet int
	test, err := strconv.Atoi(string(buffer[0:n]))
	curr_packet = test

	fmt.Println("Received", curr_packet, ", looking for", expected)
	ack_val = check_packet(curr_packet, expected)
	fmt.Println("Checked packet with code", ack_val)
	fmt.Println("Received ACK for", curr_packet, "from", addr)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	if ack_val != -1 {
		return expected
	}
	return expected-1
}

func check_packet(curr_packet int, expected int) int {
	fmt.Println("curr_packet:", curr_packet, " waiting for ACK for:", expected)
	if curr_packet == expected {
		return curr_packet
	} else {
		return curr_packet-1
	}
}