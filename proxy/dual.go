package proxy

import (
	"context"
	"net"

	M "github.com/xjasonlyu/tun2socks/v2/metadata"
)

type Dual struct {
	tcp Proxy
	udp Proxy
}

func NewDual(tcp, udp Proxy) *Dual {
	return &Dual{tcp: tcp, udp: udp}
}

func (d *Dual) DialContext(ctx context.Context, metadata *M.Metadata) (net.Conn, error) {
	return d.tcp.DialContext(ctx, metadata)
}

func (d *Dual) DialUDP(metadata *M.Metadata) (net.PacketConn, error) {
	return d.udp.DialUDP(metadata)
}
