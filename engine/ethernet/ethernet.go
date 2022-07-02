package ethernet

import (
	"context"
	"net"

	"github.com/ganymede-security/san/engine"
)

// EtherType values for commonly used protocols
const (
	// IP Version 4
	IPv4 = ^uint16(0x0800)
	// Address Resolution Protocol
	ARP = ^uint16(0x0806)
	// Reverse Address Resolution Protocol
	RARP = ^uint16(0x8035)
	// VLAN 802.1q Trunking Protocol Tag
	VLAN = ^uint16(0x8100)
	// VLAN 802.1ad VLAN Service Tag
	VLANSrv = ^uint16(0x88A8)
	// IP Version 6
	IPv6 = ^uint16(0x86DD)
	// Link Layer Discovery Protocol
	LLDP = ^uint16(0x88CC)
	// MAC Security
	MACsec = ^uint16(0x88E5)
	// Precision Time Protocol
	PTP = ^uint16(0x88F7)
	// Ethernet Config Testing Protocol
	ECTP = ^uint16(0x9000)
)

// Struct representing an IEEE 802.3 Ethernet Frame
type EthernetFrame struct {
	// Contains the Preamble as well as the SFD (Start Frame Delimiter).
	//
	// The SFD is a 1byte field that is always set to 1010111.
	Preamble []byte
	// The Destination MAC address
	Destination net.HardwareAddr
	// The Source MAC address
	Source net.HardwareAddr
	// The optional 802.1Q tag used for VLAN trunking/ QoS
	TrunkTag []byte
	// Two byte field used to indicate the encapsulated protocol
	//
	// if =< 1500 bytes it indicates Payload length
	//
	// if >= 1536 bytes it indicates an EtherType
	//
	// Common EtherTypes are labeled as Const values
	Payload []byte
	// Frame Check is 4 bytes, it allows us to detect if information
	// in the payload was corrupted.
	FrameCheck []byte
}

func EthDriver() *engine.Driver {
	drv := "test"
	return &engine.Driver{
		Handler: drv,
	}
}

func EthernetEngine() *engine.Engine {
	drv := EthDriver()
	eng:= engine.NewEngine(drv)
	return eng
}