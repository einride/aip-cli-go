package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_ScheduleBookingService = &cobra.Command{
	Use: "einride.planner.v1beta1.ScheduleBookingService",
}

var einride_planner_v1beta1_CreateScheduleBookingRequest v1beta1.CreateScheduleBookingRequest
var einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking = &cobra.Command{
	Use: "CreateScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ScheduleBookingService.CreateScheduleBooking")
		return nil
	},
}

var einride_planner_v1beta1_UpdateScheduleBookingRequest v1beta1.UpdateScheduleBookingRequest
var einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking = &cobra.Command{
	Use: "UpdateScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ScheduleBookingService.UpdateScheduleBooking")
		return nil
	},
}

var einride_planner_v1beta1_DeleteScheduleBookingRequest v1beta1.DeleteScheduleBookingRequest
var einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking = &cobra.Command{
	Use: "DeleteScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ScheduleBookingService.DeleteScheduleBooking")
		return nil
	},
}

var einride_planner_v1beta1_GetScheduleBookingRequest v1beta1.GetScheduleBookingRequest
var einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking = &cobra.Command{
	Use: "GetScheduleBooking",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ScheduleBookingService.GetScheduleBooking")
		return nil
	},
}

var einride_planner_v1beta1_SyncScheduleBookingsRequest v1beta1.SyncScheduleBookingsRequest
var einride_planner_v1beta1_ScheduleBookingService_SyncScheduleBookings = &cobra.Command{
	Use: "SyncScheduleBookings",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.ScheduleBookingService.SyncScheduleBookings")
		return nil
	},
}

func AddScheduleBookingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_ScheduleBookingService)
}

func init() {
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_SyncScheduleBookings)
}
