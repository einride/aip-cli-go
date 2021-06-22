package plannerv1beta1

import (
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planner.v1beta1.DefaultScheduleService.
var (
	einride_planner_v1beta1_DefaultScheduleServiceClient v1beta1.DefaultScheduleServiceClient
	einride_planner_v1beta1_DefaultScheduleService       = &cobra.Command{
		Use: "einride.planner.v1beta1.DefaultScheduleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_DefaultScheduleServiceClient = v1beta1.NewDefaultScheduleServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.GetDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule_Request v1beta1.GetDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule         = &cobra.Command{
		Use: "GetDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.GetDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.UpdateDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule_Request v1beta1.UpdateDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule         = &cobra.Command{
		Use: "UpdateDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.UpdateDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedule.
var (
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule_Request v1beta1.BookDefaultScheduleRequest
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule         = &cobra.Command{
		Use: "BookDefaultSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedule")
			return nil
		},
	}
)

// einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedules.
var (
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules_Request v1beta1.BookDefaultSchedulesRequest
	einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules         = &cobra.Command{
		Use: "BookDefaultSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedules")
			return nil
		},
	}
)

func AddDefaultScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_DefaultScheduleService)
}

func init() {
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules)
}
