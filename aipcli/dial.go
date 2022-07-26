package aipcli

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
	if config.Runtime.Insecure {
		return dialInsecure(ctx, config)
	}
	address, ok := config.GetAddress()
	if !ok {
		return nil, fmt.Errorf("dial: no address")
	}
	Logf(ctx, "address: %s", address)
	var opts []grpc.DialOption
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

func dialInsecure(ctx context.Context, config *Config) (*grpc.ClientConn, error) {
	token, hasToken := config.GetToken()
	switch {
	case config.Runtime.Address == "":
		return nil, fmt.Errorf("must provide --address with --insecure")
	case hasToken && !strings.HasPrefix(config.Runtime.Address, "localhost:"):
		return nil, fmt.Errorf("must connect to localhost with --insecure and --token")
	}
	Logf(ctx, "insecure address: %s", config.Runtime.Address)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if token != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(insecureTokenCredentials(token)))
	}
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
