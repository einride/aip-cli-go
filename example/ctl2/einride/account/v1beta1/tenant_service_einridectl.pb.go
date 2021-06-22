package accountv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_account_v1beta1_TenantService = &cobra.Command{
	Use: "einride.account.v1beta1.TenantService",
}

var einride_account_v1beta1_CreateTenantRequest v1beta1.CreateTenantRequest
var einride_account_v1beta1_TenantService_CreateTenant = &cobra.Command{
	Use: "CreateTenant",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.TenantService.CreateTenant")
		return nil
	},
}

var einride_account_v1beta1_GetTenantRequest v1beta1.GetTenantRequest
var einride_account_v1beta1_TenantService_GetTenant = &cobra.Command{
	Use: "GetTenant",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.TenantService.GetTenant")
		return nil
	},
}

var einride_account_v1beta1_BatchGetTenantsRequest v1beta1.BatchGetTenantsRequest
var einride_account_v1beta1_TenantService_BatchGetTenants = &cobra.Command{
	Use: "BatchGetTenants",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.TenantService.BatchGetTenants")
		return nil
	},
}

var einride_account_v1beta1_ListTenantsRequest v1beta1.ListTenantsRequest
var einride_account_v1beta1_TenantService_ListTenants = &cobra.Command{
	Use: "ListTenants",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.TenantService.ListTenants")
		return nil
	},
}

var einride_account_v1beta1_SearchTenantsRequest v1beta1.SearchTenantsRequest
var einride_account_v1beta1_TenantService_SearchTenants = &cobra.Command{
	Use: "SearchTenants",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.TenantService.SearchTenants")
		return nil
	},
}

func AddTenantServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_TenantService)
}

func init() {
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_CreateTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_GetTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_BatchGetTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ListTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_SearchTenants)
}
