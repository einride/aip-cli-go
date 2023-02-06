package aipcli

import (
	"context"
	"crypto/x509"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func dial(cmd *cobra.Command) (*grpc.ClientConn, error) {
	if isInsecure(cmd) {
		return dialInsecure(cmd)
	}
	address, ok := getAddress(cmd)
	if !ok {
		return nil, fmt.Errorf("dial: no address")
	}
	if IsVerbose(cmd) {
		cmd.PrintErrln(">> address:", address)
	}
	var opts []grpc.DialOption
	if token, ok := getToken(cmd); ok {
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
	return grpc.DialContext(cmd.Context(), withDefaultPort(address, 443), opts...)
}

func dialInsecure(cmd *cobra.Command) (*grpc.ClientConn, error) {
	token, hasToken := getToken(cmd)
	address, hasAddress := getAddress(cmd)
	switch {
	case !hasAddress:
		return nil, fmt.Errorf("must provide --address with --insecure")
	case hasToken && !strings.HasPrefix(address, "localhost:"):
		return nil, fmt.Errorf("must connect to localhost with --insecure and --token")
	}
	if IsVerbose(cmd) {
		cmd.PrintErrln(">> insecure address:", address)
	}
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if token != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(insecureTokenCredentials(token)))
	}
	return grpc.DialContext(cmd.Context(), withDefaultPort(address, 443), opts...)
}

func methodURI(method protoreflect.MethodDescriptor) string {
	return "/" +
		string(method.Parent().(protoreflect.ServiceDescriptor).FullName()) +
		"/" + string(method.Name())
}

type insecureTokenCredentials string

func (c insecureTokenCredentials) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + string(c),
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
