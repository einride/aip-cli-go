package plugin

import (
	"github.com/einride/ctl/internal/codegen"
)

func GenerateConnectFile(f *codegen.File) {
	f.P(`
package ctl

import (
	"google.golang.org/grpc"
)

func connect() (*grpc.ClientConn, error) {
	const addr = "api-g4oz7jceaa-ew.a.run.app:443"
	conn, err := grpc.Dial(addr)
	return conn, err
}
`)
}
