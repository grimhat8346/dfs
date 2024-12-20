package p2p

// Peer represent the remote node and connection
type Peer interface {
}

// Transport is everithing that hadles communications between the nodes in the network
// In form TCP, UDP, WebSocket
type Transport interface {
	ListenAndAccept() error
}
