package accountv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_account_v1beta1_AuthenticationService = &cobra.Command{
	Use: "einride.account.v1beta1.AuthenticationService",
}

var einride_account_v1beta1_AuthenticateUserIdentityTokenRequest v1beta1.AuthenticateUserIdentityTokenRequest
var einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken = &cobra.Command{
	Use: "AuthenticateUserIdentityToken",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.AuthenticationService.AuthenticateUserIdentityToken")
		return nil
	},
}

func AddAuthenticationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_AuthenticationService)
}

func init() {
	einride_account_v1beta1_AuthenticationService.AddCommand(einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken)
}
