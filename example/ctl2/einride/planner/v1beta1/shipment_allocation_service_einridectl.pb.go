package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_ShipmentAllocationService = &cobra.Command{
	Use: "einride.planner.v1beta1.ShipmentAllocationService",
}

var einride_planner_v1beta1_GetShipmentAllocationSettingsRequest v1beta1.GetShipmentAllocationSettingsRequest
var einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings = &cobra.Command{
	Use: "GetShipmentAllocationSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ShipmentAllocationService.GetShipmentAllocationSettings")
		return nil
	},
}

var einride_planner_v1beta1_UpdateShipmentAllocationSettingsRequest v1beta1.UpdateShipmentAllocationSettingsRequest
var einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings = &cobra.Command{
	Use: "UpdateShipmentAllocationSettings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ShipmentAllocationService.UpdateShipmentAllocationSettings")
		return nil
	},
}

var einride_planner_v1beta1_AllocateShipmentsRequest v1beta1.AllocateShipmentsRequest
var einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments = &cobra.Command{
	Use: "AllocateShipments",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ShipmentAllocationService.AllocateShipments")
		return nil
	},
}

var einride_planner_v1beta1_PublishShipmentAllocationStatusesRequest v1beta1.PublishShipmentAllocationStatusesRequest
var einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses = &cobra.Command{
	Use: "PublishShipmentAllocationStatuses",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ShipmentAllocationService.PublishShipmentAllocationStatuses")
		return nil
	},
}

var einride_planner_v1beta1_AllocateShipmentsAutomaticRequest v1beta1.AllocateShipmentsAutomaticRequest
var einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic = &cobra.Command{
	Use: "AllocateShipmentsAutomatic",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ShipmentAllocationService.AllocateShipmentsAutomatic")
		return nil
	},
}

func AddShipmentAllocationServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_ShipmentAllocationService)
}

func init() {
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_GetShipmentAllocationSettings)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic)
}
