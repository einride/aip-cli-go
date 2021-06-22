package planningv1

import (
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	log "log"
)

var einride_planning_v1_TaskService = &cobra.Command{
	Use: "einride.planning.v1.TaskService",
}

var einride_planning_v1_CreateTaskRequest v1.CreateTaskRequest
var einride_planning_v1_TaskService_CreateTask = &cobra.Command{
	Use: "CreateTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.CreateTask")
		return nil
	},
}

var einride_planning_v1_GetTaskRequest v1.GetTaskRequest
var einride_planning_v1_TaskService_GetTask = &cobra.Command{
	Use: "GetTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.GetTask")
		return nil
	},
}

var einride_planning_v1_BatchGetTasksRequest v1.BatchGetTasksRequest
var einride_planning_v1_TaskService_BatchGetTasks = &cobra.Command{
	Use: "BatchGetTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.BatchGetTasks")
		return nil
	},
}

var einride_planning_v1_SearchTasksRequest v1.SearchTasksRequest
var einride_planning_v1_TaskService_SearchTasks = &cobra.Command{
	Use: "SearchTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.SearchTasks")
		return nil
	},
}

var einride_planning_v1_UpdateTaskRequest v1.UpdateTaskRequest
var einride_planning_v1_TaskService_UpdateTask = &cobra.Command{
	Use: "UpdateTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.UpdateTask")
		return nil
	},
}

var einride_planning_v1_DeleteTaskRequest v1.DeleteTaskRequest
var einride_planning_v1_TaskService_DeleteTask = &cobra.Command{
	Use: "DeleteTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.DeleteTask")
		return nil
	},
}

var einride_planning_v1_UndeleteTaskRequest v1.UndeleteTaskRequest
var einride_planning_v1_TaskService_UndeleteTask = &cobra.Command{
	Use: "UndeleteTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.UndeleteTask")
		return nil
	},
}

var einride_planning_v1_AllocatePickupTaskRequest v1.AllocatePickupTaskRequest
var einride_planning_v1_TaskService_AllocatePickupTask = &cobra.Command{
	Use: "AllocatePickupTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.AllocatePickupTask")
		return nil
	},
}

var einride_planning_v1_DeallocatePickupTaskRequest v1.DeallocatePickupTaskRequest
var einride_planning_v1_TaskService_DeallocatePickupTask = &cobra.Command{
	Use: "DeallocatePickupTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.DeallocatePickupTask")
		return nil
	},
}

var einride_planning_v1_MutateTasksRequest v1.MutateTasksRequest
var einride_planning_v1_TaskService_MutateTasks = &cobra.Command{
	Use: "MutateTasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.MutateTasks")
		return nil
	},
}

var einride_planning_v1_StartTaskRequest v1.StartTaskRequest
var einride_planning_v1_TaskService_StartTask = &cobra.Command{
	Use: "StartTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.StartTask")
		return nil
	},
}

var einride_planning_v1_CompleteTaskRequest v1.CompleteTaskRequest
var einride_planning_v1_TaskService_CompleteTask = &cobra.Command{
	Use: "CompleteTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.CompleteTask")
		return nil
	},
}

var einride_planning_v1_ApproveDockTaskRequest v1.ApproveDockTaskRequest
var einride_planning_v1_TaskService_ApproveDockTask = &cobra.Command{
	Use: "ApproveDockTask",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("einride.planning.v1.TaskService.ApproveDockTask")
		return nil
	},
}

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
