package planningv1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.planning.v1.ScheduleService.
var (
	einride_planning_v1_ScheduleServiceClient v1.ScheduleServiceClient
	einride_planning_v1_ScheduleService       = &cobra.Command{
		Use:   "einride.planning.v1.ScheduleService",
		Short: "Schedule service.",
		Long:  "Schedule service.",
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
			response, err := einride_planning_v1_ScheduleServiceClient.CreateSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_CreateSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.GetSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_GetSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.BatchGetSchedules(cmd.Context(), &einride_planning_v1_ScheduleService_BatchGetSchedules_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.ListSchedules(cmd.Context(), &einride_planning_v1_ScheduleService_ListSchedules_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.UpdateSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_UpdateSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.PublishSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_PublishSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.DeleteSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_DeleteSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_ScheduleServiceClient.UndeleteSchedule(cmd.Context(), &einride_planning_v1_ScheduleService_UndeleteSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_ScheduleService)
}

func init() {
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_CreateSchedule)

	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Parent, "parent", "", "Resource name of the transport installation where this schedule will be created.")

	einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule = new(v1.Schedule)
	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule.Name, "schedule.name", "", "The resource name.")
	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule.RevisionId, "schedule.revisionId", "", "The revision ID of the schedule.\nA new revision is committed by the service whenever the schedule is\nchanged in any way.")
	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule.Etag, "schedule.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule.DisplayName, "schedule.displayName", "", "Display name of the schedule.")
	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.Schedule.Description, "schedule.description", "", "Free text description of the schedule.")
	// TODO: list: Vehicles message
	// TODO: list: Drivers message
	// TODO: list: Tasks message

	einride_planning_v1_ScheduleService_CreateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_CreateSchedule_Request.ScheduleId, "scheduleId", "", "The ID to use for the schedule, which will become the final component of\nthe schedules resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\nlowercase letters, numbers and hyphens.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_GetSchedule)

	einride_planning_v1_ScheduleService_GetSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_GetSchedule_Request.Name, "name", "", "Resource name of the schedule to get.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_BatchGetSchedules)

	einride_planning_v1_ScheduleService_BatchGetSchedules.Flags().StringVar(&einride_planning_v1_ScheduleService_BatchGetSchedules_Request.Parent, "parent", "", "Resource name of the transport installation shared by all schedules.\nTo get schedules across parents, specify a wildcard as the last segment.\nIf this is not a wildcard, the parent of all of the schedules specified in `names`\nmust match this field.")

	einride_planning_v1_ScheduleService_BatchGetSchedules.Flags().StringSliceVar(&einride_planning_v1_ScheduleService_BatchGetSchedules_Request.Names, "names", []string{}, "Resource names of the schedules to retrieve.\nA maximum of 1000 schedules can be retrieved in a batch.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_ListSchedules)

	einride_planning_v1_ScheduleService_ListSchedules.Flags().StringVar(&einride_planning_v1_ScheduleService_ListSchedules_Request.Parent, "parent", "", "Resource name of the transport installation, which\nowns the schedules to search for.")

	einride_planning_v1_ScheduleService_ListSchedules.Flags().Int32Var(&einride_planning_v1_ScheduleService_ListSchedules_Request.PageSize, "pageSize", 10, "The maximum number of schedules to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 schedules will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planning_v1_ScheduleService_ListSchedules.Flags().StringVar(&einride_planning_v1_ScheduleService_ListSchedules_Request.PageToken, "pageToken", "", "A page token, received from a previous `ListSchedules` call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided to `ListSchedules` must match\nthe call that provided the page token.")

	// TODO: enum View
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_UpdateSchedule)

	einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule = new(v1.Schedule)
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule.Name, "schedule.name", "", "The resource name.")
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule.RevisionId, "schedule.revisionId", "", "The revision ID of the schedule.\nA new revision is committed by the service whenever the schedule is\nchanged in any way.")
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule.Etag, "schedule.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule.DisplayName, "schedule.displayName", "", "Display name of the schedule.")
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.Schedule.Description, "schedule.description", "", "Free text description of the schedule.")
	// TODO: list: Vehicles message
	// TODO: list: Drivers message
	// TODO: list: Tasks message

	einride_planning_v1_ScheduleService_UpdateSchedule_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planning_v1_ScheduleService_UpdateSchedule.Flags().StringSliceVar(&einride_planning_v1_ScheduleService_UpdateSchedule_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_PublishSchedule)

	einride_planning_v1_ScheduleService_PublishSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_PublishSchedule_Request.Name, "name", "", "Resource name of the schedule to publish.\nCan only refer to the latest revision of the schedule.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_DeleteSchedule)

	einride_planning_v1_ScheduleService_DeleteSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_DeleteSchedule_Request.Name, "name", "", "Resource name of the schedule to delete.")
	einride_planning_v1_ScheduleService.AddCommand(einride_planning_v1_ScheduleService_UndeleteSchedule)

	einride_planning_v1_ScheduleService_UndeleteSchedule.Flags().StringVar(&einride_planning_v1_ScheduleService_UndeleteSchedule_Request.Name, "name", "", "Resource name of the schedule to delete.")
}
