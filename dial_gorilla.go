//go:build gorilla

package wsdial

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func Dial(ctx context.Context, u *url.URL, hdr http.Header) (conn net.Conn, err error) {
	hdr = fillAuthHeader(hdr, u)

	var wd *websocket.Dialer = &websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: true,
		},
	}

	wsconn, _, err := wd.DialContext(ctx, u.String(), hdr)
	if err != nil {
		return nil, err
	}

	return wsconn.NetConn(), nil
}
