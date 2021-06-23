package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.ShipperUserSettingsService.
var (
	einride_shipper_v1beta1_ShipperUserSettingsServiceClient v1beta1.ShipperUserSettingsServiceClient
	einride_shipper_v1beta1_ShipperUserSettingsService       = &cobra.Command{
		Use:   "shipper.v1beta1.ShipperUserSettingsService",
		Short: "Shipper user settings service.",
		Long:  "Shipper user settings service.",
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
			response, err := einride_shipper_v1beta1_ShipperUserSettingsServiceClient.GetShipperUserSettings(cmd.Context(), &einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_v1beta1_ShipperUserSettingsServiceClient.UpdateShipperUserSettings(cmd.Context(), &einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipperUserSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings)

	einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperUserSettingsService_GetShipperUserSettings_Request.Name, "name", "", "Resource name of the shipper user settings to retrieve.")
	einride_shipper_v1beta1_ShipperUserSettingsService.AddCommand(einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings)

	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.ShipperUserSettings = new(v1beta1.ShipperUserSettings)
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.ShipperUserSettings.Name, "shipperUserSettings.name", "", "The resource name of the shipper user settings.")
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.ShipperUserSettings.Etag, "shipperUserSettings.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.ShipperUserSettings.DefaultSite, "shipperUserSettings.defaultSite", "", "Resource name of the user's default site.")
	// TODO: enum Persona

	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipperUserSettingsService_UpdateShipperUserSettings_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
}
