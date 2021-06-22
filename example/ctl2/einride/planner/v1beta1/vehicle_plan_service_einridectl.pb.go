package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.VehiclePlanService.
var (
	einride_planner_v1beta1_VehiclePlanServiceClient v1beta1.VehiclePlanServiceClient
	einride_planner_v1beta1_VehiclePlanService       = &cobra.Command{
		Use: "einride.planner.v1beta1.VehiclePlanService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_VehiclePlanServiceClient = v1beta1.NewVehiclePlanServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CreateVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan_Request v1beta1.CreateVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan         = &cobra.Command{
		Use: "CreateVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.CreateVehiclePlan")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.GetVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan_Request v1beta1.GetVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan         = &cobra.Command{
		Use: "GetVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.GetVehiclePlan")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchGetVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans_Request v1beta1.BatchGetVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans         = &cobra.Command{
		Use: "BatchGetVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.BatchGetVehiclePlans")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ShipperSearchVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks_Request v1beta1.ShipperSearchVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks         = &cobra.Command{
		Use: "ShipperSearchVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.ShipperSearchVehicleTasks")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.SearchVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans_Request v1beta1.SearchVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans         = &cobra.Command{
		Use: "SearchVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.SearchVehiclePlans")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ListVehiclePlans.
var (
	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans_Request v1beta1.ListVehiclePlansRequest
	einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans         = &cobra.Command{
		Use: "ListVehiclePlans",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.ListVehiclePlans")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeleteVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan_Request v1beta1.DeleteVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan         = &cobra.Command{
		Use: "DeleteVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.DeleteVehiclePlan")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UndeleteVehiclePlan.
var (
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan_Request v1beta1.UndeleteVehiclePlanRequest
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan         = &cobra.Command{
		Use: "UndeleteVehiclePlan",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.UndeleteVehiclePlan")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CreateVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask_Request v1beta1.CreateVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask         = &cobra.Command{
		Use: "CreateVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.CreateVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchCreateVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks_Request v1beta1.BatchCreateVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks         = &cobra.Command{
		Use: "BatchCreateVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.BatchCreateVehicleTasks")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UpdateVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask_Request v1beta1.UpdateVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask         = &cobra.Command{
		Use: "UpdateVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.UpdateVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ListVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks_Request v1beta1.ListVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks         = &cobra.Command{
		Use: "ListVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.ListVehicleTasks")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.GetVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_GetVehicleTask_Request v1beta1.GetVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_GetVehicleTask         = &cobra.Command{
		Use: "GetVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.GetVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.BatchGetVehicleTasks.
var (
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks_Request v1beta1.BatchGetVehicleTasksRequest
	einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks         = &cobra.Command{
		Use: "BatchGetVehicleTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.BatchGetVehicleTasks")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask_Request v1beta1.DeleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask         = &cobra.Command{
		Use: "DeleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.DeleteVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.UndeleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask_Request v1beta1.UndeleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask         = &cobra.Command{
		Use: "UndeleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.UndeleteVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.StartVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask_Request v1beta1.StartVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_StartVehicleTask         = &cobra.Command{
		Use: "StartVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.StartVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.CompleteVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask_Request v1beta1.CompleteVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask         = &cobra.Command{
		Use: "CompleteVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.CompleteVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.AssignShipmentToPickupShipmentVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask_Request v1beta1.AssignShipmentToPickupShipmentVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask         = &cobra.Command{
		Use: "AssignShipmentToPickupShipmentVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.AssignShipmentToPickupShipmentVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.DeassignShipmentFromPickupShipmentVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask_Request v1beta1.DeassignShipmentFromPickupShipmentVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask         = &cobra.Command{
		Use: "DeassignShipmentFromPickupShipmentVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.DeassignShipmentFromPickupShipmentVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.RejectShipmentInPickupVehicleTask.
var (
	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask_Request v1beta1.RejectShipmentInPickupVehicleTaskRequest
	einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask         = &cobra.Command{
		Use: "RejectShipmentInPickupVehicleTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.RejectShipmentInPickupVehicleTask")
			return nil
		},
	}
)

// einride.planner.v1beta1.VehiclePlanService.ApproveDocking.
var (
	einride_planner_v1beta1_VehiclePlanService_ApproveDocking_Request v1beta1.ApproveDockingRequest
	einride_planner_v1beta1_VehiclePlanService_ApproveDocking         = &cobra.Command{
		Use: "ApproveDocking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.VehiclePlanService.ApproveDocking")
			return nil
		},
	}
)

func AddVehiclePlanServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_VehiclePlanService)
}

func init() {
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_GetVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_StartVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask)
	einride_planner_v1beta1_VehiclePlanService.AddCommand(einride_planner_v1beta1_VehiclePlanService_ApproveDocking)
}
