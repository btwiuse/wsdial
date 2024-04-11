//go:build gorilla && !nhooyr

package wsdial

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

func Dial(u *url.URL) (conn net.Conn, err error) {
	var wd *websocket.Dialer = &websocket.Dialer{
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig: &tls.Config{
			RootCAs:            nil,
			InsecureSkipVerify: true,
		},
	}

	wsconn, _, err := wd.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}

	return wsconn.NetConn(wsconn), nil
}
