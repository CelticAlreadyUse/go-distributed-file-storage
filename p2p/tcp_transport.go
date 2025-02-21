package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddr string
	listenner  net.Interface
	mu sync.RWMutex
	peers  map[net.Addr]Peer
}

func InitCPTransport(listenAddr string)*TCPTransport{
	return &TCPTransport{
		listenAddr: listenAddr,
	}
}