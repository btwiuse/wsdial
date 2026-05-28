//go:build js

package wsdial

import (
	"net/http"

	"github.com/coder/websocket"
)

func dialOptions(http.Header) *websocket.DialOptions {
	return &websocket.DialOptions{}
}
