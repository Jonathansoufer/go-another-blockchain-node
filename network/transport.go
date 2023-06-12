package network

type NetAddr string

type RPC struct {
	From	NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMEssage(NetAddr, []byte) error
	Addr() NetAddr
}