package proxy

import (
	"context"
	"encoding/json"

	"github.com/centrifugal/centrifuge"
)

// ConnectRequest ...
type ConnectRequest struct {
	ClientID  string
	Transport centrifuge.Transport
	Data      centrifuge.Raw
}

// ConnectCredentials ...
type ConnectCredentials struct {
	UserID     string          `json:"user"`
	ExpireAt   int64           `json:"exp"`
	Info       json.RawMessage `json:"info"`
	Base64Info string          `json:"b64info"`
}

// ConnectResult ...
type ConnectResult struct {
	Credentials *ConnectCredentials    `json:"credentials"`
	Error       *centrifuge.Error      `json:"error"`
	Disconnect  *centrifuge.Disconnect `json:"disconnect"`
}

// ConnectProxy allows to proxy connect requests to application backend to
// authenticate client connection.
type ConnectProxy interface {
	ProxyConnect(context.Context, ConnectRequest) (*ConnectResult, error)
}