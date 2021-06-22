package ctl

import (
	"context"
	"crypto/x509"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type DialConfig struct {
	Insecure bool
	Address  string
	Token    string
}

func ParseDialConfig(flags *pflag.FlagSet) (DialConfig, error) {
	insecure, err := flags.GetBool("insecure")
	if err != nil {
		return DialConfig{}, err
	}
	address, err := flags.GetString("address")
	if err != nil {
		return DialConfig{}, err
	}
	token, err := flags.GetString("token")
	if err != nil {
		return DialConfig{}, err
	}
	return DialConfig{
		Insecure: insecure,
		Address:  address,
		Token:    token,
	}, nil
}

func Dial(ctx context.Context, config DialConfig) (*grpc.ClientConn, error) {
	if config.Insecure {
		return dialInsecure(ctx, config)
	}
	var opts []grpc.DialOption
	if config.Token != "" {
		opts = append(
			opts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(
					&oauth2.Token{
						AccessToken: config.Token,
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
	return grpc.DialContext(ctx, withDefaultPort(config.Address, 443), opts...)
}

func dialInsecure(ctx context.Context, config DialConfig) (*grpc.ClientConn, error) {
	switch {
	case config.Address == "":
		log.Fatal("must provide --address with --insecure")
	case config.Token != "" && !strings.HasPrefix(config.Address, "localhost:"):
		log.Fatal("must connect to localhost with --insecure and --token")
	}
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if config.Token != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(insecureTokenCredentials(config.Token)))
	}
	return grpc.DialContext(ctx, withDefaultPort(config.Address, 443), opts...)
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
