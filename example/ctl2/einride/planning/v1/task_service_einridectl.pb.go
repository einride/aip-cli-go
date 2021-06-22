package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

// einride.planning.v1.TaskService.
var (
	einride_planning_v1_TaskServiceClient v1.TaskServiceClient
	einride_planning_v1_TaskService       = &cobra.Command{
		Use: "einride.planning.v1.TaskService",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_TaskServiceClient = v1.NewTaskServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.TaskService.CreateTask.
var (
	einride_planning_v1_TaskService_CreateTask_Request v1.CreateTaskRequest
	einride_planning_v1_TaskService_CreateTask         = &cobra.Command{
		Use: "CreateTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.CreateTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.GetTask.
var (
	einride_planning_v1_TaskService_GetTask_Request v1.GetTaskRequest
	einride_planning_v1_TaskService_GetTask         = &cobra.Command{
		Use: "GetTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.GetTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.BatchGetTasks.
var (
	einride_planning_v1_TaskService_BatchGetTasks_Request v1.BatchGetTasksRequest
	einride_planning_v1_TaskService_BatchGetTasks         = &cobra.Command{
		Use: "BatchGetTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.BatchGetTasks")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.SearchTasks.
var (
	einride_planning_v1_TaskService_SearchTasks_Request v1.SearchTasksRequest
	einride_planning_v1_TaskService_SearchTasks         = &cobra.Command{
		Use: "SearchTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.SearchTasks")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.UpdateTask.
var (
	einride_planning_v1_TaskService_UpdateTask_Request v1.UpdateTaskRequest
	einride_planning_v1_TaskService_UpdateTask         = &cobra.Command{
		Use: "UpdateTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.UpdateTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.DeleteTask.
var (
	einride_planning_v1_TaskService_DeleteTask_Request v1.DeleteTaskRequest
	einride_planning_v1_TaskService_DeleteTask         = &cobra.Command{
		Use: "DeleteTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.DeleteTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.UndeleteTask.
var (
	einride_planning_v1_TaskService_UndeleteTask_Request v1.UndeleteTaskRequest
	einride_planning_v1_TaskService_UndeleteTask         = &cobra.Command{
		Use: "UndeleteTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.UndeleteTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.AllocatePickupTask.
var (
	einride_planning_v1_TaskService_AllocatePickupTask_Request v1.AllocatePickupTaskRequest
	einride_planning_v1_TaskService_AllocatePickupTask         = &cobra.Command{
		Use: "AllocatePickupTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.AllocatePickupTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.DeallocatePickupTask.
var (
	einride_planning_v1_TaskService_DeallocatePickupTask_Request v1.DeallocatePickupTaskRequest
	einride_planning_v1_TaskService_DeallocatePickupTask         = &cobra.Command{
		Use: "DeallocatePickupTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.DeallocatePickupTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.MutateTasks.
var (
	einride_planning_v1_TaskService_MutateTasks_Request v1.MutateTasksRequest
	einride_planning_v1_TaskService_MutateTasks         = &cobra.Command{
		Use: "MutateTasks",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.MutateTasks")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.StartTask.
var (
	einride_planning_v1_TaskService_StartTask_Request v1.StartTaskRequest
	einride_planning_v1_TaskService_StartTask         = &cobra.Command{
		Use: "StartTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.StartTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.CompleteTask.
var (
	einride_planning_v1_TaskService_CompleteTask_Request v1.CompleteTaskRequest
	einride_planning_v1_TaskService_CompleteTask         = &cobra.Command{
		Use: "CompleteTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.CompleteTask")
			return nil
		},
	}
)

// einride.planning.v1.TaskService.ApproveDockTask.
var (
	einride_planning_v1_TaskService_ApproveDockTask_Request v1.ApproveDockTaskRequest
	einride_planning_v1_TaskService_ApproveDockTask         = &cobra.Command{
		Use: "ApproveDockTask",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Println("einride.planning.v1.TaskService.ApproveDockTask")
			return nil
		},
	}
)

func AddTaskServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_TaskService)
}

func init() {
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_CreateTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_GetTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_BatchGetTasks)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_SearchTasks)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_UpdateTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_DeleteTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_UndeleteTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_AllocatePickupTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_DeallocatePickupTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_MutateTasks)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_StartTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_CompleteTask)
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_ApproveDockTask)
}
