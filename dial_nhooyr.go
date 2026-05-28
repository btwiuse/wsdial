//go:build !gorilla

package wsdial

import (
	"context"
	"net"
	"net/http"
	"net/url"

	"github.com/coder/websocket"
)

func Dial(ctx context.Context, u *url.URL, hdr http.Header) (conn net.Conn, err error) {
	hdr = fillAuthHeader(hdr, u)

	wsconn, _, err := websocket.Dial(
		ctx,
		u.String(),
		dialOptions(hdr),
	)
	if err != nil {
		return nil, err
	}

	wsconn.SetReadLimit(-1)

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}

