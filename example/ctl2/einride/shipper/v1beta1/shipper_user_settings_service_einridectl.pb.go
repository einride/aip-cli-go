package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipperUserSettingsService.
var (
	einride_shipper_v1beta1_ShipperUserSettingsServiceClient v1beta1.ShipperUserSettingsServiceClient
	einride_shipper_v1beta1_ShipperUserSettingsService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipperUserSettingsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipperUserSettingsServiceClient = v1beta1.NewShipperUserSettingsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperUserSettingsService.GetShipperUserSettings.
var (
	einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings_Request v1beta1.GetShipperUserSettingsRequest
	einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings         = &cobra.Command{
		Use: "GetShipperUserSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperUserSettingsService.GetShipperUserSettings")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperUserSettingsService.UpdateShipperUserSettings.
var (
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request v1beta1.UpdateShipperUserSettingsRequest
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings         = &cobra.Command{
		Use: "UpdateShipperUserSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperUserSettingsService.UpdateShipperUserSettings")
			return nil
		},
	}
)

func AddShipperUserSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings)
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings)
}
