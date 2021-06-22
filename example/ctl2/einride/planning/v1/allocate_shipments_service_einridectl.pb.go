package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planning.v1.AllocateShipmentsService.
var (
	einride_planning_v1_AllocateShipmentsServiceClient v1.AllocateShipmentsServiceClient
	einride_planning_v1_AllocateShipmentsService       = &cobra.Command{
		Use: "einride.planning.v1.AllocateShipmentsService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_AllocateShipmentsServiceClient = v1.NewAllocateShipmentsServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.CreateAllocateShipmentsJob.
var (
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request v1.CreateAllocateShipmentsJobRequest
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob         = &cobra.Command{
		Use: "CreateAllocateShipmentsJob",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.CreateAllocateShipmentsJob")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.GetAllocateShipmentsJob.
var (
	einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob_Request v1.GetAllocateShipmentsJobRequest
	einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob         = &cobra.Command{
		Use: "GetAllocateShipmentsJob",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.GetAllocateShipmentsJob")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.UpdateAllocateShipmentsJob.
var (
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request v1.UpdateAllocateShipmentsJobRequest
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob         = &cobra.Command{
		Use: "UpdateAllocateShipmentsJob",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.UpdateAllocateShipmentsJob")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobs.
var (
	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs_Request v1.ListAllocateShipmentsJobsRequest
	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs         = &cobra.Command{
		Use: "ListAllocateShipmentsJobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobs")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.DeleteAllocateShipmentsJob.
var (
	einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob_Request v1.DeleteAllocateShipmentsJobRequest
	einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob         = &cobra.Command{
		Use: "DeleteAllocateShipmentsJob",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.DeleteAllocateShipmentsJob")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.RunAllocateShipmentsJob.
var (
	einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob_Request v1.RunAllocateShipmentsJobRequest
	einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob         = &cobra.Command{
		Use: "RunAllocateShipmentsJob",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.RunAllocateShipmentsJob")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.BatchRunAllocateShipmentsJobs.
var (
	einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs_Request v1.BatchRunAllocateShipmentsJobsRequest
	einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs         = &cobra.Command{
		Use: "BatchRunAllocateShipmentsJobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.BatchRunAllocateShipmentsJobs")
			return nil
		},
	}
)

// einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobExecutions.
var (
	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions_Request v1.ListAllocateShipmentsJobExecutionsRequest
	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions         = &cobra.Command{
		Use: "ListAllocateShipmentsJobExecutions",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.AllocateShipmentsService.ListAllocateShipmentsJobExecutions")
			return nil
		},
	}
)

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
