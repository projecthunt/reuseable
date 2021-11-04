package reuseable

import (
	"net"
	"time"
)

var netDialer = net.Dialer{
	Control: control,
}

func resolver(network, address string) (net.Addr, error) {
	switch network {
	default:
		return nil, net.UnknownNetworkError(network)
	case "ip", "ip4", "ip6":
		return net.ResolveIPAddr(network, address)
	case "tcp", "tcp4", "tcp6":
		return net.ResolveTCPAddr(network, address)
	case "udp", "udp4", "udp6":
		return net.ResolveUDPAddr(network, address)
	case "unix", "unixgram", "unixpacket":
		return net.ResolveUnixAddr(network, address)
	}
}

// Dial do the same thing with net.Dial but with socket that have SO_REUSEADDR and SO_REUSEPORT flags.
func Dial(network string, laddr string, raddr string) (net.Conn, error) {
	netDialerWithLocalAddr := netDialer
	localAddr, err := resolver(network, laddr)
	if err != nil {
		return nil, err
	}
	netDialerWithLocalAddr.LocalAddr = localAddr
	return netDialerWithLocalAddr.Dial(network, raddr)
}

// DialTimeout do the same thing with net.DialTimeout but with socket that have SO_REUSEADDR and SO_REUSEPORT flags.
func DialTimeout(network string, laddr string, raddr string, timeout time.Duration) (net.Conn, error) {
	netDialerWithTimeout := netDialer
	netDialerWithTimeout.Timeout = timeout
	localAddr, err := resolver(network, laddr)
	if err != nil {
		return nil, err
	}
	netDialerWithTimeout.LocalAddr = localAddr
	return netDialerWithTimeout.Dial(network, raddr)
}
