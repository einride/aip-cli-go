package shipperv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
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
			log.Println("einride.shipper.v1beta1.ShipmentEventService.CreateShipmentEvent")
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
			log.Println("einride.shipper.v1beta1.ShipmentEventService.GetShipmentEvent")
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
			log.Println("einride.shipper.v1beta1.ShipmentEventService.ListShipmentEvents")
			return nil
		},
	}
)

func AddShipmentEventServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentEventService)
}

func init() {
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent)
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent)
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents)
}
