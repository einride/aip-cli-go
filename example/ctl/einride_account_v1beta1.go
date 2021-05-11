package ctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	_                                             = fmt.Sprintf
	_                                             = cobra.Command{}
	einride_account_v1beta1_AuthenticationService = &cobra.Command{
		Use: "einride.account.v1beta1.Authentication",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("einride.account.v1beta1.Authentication called")
		},
	}
)

var einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken = &cobra.Command{
	Use: "AuthenticateUserIdentityToken",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AuthenticateUserIdentityToken called")
	},
}

var einride_account_v1beta1_FeatureFlagsService = &cobra.Command{
	Use: "einride.account.v1beta1.FeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.FeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags = &cobra.Command{
	Use: "GetUserFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetUserFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags = &cobra.Command{
	Use: "UpdateUserFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateUserFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag = &cobra.Command{
	Use: "AddUserFeatureFlag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AddUserFeatureFlag called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag = &cobra.Command{
	Use: "RemoveUserFeatureFlag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RemoveUserFeatureFlag called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags = &cobra.Command{
	Use: "GetTenantFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetTenantFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags = &cobra.Command{
	Use: "UpdateTenantFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateTenantFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags = &cobra.Command{
	Use: "ComputeEffectiveFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputeEffectiveFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags = &cobra.Command{
	Use: "ComputeUserFeatureFlags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ComputeUserFeatureFlags called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag = &cobra.Command{
	Use: "CreateFeatureFlag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateFeatureFlag called")
	},
}

var einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag = &cobra.Command{
	Use: "DeleteFeatureFlag",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DeleteFeatureFlag called")
	},
}

var einride_account_v1beta1_TenantService = &cobra.Command{
	Use: "einride.account.v1beta1.Tenant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.Tenant called")
	},
}

var einride_account_v1beta1_TenantService_CreateTenant = &cobra.Command{
	Use: "CreateTenant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateTenant called")
	},
}

var einride_account_v1beta1_TenantService_GetTenant = &cobra.Command{
	Use: "GetTenant",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetTenant called")
	},
}

var einride_account_v1beta1_TenantService_BatchGetTenants = &cobra.Command{
	Use: "BatchGetTenants",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetTenants called")
	},
}

var einride_account_v1beta1_TenantService_ListTenants = &cobra.Command{
	Use: "ListTenants",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListTenants called")
	},
}

var einride_account_v1beta1_TenantService_SearchTenants = &cobra.Command{
	Use: "SearchTenants",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SearchTenants called")
	},
}

var einride_account_v1beta1_UserService = &cobra.Command{
	Use: "einride.account.v1beta1.User",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("einride.account.v1beta1.User called")
	},
}

var einride_account_v1beta1_UserService_CreateUser = &cobra.Command{
	Use: "CreateUser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateUser called")
	},
}

var einride_account_v1beta1_UserService_GetUser = &cobra.Command{
	Use: "GetUser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GetUser called")
	},
}

var einride_account_v1beta1_UserService_BatchGetUsers = &cobra.Command{
	Use: "BatchGetUsers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("BatchGetUsers called")
	},
}

var einride_account_v1beta1_UserService_UpdateUser = &cobra.Command{
	Use: "UpdateUser",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateUser called")
	},
}

var einride_account_v1beta1_UserService_ListUsers = &cobra.Command{
	Use: "ListUsers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListUsers called")
	},
}

func init() {
	rootCmd.AddCommand(einride_account_v1beta1_AuthenticationService)
	einride_account_v1beta1_AuthenticationService.AddCommand(einride_account_v1beta1_AuthenticationService_AuthenticateUserIdentityToken)
	rootCmd.AddCommand(einride_account_v1beta1_FeatureFlagsService)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag)
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag)
	rootCmd.AddCommand(einride_account_v1beta1_TenantService)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_CreateTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_GetTenant)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_BatchGetTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_ListTenants)
	einride_account_v1beta1_TenantService.AddCommand(einride_account_v1beta1_TenantService_SearchTenants)
	rootCmd.AddCommand(einride_account_v1beta1_UserService)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_CreateUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_GetUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_BatchGetUsers)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_UpdateUser)
	einride_account_v1beta1_UserService.AddCommand(einride_account_v1beta1_UserService_ListUsers)
}
