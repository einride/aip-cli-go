package planningv1

import (
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Parent, "parent", "", "Resource name of the transport installation where this task will be created.")

	einride_planning_v1_TaskService_CreateTask_Request.Task = new(v1.Task)
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Name, "task.name", "", "The resource name.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Etag, "task.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedStartTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_CreateTask.Flags().Int64Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedStartTime.Seconds, "task.expectedStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_CreateTask.Flags().Int32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedStartTime.Nanos, "task.expectedStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedCompleteTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_CreateTask.Flags().Int64Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedCompleteTime.Seconds, "task.expectedCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_CreateTask.Flags().Int32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedCompleteTime.Nanos, "task.expectedCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_CreateTask.Flags().Float32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedStartRemainingEnergyJoules, "task.expectedStartRemainingEnergyJoules", 10, "The energy expected to be remaining when the task is started.")
	einride_planning_v1_TaskService_CreateTask.Flags().Float32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ExpectedCompleteRemainingEnergyJoules, "task.expectedCompleteRemainingEnergyJoules", 10, "The energy expected to be remaining when the task is completed.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Description, "task.description", "", "Free text description of the task.")
	// TODO: map: Annotationsstring message
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Carrier, "task.carrier", "", "Resource name of the carrier that the vehicle and\ndrivers executing the task belongs to.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Vehicle, "task.vehicle", "", "Resource name of the vehicle executing the task.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.Driver, "task.driver", "", "Resource name of the driver executing the entire task.\nIf this is set, both `start_driver` and `end_driver` are ignored.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.StartDriver, "task.startDriver", "", "Resource name of the driver executing the start of the task.\nIgnored if `driver` is also set.")
	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.Task.EndDriver, "task.endDriver", "", "Resource name of the driver executing the end of the task.\nIgnored if `driver` is also set.")
	// TODO: enum State
	einride_planning_v1_TaskService_CreateTask_Request.Task.ActualStartTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_CreateTask.Flags().Int64Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ActualStartTime.Seconds, "task.actualStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_CreateTask.Flags().Int32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ActualStartTime.Nanos, "task.actualStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_CreateTask_Request.Task.ActualCompleteTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_CreateTask.Flags().Int64Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ActualCompleteTime.Seconds, "task.actualCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_CreateTask.Flags().Int32Var(&einride_planning_v1_TaskService_CreateTask_Request.Task.ActualCompleteTime.Nanos, "task.actualCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: one-of: PickupShipment
	// TODO: one-of: DeliverShipment
	// TODO: one-of: Drive
	// TODO: one-of: Charge
	// TODO: one-of: Dock
	// TODO: one-of: Generic

	einride_planning_v1_TaskService_CreateTask.Flags().StringVar(&einride_planning_v1_TaskService_CreateTask_Request.TaskId, "taskId", "", "The ID to use for the task, which will become the final component of\nthe tasks resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\nlowercase letters, numbers and hyphens.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_GetTask)

	einride_planning_v1_TaskService_GetTask.Flags().StringVar(&einride_planning_v1_TaskService_GetTask_Request.Name, "name", "", "Resource name of the task to get.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_BatchGetTasks)

	einride_planning_v1_TaskService_BatchGetTasks.Flags().StringVar(&einride_planning_v1_TaskService_BatchGetTasks_Request.Parent, "parent", "", "Resource name of the parent transport installation shared by all\nvehicle plans requested.\nIf this is not a wildcard, the parent of all of the tasks specified in `names`\nmust match this field.")

	einride_planning_v1_TaskService_BatchGetTasks.Flags().StringSliceVar(&einride_planning_v1_TaskService_BatchGetTasks_Request.Names, "names", []string{}, "Resource names of the tasks requested.\nA maximum of 1000 tasks can be retrieved in one request.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_SearchTasks)

	einride_planning_v1_TaskService_SearchTasks.Flags().StringVar(&einride_planning_v1_TaskService_SearchTasks_Request.Parent, "parent", "", "Resource name of the transport installation, which\nowns the tasks to search for.\nTo search for tasks across transport installations,\nspecify a wildcard in the last segment:\n`transportInstallations/-`")

	einride_planning_v1_TaskService_SearchTasks.Flags().Int32Var(&einride_planning_v1_TaskService_SearchTasks_Request.PageSize, "pageSize", 10, "The maximum number of tasks to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 tasks will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planning_v1_TaskService_SearchTasks.Flags().StringVar(&einride_planning_v1_TaskService_SearchTasks_Request.PageToken, "pageToken", "", "A page token, received from a previous `SearchTasks` call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided to `SearchTasks` must match\nthe call that provided the page token.")

	einride_planning_v1_TaskService_SearchTasks.Flags().StringVar(&einride_planning_v1_TaskService_SearchTasks_Request.Filter, "filter", "", "Filter to search for.\n\nSupported filters:\nField                                          Operators     Note\n---\nstate                                          =\ncarrier                                        =\ndriver                                         =\nstart_driver                                   =\nend_driver                                     =\nvehicle                                        =\nexpected_start_time                            =, <, >\nexpected_complete_time                         =, <, >\nexpected_start_remaining_energy_joules         =, <, >\nexpected_complete_remaining_energy_joules      =, <, >\nannotations.{key}                              =           Search for tasks an exact annotation key value pair.\nannotations                                    :           Search for tasks with an annotation with the key (RHS)\n\nExamples:\nAll tasks in ACTIVE or COMPLETED state:\n`state = ACTIVE OR state = COMPLETED`\n\nAll tasks for a specific driver:\n`driver = \"carriers/1/drivers/1\" OR start_driver = \"carriers/1/drivers/1\" OR end_driver = \"carriers/1/drivers/1\"`")

	einride_planning_v1_TaskService_SearchTasks.Flags().StringVar(&einride_planning_v1_TaskService_SearchTasks_Request.OrderBy, "orderBy", "", "The fields to order the tasks by.\nIf not specified, order of tasks is undefined and may change.\n\nFormat:\n```\norder_by: field [{ asc | desc }] [, ...]\nfield: { expected_start_time | expected_complete_time }\n```\n\nExample:\n```\norder_by = \"expected_complete_time asc, expected_start_time asc\"\n```")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_UpdateTask)

	einride_planning_v1_TaskService_UpdateTask_Request.Task = new(v1.Task)
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Name, "task.name", "", "The resource name.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Etag, "task.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedStartTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_UpdateTask.Flags().Int64Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedStartTime.Seconds, "task.expectedStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Int32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedStartTime.Nanos, "task.expectedStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedCompleteTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_UpdateTask.Flags().Int64Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedCompleteTime.Seconds, "task.expectedCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Int32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedCompleteTime.Nanos, "task.expectedCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Float32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedStartRemainingEnergyJoules, "task.expectedStartRemainingEnergyJoules", 10, "The energy expected to be remaining when the task is started.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Float32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ExpectedCompleteRemainingEnergyJoules, "task.expectedCompleteRemainingEnergyJoules", 10, "The energy expected to be remaining when the task is completed.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Description, "task.description", "", "Free text description of the task.")
	// TODO: map: Annotationsstring message
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Carrier, "task.carrier", "", "Resource name of the carrier that the vehicle and\ndrivers executing the task belongs to.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Vehicle, "task.vehicle", "", "Resource name of the vehicle executing the task.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.Driver, "task.driver", "", "Resource name of the driver executing the entire task.\nIf this is set, both `start_driver` and `end_driver` are ignored.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.StartDriver, "task.startDriver", "", "Resource name of the driver executing the start of the task.\nIgnored if `driver` is also set.")
	einride_planning_v1_TaskService_UpdateTask.Flags().StringVar(&einride_planning_v1_TaskService_UpdateTask_Request.Task.EndDriver, "task.endDriver", "", "Resource name of the driver executing the end of the task.\nIgnored if `driver` is also set.")
	// TODO: enum State
	einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualStartTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_UpdateTask.Flags().Int64Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualStartTime.Seconds, "task.actualStartTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Int32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualStartTime.Nanos, "task.actualStartTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualCompleteTime = new(timestamppb.Timestamp)
	einride_planning_v1_TaskService_UpdateTask.Flags().Int64Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualCompleteTime.Seconds, "task.actualCompleteTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_TaskService_UpdateTask.Flags().Int32Var(&einride_planning_v1_TaskService_UpdateTask_Request.Task.ActualCompleteTime.Nanos, "task.actualCompleteTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: one-of: PickupShipment
	// TODO: one-of: DeliverShipment
	// TODO: one-of: Drive
	// TODO: one-of: Charge
	// TODO: one-of: Dock
	// TODO: one-of: Generic

	einride_planning_v1_TaskService_UpdateTask_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planning_v1_TaskService_UpdateTask.Flags().StringSliceVar(&einride_planning_v1_TaskService_UpdateTask_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_DeleteTask)

	einride_planning_v1_TaskService_DeleteTask.Flags().StringVar(&einride_planning_v1_TaskService_DeleteTask_Request.Name, "name", "", "Resource name of the task to delete.")

	einride_planning_v1_TaskService_DeleteTask.Flags().StringVar(&einride_planning_v1_TaskService_DeleteTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_UndeleteTask)

	einride_planning_v1_TaskService_UndeleteTask.Flags().StringVar(&einride_planning_v1_TaskService_UndeleteTask_Request.Name, "name", "", "Resource name of the task to delete.")

	einride_planning_v1_TaskService_UndeleteTask.Flags().StringVar(&einride_planning_v1_TaskService_UndeleteTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_AllocatePickupTask)

	einride_planning_v1_TaskService_AllocatePickupTask.Flags().StringVar(&einride_planning_v1_TaskService_AllocatePickupTask_Request.Name, "name", "", "Resource name of the task to allocate.")

	einride_planning_v1_TaskService_AllocatePickupTask.Flags().StringVar(&einride_planning_v1_TaskService_AllocatePickupTask_Request.Shipment, "shipment", "", "Resource name of the shipment to assign to the pickup task.")

	// TODO: enum ShipmentAssignmentSource

	einride_planning_v1_TaskService_AllocatePickupTask.Flags().StringVar(&einride_planning_v1_TaskService_AllocatePickupTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_DeallocatePickupTask)

	einride_planning_v1_TaskService_DeallocatePickupTask.Flags().StringVar(&einride_planning_v1_TaskService_DeallocatePickupTask_Request.Name, "name", "", "Resource name of the task to deallocate.")

	einride_planning_v1_TaskService_DeallocatePickupTask.Flags().StringVar(&einride_planning_v1_TaskService_DeallocatePickupTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_MutateTasks)

	einride_planning_v1_TaskService_MutateTasks.Flags().StringVar(&einride_planning_v1_TaskService_MutateTasks_Request.Parent, "parent", "", "Resource name of the transport installation all tasks to mutate\nbelongs to.")

	// TODO: list: Operations message
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_StartTask)

	einride_planning_v1_TaskService_StartTask.Flags().StringVar(&einride_planning_v1_TaskService_StartTask_Request.Name, "name", "", "Resource name of the task to start.")

	einride_planning_v1_TaskService_StartTask.Flags().StringVar(&einride_planning_v1_TaskService_StartTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_CompleteTask)

	einride_planning_v1_TaskService_CompleteTask.Flags().StringVar(&einride_planning_v1_TaskService_CompleteTask_Request.Name, "name", "", "Resource name of the vehicle plan to complete.")

	einride_planning_v1_TaskService_CompleteTask.Flags().StringVar(&einride_planning_v1_TaskService_CompleteTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
	einride_planning_v1_TaskService.AddCommand(einride_planning_v1_TaskService_ApproveDockTask)

	einride_planning_v1_TaskService_ApproveDockTask.Flags().StringVar(&einride_planning_v1_TaskService_ApproveDockTask_Request.Name, "name", "", "Resource name of the task to approve docking for.")

	einride_planning_v1_TaskService_ApproveDockTask.Flags().StringVar(&einride_planning_v1_TaskService_ApproveDockTask_Request.Etag, "etag", "", "The current etag of the resource.\nIf an etag is provided and does not match the current etag of the resource,\nthe request will aborted and an ABORTED error will be returned.")
}
