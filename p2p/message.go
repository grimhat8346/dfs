package p2p

import "net"

// RPC respesents any artitrary data that is being sent
// between two nodes in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
