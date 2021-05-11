package plugin

import (
	"github.com/einride/ctl/internal/codegen"
)

func GenerateConnectFile(f *codegen.File) {
	f.P(`
import (
    "context"
	"fmt"

	"google.golang.org/grpc"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/transport"
)

func connect(ctx context.Context) (*grpc.ClientConn, error) {
	if CustomConnect != nil {
	  customConnection, err := CustomConnect(ctx)
	  if err != nil {
		return nil, err
      }
	  if customConnection != nil {
		return customConnection, nil
	  }
	}
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
`)
}
