package planningv1

import (
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planning_v1_ScheduleService = &cobra.Command{
	Use: "einride.planning.v1.ScheduleService",
}

var einride_planning_v1_CreateScheduleRequest v1.CreateScheduleRequest
var einride_planning_v1_ScheduleService_CreateSchedule = &cobra.Command{
	Use: "CreateSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.CreateSchedule")
		return nil
	},
}

var einride_planning_v1_GetScheduleRequest v1.GetScheduleRequest
var einride_planning_v1_ScheduleService_GetSchedule = &cobra.Command{
	Use: "GetSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.GetSchedule")
		return nil
	},
}

var einride_planning_v1_BatchGetSchedulesRequest v1.BatchGetSchedulesRequest
var einride_planning_v1_ScheduleService_BatchGetSchedules = &cobra.Command{
	Use: "BatchGetSchedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.BatchGetSchedules")
		return nil
	},
}

var einride_planning_v1_ListSchedulesRequest v1.ListSchedulesRequest
var einride_planning_v1_ScheduleService_ListSchedules = &cobra.Command{
	Use: "ListSchedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.ListSchedules")
		return nil
	},
}

var einride_planning_v1_UpdateScheduleRequest v1.UpdateScheduleRequest
var einride_planning_v1_ScheduleService_UpdateSchedule = &cobra.Command{
	Use: "UpdateSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.UpdateSchedule")
		return nil
	},
}

var einride_planning_v1_PublishScheduleRequest v1.PublishScheduleRequest
var einride_planning_v1_ScheduleService_PublishSchedule = &cobra.Command{
	Use: "PublishSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.PublishSchedule")
		return nil
	},
}

var einride_planning_v1_DeleteScheduleRequest v1.DeleteScheduleRequest
var einride_planning_v1_ScheduleService_DeleteSchedule = &cobra.Command{
	Use: "DeleteSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.DeleteSchedule")
		return nil
	},
}

var einride_planning_v1_UndeleteScheduleRequest v1.UndeleteScheduleRequest
var einride_planning_v1_ScheduleService_UndeleteSchedule = &cobra.Command{
	Use: "UndeleteSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.ScheduleService.UndeleteSchedule")
		return nil
	},
}

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
