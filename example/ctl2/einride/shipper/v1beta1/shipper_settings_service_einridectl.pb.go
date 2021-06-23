package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
			response, err := einride_shipper_v1beta1_ShipperSettingsServiceClient.GetShipperSettings(cmd.Context(), &einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_shipper_v1beta1_ShipperSettingsServiceClient.UpdateShipperSettings(cmd.Context(), &einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipperSettingsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipperSettingsService)
}

func init() {
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings)

	einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperSettingsService_GetShipperSettings_Request.Name, "name", "", "Resource name of the shipper settings to retrieve.")
	einride_shipper_v1beta1_ShipperSettingsService.AddCommand(einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings)

	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings = new(v1beta1.ShipperSettings)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.Name, "shipperSettings.name", "", "The resource name of the shipper user settings.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().StringVar(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.Etag, "shipperSettings.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().BoolVar(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.AutoApproveShipmentChangeProposals, "shipperSettings.autoApproveShipmentChangeProposals", false, "Flag indicating if shipment change proposals created for this shipper\nshould be auto-approved.\n\nDefaults to no auto-approve when unset.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int64Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupEarliestTime.Seconds, "shipperSettings.pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int32Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupEarliestTime.Nanos, "shipperSettings.pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int64Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupLatestTime.Seconds, "shipperSettings.pickupLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int32Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.PickupLatestTime.Nanos, "shipperSettings.pickupLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int64Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryEarliestTime.Seconds, "shipperSettings.deliveryEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int32Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryEarliestTime.Nanos, "shipperSettings.deliveryEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int64Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryLatestTime.Seconds, "shipperSettings.deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().Int32Var(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.ShipperSettings.DeliveryLatestTime.Nanos, "shipperSettings.deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipperSettingsService_UpdateShipperSettings_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
}
