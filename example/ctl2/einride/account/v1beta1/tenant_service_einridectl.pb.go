package accountv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
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
			log.Println("einride.account.v1beta1.TenantService.CreateTenant")
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
			log.Println("einride.account.v1beta1.TenantService.GetTenant")
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
			log.Println("einride.account.v1beta1.TenantService.BatchGetTenants")
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
			log.Println("einride.account.v1beta1.TenantService.ListTenants")
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
			log.Println("einride.account.v1beta1.TenantService.SearchTenants")
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
			log.Println("einride.account.v1beta1.TenantService.ResolveTenant")
			return nil
		},
	}
)

func AddTenantServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_TenantService)
}

func init() {
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_CreateTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_GetTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_BatchGetTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ListTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_SearchTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ResolveTenant)
}
