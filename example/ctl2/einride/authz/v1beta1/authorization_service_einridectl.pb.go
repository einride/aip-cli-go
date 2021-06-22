package authzv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/authz/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_authz_v1beta1_AuthorizationService = &cobra.Command{
	Use: "einride.authz.v1beta1.AuthorizationService",
}

var einride_authz_v1beta1_GetAuthorizationPolicyRequest v1beta1.GetAuthorizationPolicyRequest
var einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy = &cobra.Command{
	Use: "GetAuthorizationPolicy",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.GetAuthorizationPolicy")
		return nil
	},
}

var einride_authz_v1beta1_SearchRoleBindingsRequest v1beta1.SearchRoleBindingsRequest
var einride_authz_v1beta1_AuthorizationService_SearchRoleBindings = &cobra.Command{
	Use: "SearchRoleBindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.SearchRoleBindings")
		return nil
	},
}

var einride_authz_v1beta1_AddRoleBindingRequest v1beta1.AddRoleBindingRequest
var einride_authz_v1beta1_AuthorizationService_AddRoleBinding = &cobra.Command{
	Use: "AddRoleBinding",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.AddRoleBinding")
		return nil
	},
}

var einride_authz_v1beta1_RemoveRoleBindingRequest v1beta1.RemoveRoleBindingRequest
var einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding = &cobra.Command{
	Use: "RemoveRoleBinding",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.RemoveRoleBinding")
		return nil
	},
}

var einride_authz_v1beta1_LookupEffectivePermissionsRequest v1beta1.LookupEffectivePermissionsRequest
var einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions = &cobra.Command{
	Use: "LookupEffectivePermissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.LookupEffectivePermissions")
		return nil
	},
}

var einride_authz_v1beta1_BatchLookupEffectivePermissionsRequest v1beta1.BatchLookupEffectivePermissionsRequest
var einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions = &cobra.Command{
	Use: "BatchLookupEffectivePermissions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.BatchLookupEffectivePermissions")
		return nil
	},
}

var einride_authz_v1beta1_LookupPermissionBindingsRequest v1beta1.LookupPermissionBindingsRequest
var einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings = &cobra.Command{
	Use: "LookupPermissionBindings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.authz.v1beta1.AuthorizationService.LookupPermissionBindings")
		return nil
	},
}

func AddAuthorizationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_authz_v1beta1_AuthorizationService)
}

func init() {
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SearchRoleBindings)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_AddRoleBinding)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings)
}
