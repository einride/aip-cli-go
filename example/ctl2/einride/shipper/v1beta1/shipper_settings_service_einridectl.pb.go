package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.shipper.v1beta1.ShipperSettingsService.
var (
	einride_shipper_v1beta1_ShipperSettingsServiceClient v1beta1.ShipperSettingsServiceClient
	einride_shipper_v1beta1_ShipperSettingsService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipperSettingsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipperSettingsServiceClient = v1beta1.NewShipperSettingsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperSettingsService.GetShipperSettings.
var (
	einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings_Request v1beta1.GetShipperSettingsRequest
	einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings         = &cobra.Command{
		Use: "GetShipperSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperSettingsService.GetShipperSettings")
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipperSettingsService.UpdateShipperSettings.
var (
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request v1beta1.UpdateShipperSettingsRequest
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings         = &cobra.Command{
		Use: "UpdateShipperSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.shipper.v1beta1.ShipperSettingsService.UpdateShipperSettings")
			return nil
		},
	}
)

func AddShipperSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings)
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings)
}
