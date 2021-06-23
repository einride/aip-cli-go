package accountv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.account.v1beta1.TenantService.
var (
	einride_account_v1beta1_TenantServiceClient v1beta1.TenantServiceClient
	einride_account_v1beta1_TenantService       = &cobra.Command{
		Use: "einride.account.v1beta1.TenantService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_account_v1beta1_TenantServiceClient = v1beta1.NewTenantServiceClient(conn)
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.CreateTenant.
var (
	einride_account_v1beta1_TenantService_CreateTenant_Request v1beta1.CreateTenantRequest
	einride_account_v1beta1_TenantService_CreateTenant         = &cobra.Command{
		Use: "CreateTenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.CreateTenant(cmd.Context(), &einride_account_v1beta1_TenantService_CreateTenant_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.GetTenant.
var (
	einride_account_v1beta1_TenantService_GetTenant_Request v1beta1.GetTenantRequest
	einride_account_v1beta1_TenantService_GetTenant         = &cobra.Command{
		Use: "GetTenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.GetTenant(cmd.Context(), &einride_account_v1beta1_TenantService_GetTenant_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.BatchGetTenants.
var (
	einride_account_v1beta1_TenantService_BatchGetTenants_Request v1beta1.BatchGetTenantsRequest
	einride_account_v1beta1_TenantService_BatchGetTenants         = &cobra.Command{
		Use: "BatchGetTenants",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.BatchGetTenants(cmd.Context(), &einride_account_v1beta1_TenantService_BatchGetTenants_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.ListTenants.
var (
	einride_account_v1beta1_TenantService_ListTenants_Request v1beta1.ListTenantsRequest
	einride_account_v1beta1_TenantService_ListTenants         = &cobra.Command{
		Use: "ListTenants",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.ListTenants(cmd.Context(), &einride_account_v1beta1_TenantService_ListTenants_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.SearchTenants.
var (
	einride_account_v1beta1_TenantService_SearchTenants_Request v1beta1.SearchTenantsRequest
	einride_account_v1beta1_TenantService_SearchTenants         = &cobra.Command{
		Use: "SearchTenants",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.SearchTenants(cmd.Context(), &einride_account_v1beta1_TenantService_SearchTenants_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.account.v1beta1.TenantService.ResolveTenant.
var (
	einride_account_v1beta1_TenantService_ResolveTenant_Request v1beta1.ResolveTenantRequest
	einride_account_v1beta1_TenantService_ResolveTenant         = &cobra.Command{
		Use: "ResolveTenant",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_account_v1beta1_TenantServiceClient.ResolveTenant(cmd.Context(), &einride_account_v1beta1_TenantService_ResolveTenant_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddTenantServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_TenantService)
}

func init() {
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_CreateTenant)

	einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant = new(v1beta1.Tenant)
	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant.Name, "tenant.name", "", "The resource name of the tenant.")
	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant.Etag, "tenant.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant.DisplayName, "tenant.displayName", "", "Display name of the tenant.\nFor example: \"Johnson Trucking Co.\"")
	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant.Shipper, "tenant.shipper", "", "Resource name of the tenant's shipper, if the tenant is a shipper.")
	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.Tenant.Carrier, "tenant.carrier", "", "Resource name of the tenant's carrier, if the tenant is a carrier.")

	einride_account_v1beta1_TenantService_CreateTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_CreateTenant_Request.TenantId, "tenantId", "", "The ID to use for the tenant, which will become the final component of\nthe role's resource name.\n\nThis value should be 4-20 characters, start with a letter and\nonly consist of letters, digits and hyphens.")
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_GetTenant)

	einride_account_v1beta1_TenantService_GetTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_GetTenant_Request.Name, "name", "", "Resource name of the tenant to retrieve.")
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_BatchGetTenants)

	einride_account_v1beta1_TenantService_BatchGetTenants.Flags().StringSliceVar(&einride_account_v1beta1_TenantService_BatchGetTenants_Request.Names, "names", []string{}, "Resource names of the tenants to retrieve.\nA maximum of 1000 tenants can be retrieved in a batch.")
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ListTenants)

	einride_account_v1beta1_TenantService_ListTenants.Flags().Int32Var(&einride_account_v1beta1_TenantService_ListTenants_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_account_v1beta1_TenantService_ListTenants.Flags().StringVar(&einride_account_v1beta1_TenantService_ListTenants_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_SearchTenants)

	einride_account_v1beta1_TenantService_SearchTenants.Flags().Int32Var(&einride_account_v1beta1_TenantService_SearchTenants_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_account_v1beta1_TenantService_SearchTenants.Flags().StringVar(&einride_account_v1beta1_TenantService_SearchTenants_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_account_v1beta1_TenantService_SearchTenants.Flags().StringVar(&einride_account_v1beta1_TenantService_SearchTenants_Request.UserEmail, "userEmail", "", "If set, only tenants that have a user with the provided email are returned.")
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ResolveTenant)

	einride_account_v1beta1_TenantService_ResolveTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_ResolveTenant_Request.UserEmail, "userEmail", "", "The user email to lookup tenants for")

	einride_account_v1beta1_TenantService_ResolveTenant.Flags().StringVar(&einride_account_v1beta1_TenantService_ResolveTenant_Request.UserPassword, "userPassword", "", "The user password needed to verify the user's identity")
}
