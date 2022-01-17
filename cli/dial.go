package cli

import (
	"context"
	"crypto/x509"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

func Dial(ctx context.Context) (*grpc.ClientConn, error) {
	config := ConfigFromContext(ctx)
	opts := []grpc.DialOption{grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.Runtime.MaxRecvSize))}
	if config.Runtime.Insecure {
		return dialInsecure(ctx, config, opts)
	}
	address, ok := config.GetAddress()
	if !ok {
		return nil, fmt.Errorf("dial: no address")
	}
	Log(ctx, "address: %s", address)
	if token, ok := config.GetToken(); ok {
		opts = append(
			opts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(
					&oauth2.Token{
						AccessToken: token,
						TokenType:   "Bearer",
					},
				),
			),
		)
	}
	systemCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(systemCertPool, "")))
	return grpc.DialContext(ctx, withDefaultPort(address, 443), opts...)
}

func dialInsecure(ctx context.Context, config *Config, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	token, hasToken := config.GetToken()
	switch {
	case config.Runtime.Address == "":
		return nil, fmt.Errorf("must provide --address with --insecure")
	case hasToken && !strings.HasPrefix(config.Runtime.Address, "localhost:"):
		return nil, fmt.Errorf("must connect to localhost with --insecure and --token")
	}
	Log(ctx, "insecure address: %s", config.Runtime.Address)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if token != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(insecureTokenCredentials(token)))
	}
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(config.Runtime.MaxRecvSize)))
	return grpc.DialContext(ctx, withDefaultPort(config.Runtime.Address, 443), opts...)
}

type insecureTokenCredentials string

func (c insecureTokenCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "bearer " + string(c),
	}, nil
}

func (insecureTokenCredentials) RequireTransportSecurity() bool {
	return false
}

func withDefaultPort(target string, port int) string {
	parts := strings.Split(target, ":")
	if len(parts) == 1 {
		return target + ":" + strconv.Itoa(port)
	}
	return target
}
