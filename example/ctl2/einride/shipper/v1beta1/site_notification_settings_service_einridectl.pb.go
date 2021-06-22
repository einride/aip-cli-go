package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_SiteNotificationSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.SiteNotificationSettingsService",
}

var einride_shipper_v1beta1_GetSiteNotificationSettingsRequest v1beta1.GetSiteNotificationSettingsRequest
var einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings = &cobra.Command{
	Use: "GetSiteNotificationSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteNotificationSettingsService.GetSiteNotificationSettings")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateSiteNotificationSettingsRequest v1beta1.UpdateSiteNotificationSettingsRequest
var einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings = &cobra.Command{
	Use: "UpdateSiteNotificationSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.SiteNotificationSettingsService.UpdateSiteNotificationSettings")
		return nil
	},
}

func AddSiteNotificationSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService)
}

func init() {
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_GetSiteNotificationSettings)
	einride_shipper_v1beta1_SiteNotificationSettingsService.AddCommand(einride_shipper_v1beta1_SiteNotificationSettingsService_UpdateSiteNotificationSettings)
}
