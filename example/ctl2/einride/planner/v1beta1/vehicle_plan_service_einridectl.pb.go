package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_VehiclePlanService = &cobra.Command{
	Use: "einride.planner.v1beta1.VehiclePlanService",
}

var einride_planner_v1beta1_CreateVehiclePlanRequest v1beta1.CreateVehiclePlanRequest
var einride_planner_v1beta1_VehiclePlanService_CreateVehiclePlan = &cobra.Command{
	Use: "CreateVehiclePlan",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.CreateVehiclePlan")
		return nil
	},
}

var einride_planner_v1beta1_GetVehiclePlanRequest v1beta1.GetVehiclePlanRequest
var einride_planner_v1beta1_VehiclePlanService_GetVehiclePlan = &cobra.Command{
	Use: "GetVehiclePlan",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.GetVehiclePlan")
		return nil
	},
}

var einride_planner_v1beta1_BatchGetVehiclePlansRequest v1beta1.BatchGetVehiclePlansRequest
var einride_planner_v1beta1_VehiclePlanService_BatchGetVehiclePlans = &cobra.Command{
	Use: "BatchGetVehiclePlans",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.BatchGetVehiclePlans")
		return nil
	},
}

var einride_planner_v1beta1_ShipperSearchVehicleTasksRequest v1beta1.ShipperSearchVehicleTasksRequest
var einride_planner_v1beta1_VehiclePlanService_ShipperSearchVehicleTasks = &cobra.Command{
	Use: "ShipperSearchVehicleTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.ShipperSearchVehicleTasks")
		return nil
	},
}

var einride_planner_v1beta1_SearchVehiclePlansRequest v1beta1.SearchVehiclePlansRequest
var einride_planner_v1beta1_VehiclePlanService_SearchVehiclePlans = &cobra.Command{
	Use: "SearchVehiclePlans",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.SearchVehiclePlans")
		return nil
	},
}

var einride_planner_v1beta1_ListVehiclePlansRequest v1beta1.ListVehiclePlansRequest
var einride_planner_v1beta1_VehiclePlanService_ListVehiclePlans = &cobra.Command{
	Use: "ListVehiclePlans",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.ListVehiclePlans")
		return nil
	},
}

var einride_planner_v1beta1_DeleteVehiclePlanRequest v1beta1.DeleteVehiclePlanRequest
var einride_planner_v1beta1_VehiclePlanService_DeleteVehiclePlan = &cobra.Command{
	Use: "DeleteVehiclePlan",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.DeleteVehiclePlan")
		return nil
	},
}

var einride_planner_v1beta1_UndeleteVehiclePlanRequest v1beta1.UndeleteVehiclePlanRequest
var einride_planner_v1beta1_VehiclePlanService_UndeleteVehiclePlan = &cobra.Command{
	Use: "UndeleteVehiclePlan",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.UndeleteVehiclePlan")
		return nil
	},
}

var einride_planner_v1beta1_CreateVehicleTaskRequest v1beta1.CreateVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_CreateVehicleTask = &cobra.Command{
	Use: "CreateVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.CreateVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_BatchCreateVehicleTasksRequest v1beta1.BatchCreateVehicleTasksRequest
var einride_planner_v1beta1_VehiclePlanService_BatchCreateVehicleTasks = &cobra.Command{
	Use: "BatchCreateVehicleTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.BatchCreateVehicleTasks")
		return nil
	},
}

var einride_planner_v1beta1_UpdateVehicleTaskRequest v1beta1.UpdateVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_UpdateVehicleTask = &cobra.Command{
	Use: "UpdateVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.UpdateVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_ListVehicleTasksRequest v1beta1.ListVehicleTasksRequest
var einride_planner_v1beta1_VehiclePlanService_ListVehicleTasks = &cobra.Command{
	Use: "ListVehicleTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.ListVehicleTasks")
		return nil
	},
}

var einride_planner_v1beta1_GetVehicleTaskRequest v1beta1.GetVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_GetVehicleTask = &cobra.Command{
	Use: "GetVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.GetVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_BatchGetVehicleTasksRequest v1beta1.BatchGetVehicleTasksRequest
var einride_planner_v1beta1_VehiclePlanService_BatchGetVehicleTasks = &cobra.Command{
	Use: "BatchGetVehicleTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.BatchGetVehicleTasks")
		return nil
	},
}

var einride_planner_v1beta1_DeleteVehicleTaskRequest v1beta1.DeleteVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_DeleteVehicleTask = &cobra.Command{
	Use: "DeleteVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.DeleteVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_UndeleteVehicleTaskRequest v1beta1.UndeleteVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_UndeleteVehicleTask = &cobra.Command{
	Use: "UndeleteVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.UndeleteVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_StartVehicleTaskRequest v1beta1.StartVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_StartVehicleTask = &cobra.Command{
	Use: "StartVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.StartVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_CompleteVehicleTaskRequest v1beta1.CompleteVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_CompleteVehicleTask = &cobra.Command{
	Use: "CompleteVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.CompleteVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_AssignShipmentToPickupShipmentVehicleTaskRequest v1beta1.AssignShipmentToPickupShipmentVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_AssignShipmentToPickupShipmentVehicleTask = &cobra.Command{
	Use: "AssignShipmentToPickupShipmentVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.AssignShipmentToPickupShipmentVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_DeassignShipmentFromPickupShipmentVehicleTaskRequest v1beta1.DeassignShipmentFromPickupShipmentVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_DeassignShipmentFromPickupShipmentVehicleTask = &cobra.Command{
	Use: "DeassignShipmentFromPickupShipmentVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.DeassignShipmentFromPickupShipmentVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_RejectShipmentInPickupVehicleTaskRequest v1beta1.RejectShipmentInPickupVehicleTaskRequest
var einride_planner_v1beta1_VehiclePlanService_RejectShipmentInPickupVehicleTask = &cobra.Command{
	Use: "RejectShipmentInPickupVehicleTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.RejectShipmentInPickupVehicleTask")
		return nil
	},
}

var einride_planner_v1beta1_ApproveDockingRequest v1beta1.ApproveDockingRequest
var einride_planner_v1beta1_VehiclePlanService_ApproveDocking = &cobra.Command{
	Use: "ApproveDocking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.VehiclePlanService.ApproveDocking")
		return nil
	},
}

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
