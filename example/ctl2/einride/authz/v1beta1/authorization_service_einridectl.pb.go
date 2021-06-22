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

	einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_GetAuthorizationPolicy_Request.Resource, "resource", "", "Resource name of the resource to get the authorization policy for.\n\n(-- api-linter: core::0131::request-unknown-fields=disabled\n    aip.dev/not-precedent: This is not a standard GET method. --)")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SearchRoleBindings)

	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request.Resource, "resource", "", "Resource name of the resource to search role bindings for.\nRole bindings applied to this resource, and its' descendants are\nreturned. For example, specifying shippers/1 would also return role bindings\napplied to shippers/1/sites/1.")

	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request.User, "user", "", "Resource name of the user on a role binding.\nIf set, only role bindings which this user is a part of is returned, and will\nbe the only user of the role binding.")

	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request.Tenant, "tenant", "", "Resource name of the tenant on a role binding.\nIf set, only role bindings which this tenant is a part of is returned, and will\nbe the only tenant of the role binding.")

	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings.Flags().Int32Var(&einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_authz_v1beta1_AuthorizationService_SearchRoleBindings.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchRoleBindings_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_AddRoleBinding)

	einride_authz_v1beta1_AuthorizationService_AddRoleBinding_Request.RoleBinding = new(v1beta1.RoleBinding)
	einride_authz_v1beta1_AuthorizationService_AddRoleBinding.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_AddRoleBinding_Request.RoleBinding.Resource, "roleBinding.resource", "", "Resource name of the resource the role binding applies to.")
	// TODO: enum Role
	einride_authz_v1beta1_AuthorizationService_AddRoleBinding.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_AddRoleBinding_Request.RoleBinding.Users, "roleBinding.users", []string{}, "Resource names of the users bound to the role.")
	einride_authz_v1beta1_AuthorizationService_AddRoleBinding.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_AddRoleBinding_Request.RoleBinding.Tenants, "roleBinding.tenants", []string{}, "Resource names of the tenants bound to the role.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding)

	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding_Request.RoleBinding = new(v1beta1.RoleBinding)
	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding_Request.RoleBinding.Resource, "roleBinding.resource", "", "Resource name of the resource the role binding applies to.")
	// TODO: enum Role
	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding_Request.RoleBinding.Users, "roleBinding.users", []string{}, "Resource names of the users bound to the role.")
	einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_RemoveRoleBinding_Request.RoleBinding.Tenants, "roleBinding.tenants", []string{}, "Resource names of the tenants bound to the role.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions)

	einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions_Request.Resource, "resource", "", "Resource name of the resource to lookup effective permissions for.")

	einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_LookupEffectivePermissions_Request.User, "user", "", "Resource name of the user to lookup effective permissions for.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions)

	einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions_Request.Resources, "resources", []string{}, "Resource names of the resources to lookup effective permissions for.\nA maximum of 1000 resources can be looked up in a batch.")

	einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_BatchLookupEffectivePermissions_Request.User, "user", "", "Resource name of the user to lookup effective permissions for.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings)

	// TODO: enum Permission

	einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_LookupPermissionBindings_Request.User, "user", "", "Resource name of the user to lookup permissions for.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources)

	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request.Parent, "parent", "", "Resource name of the parent to search under.\nIf not specified, all authorized resources will be returned.")

	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request.Permission, "permission", "", "The permission to lookup authorized resources for.")

	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request.Members, "members", []string{}, "The members to lookup authorized resources for.")

	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources.Flags().Int32Var(&einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 10000; values above 10000 will be coerced to 10000.")

	einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SearchAuthorizedResources_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_TestPermissions)

	einride_authz_v1beta1_AuthorizationService_TestPermissions.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_TestPermissions_Request.Permissions, "permissions", []string{}, "The permissions to test.\nAll combinations of permissions, resources and members will be tested.")

	einride_authz_v1beta1_AuthorizationService_TestPermissions.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_TestPermissions_Request.Resources, "resources", []string{}, "Resource names of the resources to test.\nAll combinations of permissions, resources and members will be tested.")

	einride_authz_v1beta1_AuthorizationService_TestPermissions.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_TestPermissions_Request.Members, "members", []string{}, "The members to test.\nAll combinations of permissions, resources and members will be tested.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_SetIamPolicy)

	einride_authz_v1beta1_AuthorizationService_SetIamPolicy.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_SetIamPolicy_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy is being specified.\nSee the operation documentation for the appropriate value for this field.")

	einride_authz_v1beta1_AuthorizationService_SetIamPolicy_Request.Policy = new(v1.Policy)
	einride_authz_v1beta1_AuthorizationService_SetIamPolicy.Flags().Int32Var(&einride_authz_v1beta1_AuthorizationService_SetIamPolicy_Request.Policy.Version, "policy.version", 10, "Specifies the format of the policy.\n\nValid values are 0, 1, and 3. Requests specifying an invalid value will be\nrejected.\n\nOperations affecting conditional bindings must specify version 3. This can\nbe either setting a conditional policy, modifying a conditional binding,\nor removing a binding (conditional or unconditional) from the stored\nconditional policy.\nOperations on non-conditional policies may specify any valid value or\nleave the field unset.\n\nIf no etag is provided in the call to `setIamPolicy`, version compliance\nchecks against the stored policy is skipped.")
	// TODO: list: Bindings message
	// TODO: bytes Etag
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetIamPolicy)

	einride_authz_v1beta1_AuthorizationService_GetIamPolicy.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_GetIamPolicy_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy is being requested.\nSee the operation documentation for the appropriate value for this field.")

	einride_authz_v1beta1_AuthorizationService_GetIamPolicy_Request.Options = new(v1.GetPolicyOptions)
	einride_authz_v1beta1_AuthorizationService_GetIamPolicy.Flags().Int32Var(&einride_authz_v1beta1_AuthorizationService_GetIamPolicy_Request.Options.RequestedPolicyVersion, "options.requestedPolicyVersion", 10, "Optional. The policy format version to be returned.\n\nValid values are 0, 1, and 3. Requests specifying an invalid value will be\nrejected.\n\nRequests for policies with any conditional bindings must specify version 3.\nPolicies without any conditional bindings may specify any valid value or\nleave the field unset.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_TestIamPermissions)

	einride_authz_v1beta1_AuthorizationService_TestIamPermissions.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_TestIamPermissions_Request.Resource, "resource", "", "REQUIRED: The resource for which the policy detail is being requested.\nSee the operation documentation for the appropriate value for this field.")

	einride_authz_v1beta1_AuthorizationService_TestIamPermissions.Flags().StringSliceVar(&einride_authz_v1beta1_AuthorizationService_TestIamPermissions_Request.Permissions, "permissions", []string{}, "The set of permissions to check for the `resource`. Permissions with\nwildcards (such as '*' or 'storage.*') are not allowed. For more\ninformation see\n[IAM Overview](https://cloud.google.com/iam/docs/overview#permissions).")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_ListRoles)

	einride_authz_v1beta1_AuthorizationService_ListRoles.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_ListRoles_Request.Parent, "parent", "", "The `parent` parameter's value depends on the target resource for the\nrequest, namely\n[`roles`](/iam/reference/rest/v1/roles),\n[`projects`](/iam/reference/rest/v1/projects.roles), or\n[`organizations`](/iam/reference/rest/v1/organizations.roles). Each\nresource type's `parent` value format is described below:\n\n* [`roles.list()`](/iam/reference/rest/v1/roles/list): An empty string.\n  This method doesn't require a resource; it simply returns all\n  [predefined roles](/iam/docs/understanding-roles#predefined_roles) in\n  Cloud IAM. Example request URL:\n  `https://iam.googleapis.com/v1/roles`\n\n* [`projects.roles.list()`](/iam/reference/rest/v1/projects.roles/list):\n  `projects/{PROJECT_ID}`. This method lists all project-level\n  [custom roles](/iam/docs/understanding-custom-roles).\n  Example request URL:\n  `https://iam.googleapis.com/v1/projects/{PROJECT_ID}/roles`\n\n* [`organizations.roles.list()`](/iam/reference/rest/v1/organizations.roles/list):\n  `organizations/{ORGANIZATION_ID}`. This method lists all\n  organization-level [custom roles](/iam/docs/understanding-custom-roles).\n  Example request URL:\n  `https://iam.googleapis.com/v1/organizations/{ORGANIZATION_ID}/roles`\n\nNote: Wildcard (*) values are invalid; you must specify a complete project\nID or organization ID.")

	einride_authz_v1beta1_AuthorizationService_ListRoles.Flags().Int32Var(&einride_authz_v1beta1_AuthorizationService_ListRoles_Request.PageSize, "pageSize", 10, "Optional limit on the number of roles to include in the response.\n\nThe default is 300, and the maximum is 1,000.")

	einride_authz_v1beta1_AuthorizationService_ListRoles.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_ListRoles_Request.PageToken, "pageToken", "", "Optional pagination token returned in an earlier ListRolesResponse.")

	// TODO: enum View

	einride_authz_v1beta1_AuthorizationService_ListRoles.Flags().BoolVar(&einride_authz_v1beta1_AuthorizationService_ListRoles_Request.ShowDeleted, "showDeleted", false, "Include Roles that have been deleted.")
	einride_authz_v1beta1_AuthorizationService.AddCommand(einride_authz_v1beta1_AuthorizationService_GetRole)

	einride_authz_v1beta1_AuthorizationService_GetRole.Flags().StringVar(&einride_authz_v1beta1_AuthorizationService_GetRole_Request.Name, "name", "", "The `name` parameter's value depends on the target resource for the\nrequest, namely\n[`roles`](/iam/reference/rest/v1/roles),\n[`projects`](/iam/reference/rest/v1/projects.roles), or\n[`organizations`](/iam/reference/rest/v1/organizations.roles). Each\nresource type's `name` value format is described below:\n\n* [`roles.get()`](/iam/reference/rest/v1/roles/get): `roles/{ROLE_NAME}`.\n  This method returns results from all\n  [predefined roles](/iam/docs/understanding-roles#predefined_roles) in\n  Cloud IAM. Example request URL:\n  `https://iam.googleapis.com/v1/roles/{ROLE_NAME}`\n\n* [`projects.roles.get()`](/iam/reference/rest/v1/projects.roles/get):\n  `projects/{PROJECT_ID}/roles/{CUSTOM_ROLE_ID}`. This method returns only\n  [custom roles](/iam/docs/understanding-custom-roles) that have been\n  created at the project level. Example request URL:\n  `https://iam.googleapis.com/v1/projects/{PROJECT_ID}/roles/{CUSTOM_ROLE_ID}`\n\n* [`organizations.roles.get()`](/iam/reference/rest/v1/organizations.roles/get):\n  `organizations/{ORGANIZATION_ID}/roles/{CUSTOM_ROLE_ID}`. This method\n  returns only [custom roles](/iam/docs/understanding-custom-roles) that\n  have been created at the organization level. Example request URL:\n  `https://iam.googleapis.com/v1/organizations/{ORGANIZATION_ID}/roles/{CUSTOM_ROLE_ID}`\n\nNote: Wildcard (*) values are invalid; you must specify a complete project\nID or organization ID.")
}
