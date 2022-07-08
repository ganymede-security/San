package demux

import (
	"flag"
	"fmt"
	"log"
	"net"
	"syscall"
)

type Driver struct {
	ListenPort	string
	ReadBuffer string
	WriteBuffer string
}

// Constants used by Driver when none are specified
const (
	DefaultWriteBuffer = 4096
	DefaultReadBuffer = 4096
)

var (
	addr = flag.String("s", "localhost", "server address")
	port = flag.Int("p", 8972, "port")
)

func Listen() {

	conn, err := net.ListenPacket("udp", *addr)
	if err != nil {
		log.Fatal(err)
	}
	cc := conn.(*net.IPConn)
	cc.SetReadBuffer(20 )

}