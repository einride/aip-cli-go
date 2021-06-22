package planningv1

import (
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planning_v1_AllocateShipmentsService = &cobra.Command{
	Use: "einride.planning.v1.AllocateShipmentsService",
}

var einride_planning_v1_CreateAllocateShipmentsJobRequest v1.CreateAllocateShipmentsJobRequest
var einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob = &cobra.Command{
	Use: "CreateAllocateShipmentsJob",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.CreateAllocateShipmentsJob")
		return nil
	},
}

var einride_planning_v1_GetAllocateShipmentsJobRequest v1.GetAllocateShipmentsJobRequest
var einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob = &cobra.Command{
	Use: "GetAllocateShipmentsJob",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.GetAllocateShipmentsJob")
		return nil
	},
}

var einride_planning_v1_UpdateAllocateShipmentsJobRequest v1.UpdateAllocateShipmentsJobRequest
var einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob = &cobra.Command{
	Use: "UpdateAllocateShipmentsJob",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.UpdateAllocateShipmentsJob")
		return nil
	},
}

var einride_planning_v1_ListAllocateShipmentsJobsRequest v1.ListAllocateShipmentsJobsRequest
var einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs = &cobra.Command{
	Use: "ListAllocateShipmentsJobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobs")
		return nil
	},
}

var einride_planning_v1_DeleteAllocateShipmentsJobRequest v1.DeleteAllocateShipmentsJobRequest
var einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob = &cobra.Command{
	Use: "DeleteAllocateShipmentsJob",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.DeleteAllocateShipmentsJob")
		return nil
	},
}

var einride_planning_v1_RunAllocateShipmentsJobRequest v1.RunAllocateShipmentsJobRequest
var einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob = &cobra.Command{
	Use: "RunAllocateShipmentsJob",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.RunAllocateShipmentsJob")
		return nil
	},
}

var einride_planning_v1_BatchRunAllocateShipmentsJobsRequest v1.BatchRunAllocateShipmentsJobsRequest
var einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs = &cobra.Command{
	Use: "BatchRunAllocateShipmentsJobs",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.BatchRunAllocateShipmentsJobs")
		return nil
	},
}

var einride_planning_v1_ListAllocateShipmentsJobExecutionsRequest v1.ListAllocateShipmentsJobExecutionsRequest
var einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions = &cobra.Command{
	Use: "ListAllocateShipmentsJobExecutions",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobExecutions")
		return nil
	},
}

func AddAllocateShipmentsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_AllocateShipmentsService)
}

func init() {
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs)
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions)
}
