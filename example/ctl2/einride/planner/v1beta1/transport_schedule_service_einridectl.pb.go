package plannerv1beta1

import (
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planner_v1beta1_TransportScheduleService = &cobra.Command{
	Use: "einride.planner.v1beta1.TransportScheduleService",
}

var einride_planner_v1beta1_CreateTransportScheduleRequest v1beta1.CreateTransportScheduleRequest
var einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule = &cobra.Command{
	Use: "CreateTransportSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.CreateTransportSchedule")
		return nil
	},
}

var einride_planner_v1beta1_GetTransportScheduleRequest v1beta1.GetTransportScheduleRequest
var einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule = &cobra.Command{
	Use: "GetTransportSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.GetTransportSchedule")
		return nil
	},
}

var einride_planner_v1beta1_UpdateTransportScheduleRequest v1beta1.UpdateTransportScheduleRequest
var einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule = &cobra.Command{
	Use: "UpdateTransportSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.UpdateTransportSchedule")
		return nil
	},
}

var einride_planner_v1beta1_ListTransportSchedulesRequest v1beta1.ListTransportSchedulesRequest
var einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules = &cobra.Command{
	Use: "ListTransportSchedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.ListTransportSchedules")
		return nil
	},
}

var einride_planner_v1beta1_BatchGetTransportSchedulesRequest v1beta1.BatchGetTransportSchedulesRequest
var einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules = &cobra.Command{
	Use: "BatchGetTransportSchedules",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.BatchGetTransportSchedules")
		return nil
	},
}

var einride_planner_v1beta1_DeleteTransportScheduleRequest v1beta1.DeleteTransportScheduleRequest
var einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule = &cobra.Command{
	Use: "DeleteTransportSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.DeleteTransportSchedule")
		return nil
	},
}

var einride_planner_v1beta1_UndeleteTransportScheduleRequest v1beta1.UndeleteTransportScheduleRequest
var einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule = &cobra.Command{
	Use: "UndeleteTransportSchedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planner.v1beta1.TransportScheduleService.UndeleteTransportSchedule")
		return nil
	},
}

func AddTransportScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportScheduleService)
}

func init() {
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule)
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule)
}
