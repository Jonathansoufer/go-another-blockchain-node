package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr NetAddr
	consumeCh chan RPC
	lock sync.Mutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr: addr,
		consumeCh: make(chan RPC),
		peers: make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumeCh
}

func (t *LocalTransport) Connect(tr *LocalTransport) error {
	trans := tr
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = trans

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.addr == to {
		return fmt.Errorf("cannot send message to self")
	}

	if peer, ok := t.peers[to]; !ok {
		return fmt.Errorf("peer %v not found", to)
	} else {
		peer.consumeCh <- RPC{t.addr, payload}
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}