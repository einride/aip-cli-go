package main

import (
	"context"

	"google.golang.org/api/transport"

	"github.com/einride/ctl/example/ctl"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

var localhost bool

func main() {
	ctl.CustomConnect = connect
	ctl.RootFlags().BoolVar(
		&localhost,
		"localhost",
		false,
		"make insecure request to localhost, will also set token if that is provided",
	)
	ctl.Execute()
}

func connect(ctx context.Context) (*grpc.ClientConn, error) {
	if localhost {
		addr := "localhost:8080"
		opts := []option.ClientOption{
			option.WithEndpoint(addr),
		}
		token, err := ctl.RootFlags().GetString("token")
		if err != nil {
			return nil, err
		}
		if token != "" {
			opts = append(opts, option.WithGRPCDialOption(grpc.WithPerRPCCredentials(tokenCredentials(token))))
		}
		return transport.DialGRPCInsecure(
			ctx,
			opts...,
		)
	}
	return nil, nil
}

type tokenCredentials string

func (t tokenCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + string(t),
	}, nil
}

func (t tokenCredentials) RequireTransportSecurity() bool {
	return false
}
