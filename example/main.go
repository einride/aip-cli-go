package main

import (
	"log"

	"github.com/einride/ctl/example/ctl2"
)

// var localhost bool

func main() {
	if err := ctl2.Command.Execute(); err != nil {
		log.Fatal(err)
	}
}

//func connect(ctx context.Context) (*grpc.ClientConn, error) {
//	if localhost {
//		addr := "localhost:8080"
//		opts := []option.ClientOption{
//			option.WithEndpoint(addr),
//		}
//		token, err := ctl.RootFlags().GetString("token")
//		if err != nil {
//			return nil, err
//		}
//		if token != "" {
//			opts = append(opts, option.WithGRPCDialOption(grpc.WithPerRPCCredentials(tokenCredentials(token))))
//		}
//		return transport.DialGRPCInsecure(
//			ctx,
//			opts...,
//		)
//	}
//	return nil, nil
//}
//
//type tokenCredentials string
//
//func (t tokenCredentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
//	return map[string]string{
//		"authorization": "bearer " + string(t),
//	}, nil
//}
//
//func (t tokenCredentials) RequireTransportSecurity() bool {
//	return false
//}
