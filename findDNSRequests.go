package main

import (
	"fmt"
	"sort"

	"github.com/CapitanShinChan/gopacket/layers"

	"github.com/CapitanShinChan/gopacket"
	"github.com/CapitanShinChan/gopacket/pcap"
)

var dnsNames []string

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

func handlePacket(pacote gopacket.Packet) {
	if appPayload := pacote.ApplicationLayer(); appPayload != nil {
		if dnsContent := pacote.Layer(layers.LayerTypeDNS); dnsContent != nil {
			// fmt.Println("This is a DNS packet")
			dnsInfo, _ := dnsContent.(*layers.DNS)
			if len(dnsInfo.Questions) > 0 {
				name := string(dnsInfo.Questions[0].Name)
				if !(include(dnsNames, name)) {
					dnsNames = append(dnsNames, name)
				}
			}
		}
	}
}

func main() {
	if handle, err := pcap.OpenOffline("D:\\18-07-23_R6-Siege\\pcaps\\02_online_gaming.pcap"); err != nil {
		panic(err)
	} else {
		handle.SetBPFFilter("udp")
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
		sort.Strings(dnsNames)
		for i := range dnsNames {
			fmt.Println(dnsNames[i])
		}
	}
}
