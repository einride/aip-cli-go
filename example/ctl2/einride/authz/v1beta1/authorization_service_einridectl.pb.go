package authzv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/authz/v1beta1"
	cobra "github.com/spf13/cobra"
	v11 "google.golang.org/genproto/googleapis/iam/admin/v1"
	v1 "google.golang.org/genproto/googleapis/iam/v1"
	log "log"
)

// einride.authz.v1beta1.AuthorizationService.
var (
	einride_authz_v1beta1_AuthorizationServiceClient v1beta1.AuthorizationServiceClient
	einride_authz_v1beta1_AuthorizationService       = &cobra.Command{
		Use: "einride.authz.v1beta1.AuthorizationService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_authz_v1beta1_AuthorizationServiceClient = v1beta1.NewAuthorizationServiceClient(conn)
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.GetAuthorizationPolicy.
var (
	einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy_Request v1beta1.GetAuthorizationPolicyRequest
	einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy         = &cobra.Command{
		Use: "GetAuthorizationPolicy",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.GetAuthorizationPolicy")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.SearchRoleBindings.
var (
	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request v1beta1.SearchRoleBindingsRequest
	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings         = &cobra.Command{
		Use: "SearchRoleBindings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.SearchRoleBindings")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.AddRoleBinding.
var (
	einride_authz_v1beta1_AuthorizationService_AddRoleBinding_Request v1beta1.AddRoleBindingRequest
	einride_authz_v1beta1_AuthorizationService_AddRoleBinding         = &cobra.Command{
		Use: "AddRoleBinding",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.AddRoleBinding")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.RemoveRoleBinding.
var (
	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding_Request v1beta1.RemoveRoleBindingRequest
	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding         = &cobra.Command{
		Use: "RemoveRoleBinding",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.RemoveRoleBinding")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.LookupEffectivePermissions.
var (
	einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions_Request v1beta1.LookupEffectivePermissionsRequest
	einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions         = &cobra.Command{
		Use: "LookupEffectivePermissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.LookupEffectivePermissions")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.BatchLookupEffectivePermissions.
var (
	einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions_Request v1beta1.BatchLookupEffectivePermissionsRequest
	einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions         = &cobra.Command{
		Use: "BatchLookupEffectivePermissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.BatchLookupEffectivePermissions")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.LookupPermissionBindings.
var (
	einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings_Request v1beta1.LookupPermissionBindingsRequest
	einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings         = &cobra.Command{
		Use: "LookupPermissionBindings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.LookupPermissionBindings")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.SearchAuthorizedResources.
var (
	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request v1beta1.SearchAuthorizedResourcesRequest
	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources         = &cobra.Command{
		Use: "SearchAuthorizedResources",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.SearchAuthorizedResources")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.TestPermissions.
var (
	einride_authz_v1beta1_AuthorizationService_TestPermissions_Request v1beta1.TestPermissionsRequest
	einride_authz_v1beta1_AuthorizationService_TestPermissions         = &cobra.Command{
		Use: "TestPermissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.TestPermissions")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.SetIamPolicy.
var (
	einride_authz_v1beta1_AuthorizationService_SetIamPolicy_Request v1.SetIamPolicyRequest
	einride_authz_v1beta1_AuthorizationService_SetIamPolicy         = &cobra.Command{
		Use: "SetIamPolicy",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.SetIamPolicy")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.GetIamPolicy.
var (
	einride_authz_v1beta1_AuthorizationService_GetIamPolicy_Request v1.GetIamPolicyRequest
	einride_authz_v1beta1_AuthorizationService_GetIamPolicy         = &cobra.Command{
		Use: "GetIamPolicy",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.GetIamPolicy")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.TestIamPermissions.
var (
	einride_authz_v1beta1_AuthorizationService_TestIamPermissions_Request v1.TestIamPermissionsRequest
	einride_authz_v1beta1_AuthorizationService_TestIamPermissions         = &cobra.Command{
		Use: "TestIamPermissions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.TestIamPermissions")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.ListRoles.
var (
	einride_authz_v1beta1_AuthorizationService_ListRoles_Request v11.ListRolesRequest
	einride_authz_v1beta1_AuthorizationService_ListRoles         = &cobra.Command{
		Use: "ListRoles",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.ListRoles")
			return nil
		},
	}
)

// einride.authz.v1beta1.AuthorizationService.GetRole.
var (
	einride_authz_v1beta1_AuthorizationService_GetRole_Request v11.GetRoleRequest
	einride_authz_v1beta1_AuthorizationService_GetRole         = &cobra.Command{
		Use: "GetRole",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.authz.v1beta1.AuthorizationService.GetRole")
			return nil
		},
	}
)

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
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_TestPermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SetIamPolicy)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetIamPolicy)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_TestIamPermissions)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_ListRoles)
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetRole)
}
