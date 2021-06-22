package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
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
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_UpdateShipmentAllocationSettings)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipments)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_PublishShipmentAllocationStatuses)
	einride_planner_v1beta1_ShipmentAllocationService.AddCommand(einride_planner_v1beta1_ShipmentAllocationService_AllocateShipmentsAutomatic)
}
