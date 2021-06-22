package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipperUserSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipperUserSettingsService",
}

var einride_shipper_v1beta1_GetShipperUserSettingsRequest v1beta1.GetShipperUserSettingsRequest
var einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings = &cobra.Command{
	Use: "GetShipperUserSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperUserSettingsService.GetShipperUserSettings")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipperUserSettingsRequest v1beta1.UpdateShipperUserSettingsRequest
var einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings = &cobra.Command{
	Use: "UpdateShipperUserSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperUserSettingsService.UpdateShipperUserSettings")
		return nil
	},
}

func AddShipperUserSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings)
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings)
}
