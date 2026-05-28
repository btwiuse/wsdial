package wsdial

import (
	"encoding/base64"
	"net/http"
	"net/url"
)

func fillAuthHeader(hdr http.Header, u *url.URL) http.Header {
	if hdr == nil {
		hdr = http.Header{}
	}
	if u.User != nil && hdr.Get("Authorization") == "" {
		hdr.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(u.User.String())))
	}
	return hdr
}
