package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	log "log"
)

// einride.planner.v1beta1.ShipmentAllocationService.
var (
	einride_planner_v1beta1_ShipmentAllocationServiceClient v1beta1.ShipmentAllocationServiceClient
	einride_planner_v1beta1_ShipmentAllocationService       = &cobra.Command{
		Use: "einride.planner.v1beta1.ShipmentAllocationService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_ShipmentAllocationServiceClient = v1beta1.NewShipmentAllocationServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.ShipmentAllocationService.GetShipmentAllocationSettings.
var (
	einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings_Request v1beta1.GetShipmentAllocationSettingsRequest
	einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings         = &cobra.Command{
		Use: "GetShipmentAllocationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ShipmentAllocationService.GetShipmentAllocationSettings")
			return nil
		},
	}
)

// einride.planner.v1beta1.ShipmentAllocationService.UpdateShipmentAllocationSettings.
var (
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request v1beta1.UpdateShipmentAllocationSettingsRequest
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings         = &cobra.Command{
		Use: "UpdateShipmentAllocationSettings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ShipmentAllocationService.UpdateShipmentAllocationSettings")
			return nil
		},
	}
)

// einride.planner.v1beta1.ShipmentAllocationService.AllocateShipments.
var (
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request v1beta1.AllocateShipmentsRequest
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments         = &cobra.Command{
		Use: "AllocateShipments",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ShipmentAllocationService.AllocateShipments")
			return nil
		},
	}
)

// einride.planner.v1beta1.ShipmentAllocationService.PublishShipmentAllocationStatuses.
var (
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request v1beta1.PublishShipmentAllocationStatusesRequest
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses         = &cobra.Command{
		Use: "PublishShipmentAllocationStatuses",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ShipmentAllocationService.PublishShipmentAllocationStatuses")
			return nil
		},
	}
)

// einride.planner.v1beta1.ShipmentAllocationService.AllocateShipmentsAutomatic.
var (
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request v1beta1.AllocateShipmentsAutomaticRequest
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic         = &cobra.Command{
		Use: "AllocateShipmentsAutomatic",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ShipmentAllocationService.AllocateShipmentsAutomatic")
			return nil
		},
	}
)

func AddShipmentAllocationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_ShipmentAllocationService)
}

func init() {
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings)

	einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings_Request.Name, "name", "", "Resource name of the allocation settings to retrieve.\nPattern: `transportInstallations/{transport_installation}/shipmentAllocationSettings`.")
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings)

	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings = new(v1beta1.ShipmentAllocationSettings)
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Name, "shipmentAllocationSettings.name", "", "The resource name of the shipment allocation settings.\nFor example: `\"transportInstallations/transportInstallation/shipmentAllocationSettings`.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Etag, "shipmentAllocationSettings.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().BoolVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Enabled, "shipmentAllocationSettings.enabled", false, "Flag indicating whether shipment allocation is enabled for the transport installation.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Shipper, "shipmentAllocationSettings.shipper", "", "Resource name of the shipper which shipments should be allocated for.\nWhen enabled = true, this field must be set.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringSliceVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Lanes, "shipmentAllocationSettings.lanes", []string{}, "Resource names of lanes on which shipments should be allocated.\nIf this list is empty, all shipments will be considered.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.CarrierReferenceId, "shipmentAllocationSettings.carrierReferenceId", "", "The carrier reference ID which shipments should have to be allocated.\nIf not set, no filtering on carrier reference ID is done.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Carrier, "shipmentAllocationSettings.carrier", "", "Resource name of the carrier who owns the vehicle tasks that shipments should be\nallocated to.\nWhen enabled = true, this field must be set.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().BoolVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.MonitoringEnabled, "shipmentAllocationSettings.monitoringEnabled", false, "Flag indicating whether monitoring is enabled for the transport installation.\nVehicle tasks on `lanes` will be monitored continuously, and if these tasks\nare unallocated errors will be triggered.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.ErrorDuration = new(durationpb.Duration)
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.ErrorDuration.Seconds, "shipmentAllocationSettings.errorDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.ErrorDuration.Nanos, "shipmentAllocationSettings.errorDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.WarningDuration = new(durationpb.Duration)
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.WarningDuration.Seconds, "shipmentAllocationSettings.warningDuration.seconds", 10, "Signed seconds of the span of time. Must be from -315,576,000,000\nto +315,576,000,000 inclusive. Note: these bounds are computed from:\n60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.WarningDuration.Nanos, "shipmentAllocationSettings.warningDuration.nanos", 10, "Signed fractions of a second at nanosecond resolution of the span\nof time. Durations less than one second are represented with a 0\n`seconds` field and a positive or negative `nanos` field. For durations\nof one second or more, a non-zero value for the `nanos` field must be\nof the same sign as the `seconds` field. Must be from -999,999,999\nto +999,999,999 inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().BoolVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.ShipmentAllocationSettings.Automatic, "shipmentAllocationSettings.automatic", false, "Flag indicating whether allocation is done automatically.\nIf set, every hour (not configurable) `AllocateShipments` will be run on the\ntransport installation.")

	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings.Flags().StringSliceVar(&einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments)

	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments.Flags().StringVar(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.TransportInstallation, "transportInstallation", "", "Resource name of the transport installation to allocate shipments for.")

	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.PickupEarliestTime.Seconds, "pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.PickupEarliestTime.Nanos, "pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.DeliveryLatestTime.Seconds, "deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments_Request.DeliveryLatestTime.Nanos, "deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses)

	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.StartTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.StartTime.Seconds, "startTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.StartTime.Nanos, "startTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.EndTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.EndTime.Seconds, "endTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses_Request.EndTime.Nanos, "endTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic)

	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.PickupEarliestTime.Seconds, "pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.PickupEarliestTime.Nanos, "pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic.Flags().Int64Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.DeliveryLatestTime.Seconds, "deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic.Flags().Int32Var(&einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic_Request.DeliveryLatestTime.Nanos, "deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
}
