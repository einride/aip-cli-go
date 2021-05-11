package plugin

import (
	"github.com/einride/ctl/internal/codegen"
)

func GenerateConnectFile(f *codegen.File) {
	f.P(`
package ctl

import (
	"fmt"
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
	fmt.Println(addr, token)

	return transport.DialGRPC(
		ctx,
		option.WithEndpoint(addr),
		option.WithTokenSource(
			oauth2.StaticTokenSource(
				&oauth2.Token{
					TokenType:   "Bearer",
					AccessToken: token,
				},
			),
		),
	)
}
`)
}
