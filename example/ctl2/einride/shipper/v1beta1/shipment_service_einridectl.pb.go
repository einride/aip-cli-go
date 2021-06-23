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

// einride.shipper.v1beta1.ShipmentService.
var (
	einride_shipper_v1beta1_ShipmentServiceClient v1beta1.ShipmentServiceClient
	einride_shipper_v1beta1_ShipmentService       = &cobra.Command{
		Use:   "shipper.v1beta1.ShipmentService",
		Short: "Shipment resource service.",
		Long:  "Shipment resource service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentServiceClient = v1beta1.NewShipmentServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.CreateShipment.
var (
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request v1beta1.CreateShipmentRequest
	einride_shipper_v1beta1_ShipmentService_CreateShipment         = &cobra.Command{
		Use: "CreateShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.CreateShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_CreateShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.GetShipment.
var (
	einride_shipper_v1beta1_ShipmentService_GetShipment_Request v1beta1.GetShipmentRequest
	einride_shipper_v1beta1_ShipmentService_GetShipment         = &cobra.Command{
		Use: "GetShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.GetShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_GetShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.BatchGetShipments.
var (
	einride_shipper_v1beta1_ShipmentService_BatchGetShipments_Request v1beta1.BatchGetShipmentsRequest
	einride_shipper_v1beta1_ShipmentService_BatchGetShipments         = &cobra.Command{
		Use: "BatchGetShipments",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.BatchGetShipments(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_BatchGetShipments_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.SearchShipments.
var (
	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request v1beta1.SearchShipmentsRequest
	einride_shipper_v1beta1_ShipmentService_SearchShipments         = &cobra.Command{
		Use: "SearchShipments",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.SearchShipments(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_SearchShipments_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UpdateShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request v1beta1.UpdateShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UpdateShipment         = &cobra.Command{
		Use: "UpdateShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.UpdateShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.DeleteShipment.
var (
	einride_shipper_v1beta1_ShipmentService_DeleteShipment_Request v1beta1.DeleteShipmentRequest
	einride_shipper_v1beta1_ShipmentService_DeleteShipment         = &cobra.Command{
		Use: "DeleteShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.DeleteShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_DeleteShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UndeleteShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UndeleteShipment_Request v1beta1.UndeleteShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UndeleteShipment         = &cobra.Command{
		Use: "UndeleteShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.UndeleteShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_UndeleteShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ReferenceShipment.
var (
	einride_shipper_v1beta1_ShipmentService_ReferenceShipment_Request v1beta1.ReferenceShipmentRequest
	einride_shipper_v1beta1_ShipmentService_ReferenceShipment         = &cobra.Command{
		Use: "ReferenceShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.ReferenceShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_ReferenceShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.PlanShipment.
var (
	einride_shipper_v1beta1_ShipmentService_PlanShipment_Request v1beta1.PlanShipmentRequest
	einride_shipper_v1beta1_ShipmentService_PlanShipment         = &cobra.Command{
		Use: "PlanShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.PlanShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_PlanShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnplanShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnplanShipment_Request v1beta1.UnplanShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnplanShipment         = &cobra.Command{
		Use: "UnplanShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.UnplanShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_UnplanShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.LoadShipment.
var (
	einride_shipper_v1beta1_ShipmentService_LoadShipment_Request v1beta1.LoadShipmentRequest
	einride_shipper_v1beta1_ShipmentService_LoadShipment         = &cobra.Command{
		Use: "LoadShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.LoadShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_LoadShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.LoadShipmentUnknownVehicle.
var (
	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle_Request v1beta1.LoadShipmentUnknownVehicleRequest
	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle         = &cobra.Command{
		Use: "LoadShipmentUnknownVehicle",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.LoadShipmentUnknownVehicle(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.TransitShipment.
var (
	einride_shipper_v1beta1_ShipmentService_TransitShipment_Request v1beta1.TransitShipmentRequest
	einride_shipper_v1beta1_ShipmentService_TransitShipment         = &cobra.Command{
		Use: "TransitShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.TransitShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_TransitShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ArriveShipment.
var (
	einride_shipper_v1beta1_ShipmentService_ArriveShipment_Request v1beta1.ArriveShipmentRequest
	einride_shipper_v1beta1_ShipmentService_ArriveShipment         = &cobra.Command{
		Use: "ArriveShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.ArriveShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_ArriveShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnloadShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnloadShipment_Request v1beta1.UnloadShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnloadShipment         = &cobra.Command{
		Use: "UnloadShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.UnloadShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_UnloadShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.DeliverShipment.
var (
	einride_shipper_v1beta1_ShipmentService_DeliverShipment_Request v1beta1.DeliverShipmentRequest
	einride_shipper_v1beta1_ShipmentService_DeliverShipment         = &cobra.Command{
		Use: "DeliverShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.DeliverShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_DeliverShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.HandleShipmentExternally.
var (
	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally_Request v1beta1.HandleShipmentExternallyRequest
	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally         = &cobra.Command{
		Use: "HandleShipmentExternally",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.HandleShipmentExternally(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.RejectShipment.
var (
	einride_shipper_v1beta1_ShipmentService_RejectShipment_Request v1beta1.RejectShipmentRequest
	einride_shipper_v1beta1_ShipmentService_RejectShipment         = &cobra.Command{
		Use: "RejectShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.RejectShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_RejectShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.UnrejectShipment.
var (
	einride_shipper_v1beta1_ShipmentService_UnrejectShipment_Request v1beta1.UnrejectShipmentRequest
	einride_shipper_v1beta1_ShipmentService_UnrejectShipment         = &cobra.Command{
		Use: "UnrejectShipment",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.UnrejectShipment(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_UnrejectShipment_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentService.ConsumeShipmentChangedEvent.
var (
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request v1beta1.ConsumeShipmentChangedEventRequest
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent         = &cobra.Command{
		Use: "ConsumeShipmentChangedEvent",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentServiceClient.ConsumeShipmentChangedEvent(cmd.Context(), &einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipmentServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentService)
}

func init() {
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_CreateShipment)

	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Parent, "parent", "", "Resource name of the parent shipper where this shipment will be created.")

	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment = new(v1beta1.Shipment)
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.Name, "shipment.name", "", "The resource name of the shipment.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.Etag, "shipment.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.IntegrationFile, "shipment.integrationFile", "", "Resource name of the integration file the shipment was ingested from, if the shipment was ingested.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.OriginSite, "shipment.originSite", "", "Resource name of the site from which the shipment originates.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DestinationSite, "shipment.destinationSite", "", "Resource name of the site to which the shipment is destined.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.ReferenceId, "shipment.referenceId", "", "The reference ID for this shipment from e.g. external ERP systems.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().BoolVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.Editable, "shipment.editable", false, "Flag indicating if the shipment should be editable.\nShipments may be uneditable when the master copy is stored in an external\nERP system.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupEarliestTime.Seconds, "shipment.pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupEarliestTime.Nanos, "shipment.pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupLatestTime.Seconds, "shipment.pickupLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupLatestTime.Nanos, "shipment.pickupLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryEarliestTime.Seconds, "shipment.deliveryEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryEarliestTime.Nanos, "shipment.deliveryEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryLatestTime.Seconds, "shipment.deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryLatestTime.Nanos, "shipment.deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: enum ShipmentStatus
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.Vehicle, "shipment.vehicle", "", "Resource name of the vehicle performing the delivery.\nEmpty for untracked vehicles.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.PickupInstructions, "shipment.pickupInstructions", "", "Instructions on the pickup, for example which trailer to pick up.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.DeliveryInstructions, "shipment.deliveryInstructions", "", "Instructions on delivery, which could say things like park the trailer at space 123 at lot B.")
	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.Shipment.CarrierReferenceId, "shipment.carrierReferenceId", "", "The carrier reference ID for this shipment from. Unique\nidentifier carrier to use, as listed in the Freight Mobility\nPlatform Integration Specification.")
	// TODO: list: LineItems message

	einride_shipper_v1beta1_ShipmentService_CreateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_CreateShipment_Request.ShipmentId, "shipmentId", "", "The ID to use for the shipment, which will become the final component of\nthe shipment's resource name.\n\nThis value should be 4-63 characters, and valid characters\nare /[a-z][0-9]-/.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_GetShipment)

	einride_shipper_v1beta1_ShipmentService_GetShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_GetShipment_Request.Name, "name", "", "Resource name of the shipment to retrieve.")

	// TODO: enum View
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_BatchGetShipments)

	einride_shipper_v1beta1_ShipmentService_BatchGetShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_BatchGetShipments_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all shipments being retrieved.\nThe parent of all of the shipments specified in `names`\nmust match this field.")

	einride_shipper_v1beta1_ShipmentService_BatchGetShipments.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentService_BatchGetShipments_Request.Names, "names", []string{}, "Resource names of the shipments to retrieve.\nA maximum of 1000 shipments can be retrieved in a batch.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_SearchShipments)

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all shipments being searched for.")

	// TODO: list: SiteFilters message

	// TODO: list: ShipmentStatuses enum

	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupEarliestTime.Seconds, "pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupEarliestTime.Nanos, "pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupLatestTime.Seconds, "pickupLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PickupLatestTime.Nanos, "pickupLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryEarliestTime.Seconds, "deliveryEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryEarliestTime.Nanos, "deliveryEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryLatestTime.Seconds, "deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.DeliveryLatestTime.Nanos, "deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.ReferenceIdQuery, "referenceIdQuery", "", "Search for shipments with reference IDs that match the provided query.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().BoolVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.ShowDeleted, "showDeleted", false, "Flag indicating whether to include deleted results.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.Vehicles, "vehicles", []string{}, "Resource names of connected vehicles.\nIf not empty, search for shipments connected to one of these vehicles.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.ShipmentIdPrefix, "shipmentIdPrefix", "", "Search for all shipments with the given shipment ID prefix.\nThe shipment ID is the last segment of the shipment resource name.\nMust be >=3 characters.")

	einride_shipper_v1beta1_ShipmentService_SearchShipments.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_SearchShipments_Request.CarrierReferenceId, "carrierReferenceId", "", "Search for all shipments with a specific carrier reference id.\nIf not set, then no filtering on carrier reference id is done.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UpdateShipment)

	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment = new(v1beta1.Shipment)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.Name, "shipment.name", "", "The resource name of the shipment.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.Etag, "shipment.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the\nclient has an up-to-date value before proceeding.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.IntegrationFile, "shipment.integrationFile", "", "Resource name of the integration file the shipment was ingested from, if the shipment was ingested.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.OriginSite, "shipment.originSite", "", "Resource name of the site from which the shipment originates.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DestinationSite, "shipment.destinationSite", "", "Resource name of the site to which the shipment is destined.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.ReferenceId, "shipment.referenceId", "", "The reference ID for this shipment from e.g. external ERP systems.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().BoolVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.Editable, "shipment.editable", false, "Flag indicating if the shipment should be editable.\nShipments may be uneditable when the master copy is stored in an external\nERP system.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupEarliestTime.Seconds, "shipment.pickupEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupEarliestTime.Nanos, "shipment.pickupEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupLatestTime.Seconds, "shipment.pickupLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupLatestTime.Nanos, "shipment.pickupLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryEarliestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryEarliestTime.Seconds, "shipment.deliveryEarliestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryEarliestTime.Nanos, "shipment.deliveryEarliestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryLatestTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryLatestTime.Seconds, "shipment.deliveryLatestTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryLatestTime.Nanos, "shipment.deliveryLatestTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: enum ShipmentStatus
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.Vehicle, "shipment.vehicle", "", "Resource name of the vehicle performing the delivery.\nEmpty for untracked vehicles.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.PickupInstructions, "shipment.pickupInstructions", "", "Instructions on the pickup, for example which trailer to pick up.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.DeliveryInstructions, "shipment.deliveryInstructions", "", "Instructions on delivery, which could say things like park the trailer at space 123 at lot B.")
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.Shipment.CarrierReferenceId, "shipment.carrierReferenceId", "", "The carrier reference ID for this shipment from. Unique\nidentifier carrier to use, as listed in the Freight Mobility\nPlatform Integration Specification.")
	// TODO: list: LineItems message

	einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_ShipmentService_UpdateShipment.Flags().StringSliceVar(&einride_shipper_v1beta1_ShipmentService_UpdateShipment_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeleteShipment)

	einride_shipper_v1beta1_ShipmentService_DeleteShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_DeleteShipment_Request.Name, "name", "", "Resource name of the shipment to delete.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UndeleteShipment)

	einride_shipper_v1beta1_ShipmentService_UndeleteShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UndeleteShipment_Request.Name, "name", "", "Resource name of the shipment to undelete.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ReferenceShipment)

	einride_shipper_v1beta1_ShipmentService_ReferenceShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ReferenceShipment_Request.Parent, "parent", "", "Resource name of the shipper to look up the shipment for.")

	einride_shipper_v1beta1_ShipmentService_ReferenceShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ReferenceShipment_Request.ReferenceId, "referenceId", "", "Reference ID of the shipment to retrieve.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_PlanShipment)

	einride_shipper_v1beta1_ShipmentService_PlanShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_PlanShipment_Request.Name, "name", "", "Resource name of the shipment to plan.")

	einride_shipper_v1beta1_ShipmentService_PlanShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_PlanShipment_Request.Vehicle, "vehicle", "", "Resource name of the vehicle to plan the shipment on.")

	einride_shipper_v1beta1_ShipmentService_PlanShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_PlanShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnplanShipment)

	einride_shipper_v1beta1_ShipmentService_UnplanShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UnplanShipment_Request.Name, "name", "", "Resource name of the shipment to plan.")

	einride_shipper_v1beta1_ShipmentService_UnplanShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UnplanShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipment)

	einride_shipper_v1beta1_ShipmentService_LoadShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_LoadShipment_Request.Name, "name", "", "Resource name of the shipment to load.")

	einride_shipper_v1beta1_ShipmentService_LoadShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_LoadShipment_Request.Vehicle, "vehicle", "", "Resource name of the vehicle to load the shipment into.")

	einride_shipper_v1beta1_ShipmentService_LoadShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_LoadShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle)

	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle_Request.Name, "name", "", "Resource name of the shipment to load.")

	einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_LoadShipmentUnknownVehicle_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_TransitShipment)

	einride_shipper_v1beta1_ShipmentService_TransitShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_TransitShipment_Request.Name, "name", "", "Resource name of the shipment to set in transit.")

	einride_shipper_v1beta1_ShipmentService_TransitShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_TransitShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ArriveShipment)

	einride_shipper_v1beta1_ShipmentService_ArriveShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ArriveShipment_Request.Name, "name", "", "Resource name of the shipment to set as arrived.")

	einride_shipper_v1beta1_ShipmentService_ArriveShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ArriveShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnloadShipment)

	einride_shipper_v1beta1_ShipmentService_UnloadShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UnloadShipment_Request.Name, "name", "", "Resource name of the shipment to unload.")

	einride_shipper_v1beta1_ShipmentService_UnloadShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UnloadShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_DeliverShipment)

	einride_shipper_v1beta1_ShipmentService_DeliverShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_DeliverShipment_Request.Name, "name", "", "Resource name of the shipment to deliver.")

	einride_shipper_v1beta1_ShipmentService_DeliverShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_DeliverShipment_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally)

	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally_Request.Name, "name", "", "Resource name of the shipment to deliver.")

	einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_HandleShipmentExternally_Request.User, "user", "", "Resource name of the user initiating the request.\nIntended to be used when background jobs are initiating the request.\nIf a user is set in the calling context, it will take precedence over this field.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_RejectShipment)

	einride_shipper_v1beta1_ShipmentService_RejectShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_RejectShipment_Request.Name, "name", "", "Resource name of the shipment to reject.")

	einride_shipper_v1beta1_ShipmentService_RejectShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_RejectShipment_Request.Comment, "comment", "", "Comment on why the shipment was rejected.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_UnrejectShipment)

	einride_shipper_v1beta1_ShipmentService_UnrejectShipment.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_UnrejectShipment_Request.Name, "name", "", "Resource name of the shipment to un-reject.")
	einride_shipper_v1beta1_ShipmentService.AddCommand(einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent)

	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message = new(v1beta1.PubsubMessage)
	// TODO: bytes Data
	// TODO: map: Attributesstring message
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message.MessageId, "message.messageId", "", "ID of this message, assigned by the server when the message is published.\nGuaranteed to be unique within the topic. This value may be read by a\nsubscriber that receives a `PubsubMessage` via a `Pull` call or a push\ndelivery. It must not be populated by the publisher in a `Publish` call.")
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message.PublishTime = new(timestamppb.Timestamp)
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent.Flags().Int64Var(&einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message.PublishTime.Seconds, "message.publishTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message.PublishTime.Nanos, "message.publishTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Message.OrderingKey, "message.orderingKey", "", "If non-empty, identifies related messages for which publish order should be\nrespected. If a `Subscription` has `enable_message_ordering` set to `true`,\nmessages published with the same non-empty `ordering_key` value will be\ndelivered to subscribers in the order in which they are received by the\nPub/Sub system. All `PubsubMessage`s published in a given `PublishRequest`\nmust specify the same `ordering_key` value.")

	einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentService_ConsumeShipmentChangedEvent_Request.Subscription, "subscription", "", "Subscription")
}
