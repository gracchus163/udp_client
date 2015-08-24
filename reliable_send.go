package main

import (
	"fmt"
	"net"
	"strconv"
)

func wait_for_ack(serv_conn *net.UDPConn, expected int, i chan int) {
	fmt.Println("Waiting for ACK for", expected)

	buffer := make([]byte, 1024)
	ack_val := -1
	n, addr, err := serv_conn.ReadFromUDP(buffer)
	var curr_packet int
	if addr != nil {
		test, err := strconv.Atoi(string(buffer[0:n]))
		if err != nil {
			fmt.Println("Error: ", err)
		}
		curr_packet = test

		fmt.Println("Received", curr_packet, ", looking for", expected)
		ack_val = check_packet(curr_packet, expected)
		fmt.Println("Checked packet with code", ack_val)
		fmt.Println("Received ACK for", curr_packet, "from", addr)
		fmt.Println(ack_val)
	}
	if err != nil {
		fmt.Println("Error: ", err)
	}
	if ack_val != -1 {
		fmt.Println("writing", curr_packet, "to inbound channel")
		i <- curr_packet
		return
	}

	go wait_for_ack(serv_conn, expected, i)
}

func check_packet(curr_packet int, expected int) int {
	fmt.Println("curr_packet:", curr_packet, " waiting for ACK for:", expected)
	if curr_packet == expected {
		return curr_packet
	} else {
		return curr_packet-1
	}
}