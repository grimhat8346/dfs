package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	//connection of the Peer
	conn net.Conn
	//if we dial and retrive a conn => outbound == true
	//if we dial and accept and retrive a conn => outbound == false
	outbound bool
}
type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()
	return nil

}

func (t *TCPTransport) startAcceptLoop() error {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("tcp listener accept error: %v\n", err)
		}
		fmt.Printf("tcp new connection from %+v\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := t.HandshakeFunc(peer); err != nil {
		fmt.Printf("tcp handshake error: %v\n", err)
		conn.Close()
		return
	}
	msg := &Message{}
	for {
		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp decoder error: %v\n", err)
			continue
		}

		fmt.Printf("tcp decoder msg: %+v/n", msg)
	}
}
