package accountv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/account/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

	einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_GetUserFeatureFlags_Request.Name, "name", "", "Resource name of the user feature flags to retrieve.\nFormat: `tenants/{tenant}/users/{user}/featureFlags`.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags)

	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags = new(v1beta1.UserFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.Name, "userFeatureFlags.name", "", "The resource name.")
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.FeatureFlags = new(v1beta1.FeatureFlags)
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.FeatureFlags.Onboarding, "userFeatureFlags.featureFlags.onboarding", false, "The user or tenant is currently onboarding.")
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.FeatureFlags.TransportSchedules, "userFeatureFlags.featureFlags.transportSchedules", false, "The user or tenant should view transport schedules in Planning view.")
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.FeatureFlags.ShippingCost, "userFeatureFlags.featureFlags.shippingCost", false, "The user or tenant should view shipping costs in Dashboard view.")
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UserFeatureFlags.FeatureFlags.LastMinuteChanges, "userFeatureFlags.featureFlags.lastMinuteChanges", false, "The user or tenant should be able to do last-minute-changes to a shipment in Driver app.")

	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags.Flags().StringSliceVar(&einride_account_v1beta1_FeatureFlagsService_UpdateUserFeatureFlags_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag)

	einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag_Request.Name, "name", "", "Resource name of the user feature flag.\nFormat: `tenants/{tenant}/users/{user}/featureFlags`")

	einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_AddUserFeatureFlag_Request.FeatureFlag, "featureFlag", "", "Resource name of the feature flag to add.\nFormat: `featureFlags/{feature_flag}")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag)

	einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag_Request.Name, "name", "", "Resource name of the user feature flag.\nFormat: `tenants/{tenant}/users/{user}/featureFlags`")

	einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_RemoveUserFeatureFlag_Request.FeatureFlag, "featureFlag", "", "Resource name of the feature flag to remove.\nFormat: `featureFlags/{feature_flag}")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags)

	einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_GetTenantFeatureFlags_Request.Name, "name", "", "Resource name of the tenant feature flags to retrieve.\nFormat: `tenants/{tenant}/featureFlags`.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags)

	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags = new(v1beta1.TenantFeatureFlags)
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.Name, "tenantFeatureFlags.name", "", "The resource name.")
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.FeatureFlags = new(v1beta1.FeatureFlags)
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.FeatureFlags.Onboarding, "tenantFeatureFlags.featureFlags.onboarding", false, "The user or tenant is currently onboarding.")
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.FeatureFlags.TransportSchedules, "tenantFeatureFlags.featureFlags.transportSchedules", false, "The user or tenant should view transport schedules in Planning view.")
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.FeatureFlags.ShippingCost, "tenantFeatureFlags.featureFlags.shippingCost", false, "The user or tenant should view shipping costs in Dashboard view.")
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().BoolVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.TenantFeatureFlags.FeatureFlags.LastMinuteChanges, "tenantFeatureFlags.featureFlags.lastMinuteChanges", false, "The user or tenant should be able to do last-minute-changes to a shipment in Driver app.")

	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags.Flags().StringSliceVar(&einride_account_v1beta1_FeatureFlagsService_UpdateTenantFeatureFlags_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags)

	einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_ComputeEffectiveFeatureFlags_Request.User, "user", "", "Resource name of the user to compute effective feature flags for.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags)

	einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_ComputeUserFeatureFlags_Request.User, "user", "", "Resource name of the user to compute feature flags for.\nFormat: `tenants/{tenant}/users/{user}`.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag)

	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request.FeatureFlag = new(v1beta1.FeatureFlag)
	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request.FeatureFlag.Name, "featureFlag.name", "", "The resource name of the feature flag.")
	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request.FeatureFlag.Description, "featureFlag.description", "", "A description of the feature flag.")
	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request.FeatureFlag.Owner, "featureFlag.owner", "", "The person or team responsible for managing this feature flag.")

	einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_CreateFeatureFlag_Request.FeatureFlagId, "featureFlagId", "", "The ID to use for the feature flag, which will become the final component of\nthe feature flags resource name.\n\nThis value should be 4-20 characters, start with a letter and\nonly consist of Upper and Lower case letters, digits and hyphens.")
	einride_account_v1beta1_FeatureFlagsService.AddCommand(einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag)

	einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag.Flags().StringVar(&einride_account_v1beta1_FeatureFlagsService_DeleteFeatureFlag_Request.Name, "name", "", "Resource name of the feature flag to delete.\nFormat: `featureFlags/{feature_flag}")
}
