package ctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

var _ = fmt.Sprintf
var _ = cobra.Command{}
var einride_authz_v1beta1_AuthorizationService = &cobra.Command{
	Use: "einride.authz.v1beta1.Authorization",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.authz.v1beta1.Authorization called")
	},
}

func init() {
	rootCmd.AddCommand(einride_authz_v1beta1_AuthorizationService)
}
