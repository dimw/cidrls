package cidr_test

import (
	"github.com/dimw/cidrls/cidr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostsWith32BitMask(t *testing.T) {
	hosts, _ := cidr.Hosts("127.0.0.1/32")

	assert.Equal(t, len(hosts), 1)
	assert.Equal(t, hosts[0], "127.0.0.1")
}

func TestHostsWith28BitMask(t *testing.T) {
	hosts, _ := cidr.Hosts("127.0.0.1/30")

	assert.Equal(t, len(hosts), 4)
	assert.Equal(t, hosts[0], "127.0.0.0")
	assert.Equal(t, hosts[1], "127.0.0.1")
	assert.Equal(t, hosts[2], "127.0.0.2")
	assert.Equal(t, hosts[3], "127.0.0.3")
}

func TestHostsWithoutMask(t *testing.T) {
	hosts, _ := cidr.Hosts("127.0.0.1")

	assert.Equal(t, len(hosts), 1)
	assert.Equal(t, hosts[0], "127.0.0.1")
}

func TestCalculateIpsWithNetworkAndBroadcast(t *testing.T) {
	ips := cidr.CalculateIps("127.0.0.1/30", false)

	assert.Equal(t, len(ips), 4)
	assert.Equal(t, ips[0], "127.0.0.0")
	assert.Equal(t, ips[1], "127.0.0.1")
	assert.Equal(t, ips[2], "127.0.0.2")
	assert.Equal(t, ips[3], "127.0.0.3")
}

func TestCalculateIpsWithoutNetworkAndBroadcast(t *testing.T) {
	ips := cidr.CalculateIps("127.0.0.1/30", true)

	assert.Equal(t, len(ips), 2)
	assert.Equal(t, ips[0], "127.0.0.1")
	assert.Equal(t, ips[1], "127.0.0.2")
}
