package control

import (
	"bazil.org/bazil/server/control/wire"
	"golang.org/x/net/context"
)

func (c controlRPC) Ping(ctx context.Context, req *wire.PingRequest) (*wire.PingResponse, error) {
	return &wire.PingResponse{}, nil
}
