package main

import (
	"time"

	"github.com/Jonathansoufer/go-another-blockchain-node/network"
)

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func(){
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{},
	}

	s := network.NewServer(opts)
	s.Start()
}