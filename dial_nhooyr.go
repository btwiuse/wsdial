//go:build !gorilla

package wsdial

import (
	"context"
	"encoding/base64"
	"net"
	"net/http"
	"net/url"

	"github.com/coder/websocket"
)

func Dial(u *url.URL) (conn net.Conn, err error) {
	hdr := http.Header{}
	if u.User != nil {
		hdr.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(u.User.String())))
	}

	wsconn, _, err := websocket.Dial(
		context.Background(),
		u.String(),
		dialOptions(hdr),
	)
	if err != nil {
		return nil, err
	}

	return websocket.NetConn(context.Background(), wsconn, websocket.MessageBinary), nil
}

