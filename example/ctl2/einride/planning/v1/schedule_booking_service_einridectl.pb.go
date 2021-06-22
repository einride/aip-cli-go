package planningv1

import (
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planning_v1_ScheduleBookingService = &cobra.Command{
	Use: "einride.planning.v1.ScheduleBookingService",
}

var einride_planning_v1_CreateScheduleBookingRequest v1.CreateScheduleBookingRequest
var einride_planning_v1_ScheduleBookingService_CreateScheduleBooking = &cobra.Command{
	Use: "CreateScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.CreateScheduleBooking")
		return nil
	},
}

var einride_planning_v1_GetScheduleBookingRequest v1.GetScheduleBookingRequest
var einride_planning_v1_ScheduleBookingService_GetScheduleBooking = &cobra.Command{
	Use: "GetScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.GetScheduleBooking")
		return nil
	},
}

var einride_planning_v1_BatchGetScheduleBookingsRequest v1.BatchGetScheduleBookingsRequest
var einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings = &cobra.Command{
	Use: "BatchGetScheduleBookings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookings")
		return nil
	},
}

var einride_planning_v1_UpdateScheduleBookingRequest v1.UpdateScheduleBookingRequest
var einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking = &cobra.Command{
	Use: "UpdateScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.UpdateScheduleBooking")
		return nil
	},
}

var einride_planning_v1_GetScheduleBookingInstanceRequest v1.GetScheduleBookingInstanceRequest
var einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance = &cobra.Command{
	Use: "GetScheduleBookingInstance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.GetScheduleBookingInstance")
		return nil
	},
}

var einride_planning_v1_BatchGetScheduleBookingInstancesRequest v1.BatchGetScheduleBookingInstancesRequest
var einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances = &cobra.Command{
	Use: "BatchGetScheduleBookingInstances",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookingInstances")
		return nil
	},
}

var einride_planning_v1_ListScheduleBookingInstancesRequest v1.ListScheduleBookingInstancesRequest
var einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances = &cobra.Command{
	Use: "ListScheduleBookingInstances",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.ListScheduleBookingInstances")
		return nil
	},
}

var einride_planning_v1_UpdateScheduleBookingInstanceRequest v1.UpdateScheduleBookingInstanceRequest
var einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance = &cobra.Command{
	Use: "UpdateScheduleBookingInstance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.UpdateScheduleBookingInstance")
		return nil
	},
}

var einride_planning_v1_ReconcileScheduleBookingInstanceRequest v1.ReconcileScheduleBookingInstanceRequest
var einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance = &cobra.Command{
	Use: "ReconcileScheduleBookingInstance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.ReconcileScheduleBookingInstance")
		return nil
	},
}

var einride_planning_v1_DeleteScheduleBookingInstanceRequest v1.DeleteScheduleBookingInstanceRequest
var einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance = &cobra.Command{
	Use: "DeleteScheduleBookingInstance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.DeleteScheduleBookingInstance")
		return nil
	},
}

var einride_planning_v1_UndeleteScheduleBookingInstanceRequest v1.UndeleteScheduleBookingInstanceRequest
var einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance = &cobra.Command{
	Use: "UndeleteScheduleBookingInstance",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleBookingService.UndeleteScheduleBookingInstance")
		return nil
	},
}

func AddScheduleBookingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_ScheduleBookingService)
}

func init() {
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_CreateScheduleBooking)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_GetScheduleBooking)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance)
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance)
}
