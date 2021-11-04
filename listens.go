package reuseable

import (
	"context"
	"net"
)

var listenConf = net.ListenConfig{
	Control: control,
}

// Listen do the same thing with net.Listen but with socket that have SO_REUSEADDR and SO_REUSEPORT flags.
func Listen(network string, address string) (net.Listener, error) {
	return listenConf.Listen(context.Background(), network, address)
}

// ListenPacket do the same thing with net.ListenPacket but with socket that have SO_REUSEADDR and SO_REUSEPORT flags.
func ListenPacket(network string, address string) (net.PacketConn, error) {
	return listenConf.ListenPacket(context.Background(), network, address)
}
