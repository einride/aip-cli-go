package accountv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.account.v1beta1.AuthenticationService.
var (
	einride_account_v1beta1_AuthenticationServiceClient v1beta1.AuthenticationServiceClient
	einride_account_v1beta1_AuthenticationService       = &cobra.Command{
		Use: "einride.account.v1beta1.AuthenticationService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_account_v1beta1_AuthenticationServiceClient = v1beta1.NewAuthenticationServiceClient(conn)
			return nil
		},
	}
)

// einride.account.v1beta1.AuthenticationService.AuthenticateUserIdentityToken.
var (
	einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken_Request v1beta1.AuthenticateUserIdentityTokenRequest
	einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken         = &cobra.Command{
		Use: "AuthenticateUserIdentityToken",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.AuthenticationService.AuthenticateUserIdentityToken")
			return nil
		},
	}
)

func AddAuthenticationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_AuthenticationService)
}

func init() {
	einride_account_v1beta1_AuthenticationService.AddCommand(einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken)

	einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken.Flags().StringVar(&einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken_Request.IdentityToken, "identityToken", "", "The raw user identity token to authenticate.")
}
