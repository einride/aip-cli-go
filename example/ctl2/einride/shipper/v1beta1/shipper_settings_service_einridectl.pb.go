package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipperSettingsService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipperSettingsService",
}

var einride_shipper_v1beta1_GetShipperSettingsRequest v1beta1.GetShipperSettingsRequest
var einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings = &cobra.Command{
	Use: "GetShipperSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperSettingsService.GetShipperSettings")
		return nil
	},
}

var einride_shipper_v1beta1_UpdateShipperSettingsRequest v1beta1.UpdateShipperSettingsRequest
var einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings = &cobra.Command{
	Use: "UpdateShipperSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipperSettingsService.UpdateShipperSettings")
		return nil
	},
}

func AddShipperSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings)
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings)
}
