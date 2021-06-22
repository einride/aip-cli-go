package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planning.v1.ScheduleBookingService.
var (
	einride_planning_v1_ScheduleBookingServiceClient v1.ScheduleBookingServiceClient
	einride_planning_v1_ScheduleBookingService       = &cobra.Command{
		Use: "einride.planning.v1.ScheduleBookingService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_ScheduleBookingServiceClient = v1.NewScheduleBookingServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.CreateScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request v1.CreateScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking         = &cobra.Command{
		Use: "CreateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.CreateScheduleBooking")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.GetScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_GetScheduleBooking_Request v1.GetScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_GetScheduleBooking         = &cobra.Command{
		Use: "GetScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.GetScheduleBooking")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookings.
var (
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings_Request v1.BatchGetScheduleBookingsRequest
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings         = &cobra.Command{
		Use: "BatchGetScheduleBookings",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookings")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UpdateScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request v1.UpdateScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking         = &cobra.Command{
		Use: "UpdateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.UpdateScheduleBooking")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.GetScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance_Request v1.GetScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance         = &cobra.Command{
		Use: "GetScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.GetScheduleBookingInstance")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookingInstances.
var (
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances_Request v1.BatchGetScheduleBookingInstancesRequest
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances         = &cobra.Command{
		Use: "BatchGetScheduleBookingInstances",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookingInstances")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.ListScheduleBookingInstances.
var (
	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request v1.ListScheduleBookingInstancesRequest
	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances         = &cobra.Command{
		Use: "ListScheduleBookingInstances",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.ListScheduleBookingInstances")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UpdateScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request v1.UpdateScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance         = &cobra.Command{
		Use: "UpdateScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.UpdateScheduleBookingInstance")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.ReconcileScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance_Request v1.ReconcileScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance         = &cobra.Command{
		Use: "ReconcileScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.ReconcileScheduleBookingInstance")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.DeleteScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance_Request v1.DeleteScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance         = &cobra.Command{
		Use: "DeleteScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.DeleteScheduleBookingInstance")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UndeleteScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance_Request v1.UndeleteScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance         = &cobra.Command{
		Use: "UndeleteScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleBookingService.UndeleteScheduleBookingInstance")
			return nil
		},
	}
)

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
