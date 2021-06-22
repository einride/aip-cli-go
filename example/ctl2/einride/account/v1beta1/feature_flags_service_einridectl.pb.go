package accountv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_account_v1beta1_FeatureFlagsService = &cobra.Command{
	Use: "einride.account.v1beta1.FeatureFlagsService",
}

var einride_account_v1beta1_GetUserFeatureFlagsRequest v1beta1.GetUserFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags = &cobra.Command{
	Use: "GetUserFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.GetUserFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_UpdateUserFeatureFlagsRequest v1beta1.UpdateUserFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags = &cobra.Command{
	Use: "UpdateUserFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.UpdateUserFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_AddUserFeatureFlagRequest v1beta1.AddUserFeatureFlagRequest
var einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag = &cobra.Command{
	Use: "AddUserFeatureFlag",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.AddUserFeatureFlag")
		return nil
	},
}

var einride_account_v1beta1_RemoveUserFeatureFlagRequest v1beta1.RemoveUserFeatureFlagRequest
var einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag = &cobra.Command{
	Use: "RemoveUserFeatureFlag",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.RemoveUserFeatureFlag")
		return nil
	},
}

var einride_account_v1beta1_GetTenantFeatureFlagsRequest v1beta1.GetTenantFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags = &cobra.Command{
	Use: "GetTenantFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.GetTenantFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_UpdateTenantFeatureFlagsRequest v1beta1.UpdateTenantFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags = &cobra.Command{
	Use: "UpdateTenantFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.UpdateTenantFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_ComputeEffectiveFeatureFlagsRequest v1beta1.ComputeEffectiveFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags = &cobra.Command{
	Use: "ComputeEffectiveFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.ComputeEffectiveFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_ComputeUserFeatureFlagsRequest v1beta1.ComputeUserFeatureFlagsRequest
var einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags = &cobra.Command{
	Use: "ComputeUserFeatureFlags",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.ComputeUserFeatureFlags")
		return nil
	},
}

var einride_account_v1beta1_CreateFeatureFlagRequest v1beta1.CreateFeatureFlagRequest
var einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag = &cobra.Command{
	Use: "CreateFeatureFlag",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.CreateFeatureFlag")
		return nil
	},
}

var einride_account_v1beta1_DeleteFeatureFlagRequest v1beta1.DeleteFeatureFlagRequest
var einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag = &cobra.Command{
	Use: "DeleteFeatureFlag",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.account.v1beta1.FeatureFlagsService.DeleteFeatureFlag")
		return nil
	},
}

func AddFeatureFlagsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_account_v1beta1_FeatureFlagsService)
}

func init() {
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
}
