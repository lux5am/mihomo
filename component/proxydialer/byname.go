package proxydialer

import (
	"context"
	"fmt"
	"net"
	"net/netip"

	C "github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/tunnel"
)

func findProxyByName(name string) C.Proxy {
	if proxy, exists := tunnel.Proxies()[name]; exists {
		return proxy
	}
	for _, p := range tunnel.Providers() {
		if proxy := p.ProxyByName(name); proxy != nil {
			return proxy
		}
	}
	return nil
}

type byNameProxyDialer struct {
	proxyName string
}

func (d byNameProxyDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	proxy := findProxyByName(d.proxyName)
	if proxy == nil {
		return nil, fmt.Errorf("proxyName[%s] not found", d.proxyName)
	}
	return New(proxy, true).DialContext(ctx, network, address)
}

func (d byNameProxyDialer) ListenPacket(ctx context.Context, network, address string, rAddrPort netip.AddrPort) (net.PacketConn, error) {
	proxy := findProxyByName(d.proxyName)
	if proxy == nil {
		return nil, fmt.Errorf("proxyName[%s] not found", d.proxyName)
	}
	return New(proxy, true).ListenPacket(ctx, network, address, rAddrPort)
}

func NewByName(proxyName string) C.Dialer {
	return byNameProxyDialer{proxyName: proxyName}
}
