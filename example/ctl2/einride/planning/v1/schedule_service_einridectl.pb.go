package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planning.v1.ScheduleService.
var (
	einride_planning_v1_ScheduleServiceClient v1.ScheduleServiceClient
	einride_planning_v1_ScheduleService       = &cobra.Command{
		Use: "einride.planning.v1.ScheduleService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_ScheduleServiceClient = v1.NewScheduleServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.CreateSchedule.
var (
	einride_planning_v1_ScheduleService_CreateSchedule_Request v1.CreateScheduleRequest
	einride_planning_v1_ScheduleService_CreateSchedule         = &cobra.Command{
		Use: "CreateSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.CreateSchedule")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.GetSchedule.
var (
	einride_planning_v1_ScheduleService_GetSchedule_Request v1.GetScheduleRequest
	einride_planning_v1_ScheduleService_GetSchedule         = &cobra.Command{
		Use: "GetSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.GetSchedule")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.BatchGetSchedules.
var (
	einride_planning_v1_ScheduleService_BatchGetSchedules_Request v1.BatchGetSchedulesRequest
	einride_planning_v1_ScheduleService_BatchGetSchedules         = &cobra.Command{
		Use: "BatchGetSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.BatchGetSchedules")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.ListSchedules.
var (
	einride_planning_v1_ScheduleService_ListSchedules_Request v1.ListSchedulesRequest
	einride_planning_v1_ScheduleService_ListSchedules         = &cobra.Command{
		Use: "ListSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.ListSchedules")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.UpdateSchedule.
var (
	einride_planning_v1_ScheduleService_UpdateSchedule_Request v1.UpdateScheduleRequest
	einride_planning_v1_ScheduleService_UpdateSchedule         = &cobra.Command{
		Use: "UpdateSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.UpdateSchedule")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.PublishSchedule.
var (
	einride_planning_v1_ScheduleService_PublishSchedule_Request v1.PublishScheduleRequest
	einride_planning_v1_ScheduleService_PublishSchedule         = &cobra.Command{
		Use: "PublishSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.PublishSchedule")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.DeleteSchedule.
var (
	einride_planning_v1_ScheduleService_DeleteSchedule_Request v1.DeleteScheduleRequest
	einride_planning_v1_ScheduleService_DeleteSchedule         = &cobra.Command{
		Use: "DeleteSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.DeleteSchedule")
			return nil
		},
	}
)

// einride.planning.v1.ScheduleService.UndeleteSchedule.
var (
	einride_planning_v1_ScheduleService_UndeleteSchedule_Request v1.UndeleteScheduleRequest
	einride_planning_v1_ScheduleService_UndeleteSchedule         = &cobra.Command{
		Use: "UndeleteSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.ScheduleService.UndeleteSchedule")
			return nil
		},
	}
)

func AddScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_ScheduleService)
}

func init() {
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_CreateSchedule)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_GetSchedule)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_BatchGetSchedules)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_ListSchedules)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_UpdateSchedule)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_PublishSchedule)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_DeleteSchedule)
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_UndeleteSchedule)
}
