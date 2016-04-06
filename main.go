package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

// Calculate the TCP/IP checksum defined in rfc1071.
func genChecksum(data []byte, csum uint32) uint16 {
	length := len(data) - 1
	for i := 0; i < length; i += 2 {
		csum += uint32(data[i]) << 8
		csum += uint32(data[i+1])
	}
	if len(data)%2 == 1 {
		csum += uint32(data[length]) << 8
	}
	for csum > 0xffff {
		csum = (csum & 0xffff) + (csum >> 16)
	}
	return ^uint16((csum >> 16) + csum)
}

func main() {
	var payload []byte

	timeoutDuration, _ := time.ParseDuration("10s")
	grpAddress := net.ParseIP("0.0.0.0") // TODO: Make arg
	dstAddress := net.ParseIP("224.0.0.1")

	// IGMP https://tools.ietf.org/html/rfc2236#section-2
	payload = make([]byte, 8, 8)
	payload[0] = uint8(0x11)
	payload[1] = uint8(100) // TODO: Make arg
	payload[4] = grpAddress.To4()[0]
	payload[5] = grpAddress.To4()[1]
	payload[6] = grpAddress.To4()[2]
	payload[7] = grpAddress.To4()[3]
	binary.BigEndian.PutUint16(payload[2:], genChecksum(payload, 0))

	tickC := time.NewTicker(time.Second * 5).C // TODO: make arg
	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, os.Interrupt)
	go func() {
		// for t := range tickC {
		for _ = range tickC {
			// fmt.Println("Tick at", t)
			// Send packet
			conn, err := net.DialTimeout("ip:igmp", dstAddress.String(), timeoutDuration)
			_, err = conn.Write(payload)
			if err != nil {
				fmt.Println("Error occured. ", err)
			}
		}
	}()
	<-signalC
}
