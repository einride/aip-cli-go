package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.ScheduleBookingService.
var (
	einride_planner_v1beta1_ScheduleBookingServiceClient v1beta1.ScheduleBookingServiceClient
	einride_planner_v1beta1_ScheduleBookingService       = &cobra.Command{
		Use: "einride.planner.v1beta1.ScheduleBookingService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_ScheduleBookingServiceClient = v1beta1.NewScheduleBookingServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.ScheduleBookingService.CreateScheduleBooking.
var (
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request v1beta1.CreateScheduleBookingRequest
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking         = &cobra.Command{
		Use: "CreateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ScheduleBookingService.CreateScheduleBooking")
			return nil
		},
	}
)

// einride.planner.v1beta1.ScheduleBookingService.UpdateScheduleBooking.
var (
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request v1beta1.UpdateScheduleBookingRequest
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking         = &cobra.Command{
		Use: "UpdateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ScheduleBookingService.UpdateScheduleBooking")
			return nil
		},
	}
)

// einride.planner.v1beta1.ScheduleBookingService.DeleteScheduleBooking.
var (
	einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking_Request v1beta1.DeleteScheduleBookingRequest
	einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking         = &cobra.Command{
		Use: "DeleteScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ScheduleBookingService.DeleteScheduleBooking")
			return nil
		},
	}
)

// einride.planner.v1beta1.ScheduleBookingService.GetScheduleBooking.
var (
	einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking_Request v1beta1.GetScheduleBookingRequest
	einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking         = &cobra.Command{
		Use: "GetScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ScheduleBookingService.GetScheduleBooking")
			return nil
		},
	}
)

// einride.planner.v1beta1.ScheduleBookingService.SyncScheduleBookings.
var (
	einride_planner_v1beta1_ScheduleBookingService_SyncScheduleBookings_Request v1beta1.SyncScheduleBookingsRequest
	einride_planner_v1beta1_ScheduleBookingService_SyncScheduleBookings         = &cobra.Command{
		Use: "SyncScheduleBookings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.ScheduleBookingService.SyncScheduleBookings")
			return nil
		},
	}
)

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
