package main

import (
	"github/grimhat8346/dfs/p2p"
	"log"
)

func main() {

	TcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	tr := p2p.NewTCPTransport(TcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}

}
