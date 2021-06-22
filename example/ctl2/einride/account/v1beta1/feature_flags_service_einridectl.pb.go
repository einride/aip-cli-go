package accountv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.account.v1beta1.FeatureFlagsService.
var (
	einride_account_v1beta1_FeatureFlagsServiceClient v1beta1.FeatureFlagsServiceClient
	einride_account_v1beta1_FeatureFlagsService       = &cobra.Command{
		Use: "einride.account.v1beta1.FeatureFlagsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_account_v1beta1_FeatureFlagsServiceClient = v1beta1.NewFeatureFlagsServiceClient(conn)
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.GetUserFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags_Request v1beta1.GetUserFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags         = &cobra.Command{
		Use: "GetUserFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.GetUserFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.UpdateUserFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request v1beta1.UpdateUserFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags         = &cobra.Command{
		Use: "UpdateUserFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.UpdateUserFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.AddUserFeatureFlag.
var (
	einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag_Request v1beta1.AddUserFeatureFlagRequest
	einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag         = &cobra.Command{
		Use: "AddUserFeatureFlag",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.AddUserFeatureFlag")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.RemoveUserFeatureFlag.
var (
	einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag_Request v1beta1.RemoveUserFeatureFlagRequest
	einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag         = &cobra.Command{
		Use: "RemoveUserFeatureFlag",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.RemoveUserFeatureFlag")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.GetTenantFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags_Request v1beta1.GetTenantFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags         = &cobra.Command{
		Use: "GetTenantFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.GetTenantFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.UpdateTenantFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request v1beta1.UpdateTenantFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags         = &cobra.Command{
		Use: "UpdateTenantFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.UpdateTenantFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.ComputeEffectiveFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags_Request v1beta1.ComputeEffectiveFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags         = &cobra.Command{
		Use: "ComputeEffectiveFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.ComputeEffectiveFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.ComputeUserFeatureFlags.
var (
	einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags_Request v1beta1.ComputeUserFeatureFlagsRequest
	einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags         = &cobra.Command{
		Use: "ComputeUserFeatureFlags",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.ComputeUserFeatureFlags")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.CreateFeatureFlag.
var (
	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request v1beta1.CreateFeatureFlagRequest
	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag         = &cobra.Command{
		Use: "CreateFeatureFlag",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.CreateFeatureFlag")
			return nil
		},
	}
)

// einride.account.v1beta1.FeatureFlagsService.DeleteFeatureFlag.
var (
	einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag_Request v1beta1.DeleteFeatureFlagRequest
	einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag         = &cobra.Command{
		Use: "DeleteFeatureFlag",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.account.v1beta1.FeatureFlagsService.DeleteFeatureFlag")
			return nil
		},
	}
)

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
