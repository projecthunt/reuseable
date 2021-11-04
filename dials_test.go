package reuseable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDial_On_SamePort_SameIP(t *testing.T) {
	host := "127.0.0.1"

	l1, err := Listen("tcp", host+":0")
	t.Logf("Listener 1 Error: %v", err)
	require.NoError(t, err)

	l2, err := Listen("tcp", host+":0")
	t.Logf("Listener 2 Error: %v", err)
	require.NoError(t, err)

	conn1, err := Dial("tcp", l1.Addr().String(), l2.Addr().String())
	t.Logf("Dial 1 Error: %v", err)
	require.NoError(t, err)

	l1.Close()
	l2.Close()
	conn1.Close()
}

func TestDialTimeout_On_SamePort_SameIP(t *testing.T) {
	host := "127.0.0.1"

	l1, err := Listen("tcp", host+":0")
	t.Logf("Listener 1 Error: %v", err)
	require.NoError(t, err)

	l2, err := Listen("tcp", host+":0")
	t.Logf("Listener 2 Error: %v", err)
	require.NoError(t, err)

	conn1, err := DialTimeout("tcp", l1.Addr().String(), l2.Addr().String(), time.Second*5)
	t.Logf("Dial 1 Error: %v", err)
	require.NoError(t, err)

	l1.Close()
	l2.Close()
	conn1.Close()
}

func TestResolver(t *testing.T) {
	tests := []map[string]interface{}{
		{
			"network":   "ip",
			"address":   "127.0.0.1",
			"expectErr": false,
		},
		{
			"network":   "ip4",
			"address":   "127.0.0.1",
			"expectErr": false,
		},
		{
			"network":   "ip6",
			"address":   "0:0:0:0:0:0:0:0",
			"expectErr": false,
		},
		{
			"network":   "tcp",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "tcp4",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "tcp6",
			"address":   "[0:0:0:0:0:0:0:0]:80",
			"expectErr": false,
		},
		{
			"network":   "udp",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "udp4",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "udp6",
			"address":   "[0:0:0:0:0:0:0:0]:80",
			"expectErr": false,
		},
		{
			"network":   "unix",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "unixgram",
			"address":   "127.0.0.1:80",
			"expectErr": false,
		},
		{
			"network":   "unixpacket",
			"address":   "[0:0:0:0:0:0:0:0]:80",
			"expectErr": false,
		},
		{
			"network":   "unknown",
			"address":   "trying",
			"expectErr": true,
		},
	}

	for _, test := range tests {
		network := test["network"].(string)
		address := test["address"].(string)
		isErrExpected := test["expectErr"].(bool)

		_, err := resolver(network, address)
		if !isErrExpected {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
	}
}
