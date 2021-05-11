package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_account_v1beta1_AuthenticationService = &cobra.Command{
	Use: "einride.account.v1beta1.Authentication",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.Authentication called")
	},
}
var einride_account_v1beta1_FeatureFlagsService = &cobra.Command{
	Use: "einride.account.v1beta1.FeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.FeatureFlags called")
	},
}
var einride_account_v1beta1_TenantService = &cobra.Command{
	Use: "einride.account.v1beta1.Tenant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.Tenant called")
	},
}
var einride_account_v1beta1_UserService = &cobra.Command{
	Use: "einride.account.v1beta1.User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.User called")
	},
}

func init() {
	rootCmd.AddCommand(einride_account_v1beta1_AuthenticationService)
	rootCmd.AddCommand(einride_account_v1beta1_FeatureFlagsService)
	rootCmd.AddCommand(einride_account_v1beta1_TenantService)
	rootCmd.AddCommand(einride_account_v1beta1_UserService)
}
