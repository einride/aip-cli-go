package plugin

import (
	"github.com/einride/ctl/internal/codegen"
)

func GenerateConnectFile(f *codegen.File) {
	f.P(`
import (
    "context"

	"google.golang.org/grpc"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func connect(ctx context.Context) (*grpc.ClientConn, error) {
	addr := devHost
	if prod {
      addr = prodHost
    }
	if address != "" {
	  addr = address
	}

	opts := []option.ClientOption{
	  option.WithEndpoint(addr),
	}

	if insecure {
	  if address == "" {
	    return nil, fmt.Errorf("cannot use insecure connection without specifying address")
	  }
	  return transport.DialGRPCInsecure(
		ctx,
		opts...,
	  )
	}
	opts = append(opts, option.WithTokenSource(
	  oauth2.StaticTokenSource(
		  &oauth2.Token{
			  TokenType:   "Bearer",
			  AccessToken: token,
		  },
	  ),
	),
	)
	return transport.DialGRPC(
		ctx,
		opts...,
	)
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
`)
}
