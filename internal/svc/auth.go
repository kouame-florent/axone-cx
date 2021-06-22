package svc

import (
	"context"
	"encoding/base64"
)

type BasicAuth struct {
	Login    string
	Password string
}

func (b BasicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := b.Login + ":" + b.Password
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{
		"authorization": "Basic " + enc,
	}, nil
}

func (b BasicAuth) RequireTransportSecurity() bool {
	return true
}
