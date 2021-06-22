package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsServiceClient v1beta1.SiteNotificationSettingsServiceClient
	einride_shipper_v1beta1_SiteNotificationSettingsService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.SiteNotificationSettingsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_SiteNotificationSettingsServiceClient = v1beta1.NewSiteNotificationSettingsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.GetSiteNotificationSettings.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings_Request v1beta1.GetSiteNotificationSettingsRequest
	einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings         = &cobra.Command{
		Use: "GetSiteNotificationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteNotificationSettingsService.GetSiteNotificationSettings")
			return nil
		},
	}
)

// einride.shipper.v1beta1.SiteNotificationSettingsService.UpdateSiteNotificationSettings.
var (
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings_Request v1beta1.UpdateSiteNotificationSettingsRequest
	einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings         = &cobra.Command{
		Use: "UpdateSiteNotificationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.SiteNotificationSettingsService.UpdateSiteNotificationSettings")
			return nil
		},
	}
)

func AddSiteNotificationSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService)
}

func init() {
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings)
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings)
}
