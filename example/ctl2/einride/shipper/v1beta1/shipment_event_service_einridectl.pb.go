package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.shipper.v1beta1.ShipmentEventService.
var (
	einride_shipper_v1beta1_ShipmentEventServiceClient v1beta1.ShipmentEventServiceClient
	einride_shipper_v1beta1_ShipmentEventService       = &cobra.Command{
		Use: "einride.shipper.v1beta1.ShipmentEventService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_ShipmentEventServiceClient = v1beta1.NewShipmentEventServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentEventService.CreateShipmentEvent.
var (
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request v1beta1.CreateShipmentEventRequest
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent         = &cobra.Command{
		Use: "CreateShipmentEvent",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentEventServiceClient.CreateShipmentEvent(cmd.Context(), &einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentEventService.GetShipmentEvent.
var (
	einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent_Request v1beta1.GetShipmentEventRequest
	einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent         = &cobra.Command{
		Use: "GetShipmentEvent",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentEventServiceClient.GetShipmentEvent(cmd.Context(), &einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.ShipmentEventService.ListShipmentEvents.
var (
	einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents_Request v1beta1.ListShipmentEventsRequest
	einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents         = &cobra.Command{
		Use: "ListShipmentEvents",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_ShipmentEventServiceClient.ListShipmentEvents(cmd.Context(), &einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddShipmentEventServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentEventService)
}

func init() {
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent)

	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.Parent, "parent", "", "Resource name of the parent shipment where this shipment event will be created.")

	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent = new(v1beta1.ShipmentEvent)
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent.Name, "shipmentEvent.name", "", "The resource name of the shipment event.")
	// TODO: enum PreviousStatus
	// TODO: enum NewStatus
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent.IntegrationFile, "shipmentEvent.integrationFile", "", "Resource name of the integration file that made the transition, if ingested.")
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent.Vehicle, "shipmentEvent.vehicle", "", "Resource name of the vehicle involved in the event.\nFor example: When unloading, this field will be set.")
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent.Site, "shipmentEvent.site", "", "Resource name of the site involved in the event.\nFor example: When arriving to a site, this field will be set.")
	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEvent.Comment, "shipmentEvent.comment", "", "Comment for the shipment event, for example why a shipment is rejected.")

	einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent_Request.ShipmentEventId, "shipmentEventId", "", "The ID to use for the shipment event, which will become the final component of\nthe shipment event's resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be a valid unix nanosecond timestamp.")
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent)

	einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent_Request.Name, "name", "", "Resource name of the shipment event to retrieve.")
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents)

	einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents_Request.Parent, "parent", "", "Resource name of the parent shipment.")

	einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents.Flags().Int32Var(&einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents.Flags().StringVar(&einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
}
