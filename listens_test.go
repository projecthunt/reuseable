package reuseable

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListen_On_SamePort_SameIP(t *testing.T) {
	host := "127.0.0.1"

	// Listen on an automatically chosen port
	l1, err := Listen("tcp", host+":0")
	t.Logf("Listener 1 Error: %v", err)
	require.NoError(t, err)

	chosenPort := strings.Split(l1.Addr().String(), ":")[1]

	// Listen on same port and ip address
	l2, err := Listen("tcp", host+":"+chosenPort)
	t.Logf("Listener 2 Error: %v", err)
	require.NoError(t, err)

	t.Logf("Listener 1: %s", l1.Addr().String())
	t.Logf("Listener 2: %s", l2.Addr().String())

	l1.Close()
	l2.Close()
}

func TestListenPacket_On_SamePort_SameIP(t *testing.T) {
	host := "127.0.0.1"

	// Listen on an automatically chosen port
	l1, err := ListenPacket("udp", host+":0")
	t.Logf("Listener 1 Error: %v", err)
	require.NoError(t, err)

	chosenPort := strings.Split(l1.LocalAddr().String(), ":")[1]

	// Listen on same port and ip address
	l2, err := ListenPacket("udp", host+":"+chosenPort)
	t.Logf("Listener 2 Error: %v", err)
	require.NoError(t, err)

	t.Logf("Listener 1: %s", l1.LocalAddr().String())
	t.Logf("Listener 2: %s", l2.LocalAddr().String())

	l1.Close()
	l2.Close()
}
