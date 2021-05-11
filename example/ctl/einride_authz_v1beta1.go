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

var einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy = &cobra.Command{
	Use: "GetAuthorizationPolicy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetAuthorizationPolicy called")
	},
}

var einride_authz_v1beta1_AuthorizationService_SearchRoleBindings = &cobra.Command{
	Use: "SearchRoleBindings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SearchRoleBindings called")
	},
}

var einride_authz_v1beta1_AuthorizationService_AddRoleBinding = &cobra.Command{
	Use: "AddRoleBinding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AddRoleBinding called")
	},
}

var einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding = &cobra.Command{
	Use: "RemoveRoleBinding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RemoveRoleBinding called")
	},
}

var einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions = &cobra.Command{
	Use: "LookupEffectivePermissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LookupEffectivePermissions called")
	},
}

var einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions = &cobra.Command{
	Use: "BatchLookupEffectivePermissions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchLookupEffectivePermissions called")
	},
}

var einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings = &cobra.Command{
	Use: "LookupPermissionBindings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LookupPermissionBindings called")
	},
}

func init() {
	rootCmd.AddCommand(einride_authz_v1beta1_AuthorizationService)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SearchRoleBindings)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_AddRoleBinding)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings)
}
