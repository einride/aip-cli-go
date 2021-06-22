package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_DefaultScheduleService = &cobra.Command{
	Use: "einride.planner.v1beta1.DefaultScheduleService",
}

var einride_planner_v1beta1_GetDefaultScheduleRequest v1beta1.GetDefaultScheduleRequest
var einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule = &cobra.Command{
	Use: "GetDefaultSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.DefaultScheduleService.GetDefaultSchedule")
		return nil
	},
}

var einride_planner_v1beta1_UpdateDefaultScheduleRequest v1beta1.UpdateDefaultScheduleRequest
var einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule = &cobra.Command{
	Use: "UpdateDefaultSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.DefaultScheduleService.UpdateDefaultSchedule")
		return nil
	},
}

var einride_planner_v1beta1_BookDefaultScheduleRequest v1beta1.BookDefaultScheduleRequest
var einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule = &cobra.Command{
	Use: "BookDefaultSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedule")
		return nil
	},
}

var einride_planner_v1beta1_BookDefaultSchedulesRequest v1beta1.BookDefaultSchedulesRequest
var einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules = &cobra.Command{
	Use: "BookDefaultSchedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.DefaultScheduleService.BookDefaultSchedules")
		return nil
	},
}

func AddDefaultScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_DefaultScheduleService)
}

func init() {
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_GetDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_UpdateDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedule)
	einride_planner_v1beta1_DefaultScheduleService.AddCommand(einride_planner_v1beta1_DefaultScheduleService_BookDefaultSchedules)
}
