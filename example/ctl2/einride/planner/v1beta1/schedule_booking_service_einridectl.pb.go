package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.Parent, "parent", "", "Resource name of the parent of the booking.")

	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking = new(v1beta1.ScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Name, "scheduleBooking.name", "", "The resource name of the schedule booking.\nThe date part must be a valid ISO 8601 date.\nFor example: `\"transportInstallations/transportInstallation/scheduleBookings/2020-03-02\"`.")
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Etag, "scheduleBooking.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Schedule, "scheduleBooking.schedule", "", "Resource name of the schedule to execute.")
	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().BoolVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.IsDefaultSchedule, "scheduleBooking.isDefaultSchedule", false, "Is this the default schedule for the day.")

	einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBookingId, "scheduleBookingId", "", "The date to execute the schedule booking on, which will become the final component\nof the schedule booking's resource name.\n\nThis value should be a date in ISO 8601 format YYYY-MM-DD (2020-08-24)")
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking)

	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking = new(v1beta1.ScheduleBooking)
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Name, "scheduleBooking.name", "", "The resource name of the schedule booking.\nThe date part must be a valid ISO 8601 date.\nFor example: `\"transportInstallations/transportInstallation/scheduleBookings/2020-03-02\"`.")
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Etag, "scheduleBooking.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Schedule, "scheduleBooking.schedule", "", "Resource name of the schedule to execute.")
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking.Flags().BoolVar(&einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.IsDefaultSchedule, "scheduleBooking.isDefaultSchedule", false, "Is this the default schedule for the day.")

	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringSliceVar(&einride_planner_v1beta1_ScheduleBookingService_UpdateScheduleBooking_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking)

	einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_DeleteScheduleBooking_Request.Name, "name", "", "Resource name of the schedule booking to delete.")
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking)

	einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking.Flags().StringVar(&einride_planner_v1beta1_ScheduleBookingService_GetScheduleBooking_Request.Name, "name", "", "Resource name of the schedule booking to retrieve.\nPattern: `transportInstallations/{transport_installation}/scheduleBookings/{schedule_booking}`.")
	einride_planner_v1beta1_ScheduleBookingService.AddCommand(einride_planner_v1beta1_ScheduleBookingService_SyncScheduleBookings)
}
