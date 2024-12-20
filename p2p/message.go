package p2p

import "net"

// Message respesents any artitrary data that is being sent
// between two nodes in the network
type Message struct {
	From    net.Addr
	Payload []byte
}
