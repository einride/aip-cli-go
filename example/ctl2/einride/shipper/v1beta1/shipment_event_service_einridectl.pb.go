package shipperv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_shipper_v1beta1_ShipmentEventService = &cobra.Command{
	Use: "einride.shipper.v1beta1.ShipmentEventService",
}

var einride_shipper_v1beta1_CreateShipmentEventRequest v1beta1.CreateShipmentEventRequest
var einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent = &cobra.Command{
	Use: "CreateShipmentEvent",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentEventService.CreateShipmentEvent")
		return nil
	},
}

var einride_shipper_v1beta1_GetShipmentEventRequest v1beta1.GetShipmentEventRequest
var einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent = &cobra.Command{
	Use: "GetShipmentEvent",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentEventService.GetShipmentEvent")
		return nil
	},
}

var einride_shipper_v1beta1_ListShipmentEventsRequest v1beta1.ListShipmentEventsRequest
var einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents = &cobra.Command{
	Use: "ListShipmentEvents",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.shipper.v1beta1.ShipmentEventService.ListShipmentEvents")
		return nil
	},
}

func AddShipmentEventServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_ShipmentEventService)
}

func init() {
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_CreateShipmentEvent)
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_GetShipmentEvent)
	einride_shipper_v1beta1_ShipmentEventService.AddCommand(einride_shipper_v1beta1_ShipmentEventService_ListShipmentEvents)
}
