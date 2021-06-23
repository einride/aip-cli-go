package planningv1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// einride.planning.v1.ScheduleBookingService.
var (
	einride_planning_v1_ScheduleBookingServiceClient v1.ScheduleBookingServiceClient
	einride_planning_v1_ScheduleBookingService       = &cobra.Command{
		Use:   "einride.planning.v1.ScheduleBookingService",
		Short: "ScheduleBooking service.",
		Long:  "ScheduleBooking service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planning_v1_ScheduleBookingServiceClient = v1.NewScheduleBookingServiceClient(conn)
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.CreateScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request v1.CreateScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking         = &cobra.Command{
		Use: "CreateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.CreateScheduleBooking(cmd.Context(), &einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.GetScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_GetScheduleBooking_Request v1.GetScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_GetScheduleBooking         = &cobra.Command{
		Use: "GetScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.GetScheduleBooking(cmd.Context(), &einride_planning_v1_ScheduleBookingService_GetScheduleBooking_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookings.
var (
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings_Request v1.BatchGetScheduleBookingsRequest
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings         = &cobra.Command{
		Use: "BatchGetScheduleBookings",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.BatchGetScheduleBookings(cmd.Context(), &einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UpdateScheduleBooking.
var (
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request v1.UpdateScheduleBookingRequest
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking         = &cobra.Command{
		Use: "UpdateScheduleBooking",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.UpdateScheduleBooking(cmd.Context(), &einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.GetScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance_Request v1.GetScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance         = &cobra.Command{
		Use: "GetScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.GetScheduleBookingInstance(cmd.Context(), &einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.BatchGetScheduleBookingInstances.
var (
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances_Request v1.BatchGetScheduleBookingInstancesRequest
	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances         = &cobra.Command{
		Use: "BatchGetScheduleBookingInstances",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.BatchGetScheduleBookingInstances(cmd.Context(), &einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.ListScheduleBookingInstances.
var (
	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request v1.ListScheduleBookingInstancesRequest
	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances         = &cobra.Command{
		Use: "ListScheduleBookingInstances",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.ListScheduleBookingInstances(cmd.Context(), &einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UpdateScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request v1.UpdateScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance         = &cobra.Command{
		Use: "UpdateScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.UpdateScheduleBookingInstance(cmd.Context(), &einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.ReconcileScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance_Request v1.ReconcileScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance         = &cobra.Command{
		Use: "ReconcileScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.ReconcileScheduleBookingInstance(cmd.Context(), &einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.DeleteScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance_Request v1.DeleteScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance         = &cobra.Command{
		Use: "DeleteScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.DeleteScheduleBookingInstance(cmd.Context(), &einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planning.v1.ScheduleBookingService.UndeleteScheduleBookingInstance.
var (
	einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance_Request v1.UndeleteScheduleBookingInstanceRequest
	einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance         = &cobra.Command{
		Use: "UndeleteScheduleBookingInstance",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planning_v1_ScheduleBookingServiceClient.UndeleteScheduleBookingInstance(cmd.Context(), &einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddScheduleBookingServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_ScheduleBookingService)
}

func init() {
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_CreateScheduleBooking)

	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.Parent, "parent", "", "Resource name of the transport installation where this schedule will be created.")

	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking = new(v1.ScheduleBooking)
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Name, "scheduleBooking.name", "", "The resource name.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Etag, "scheduleBooking.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Spec = new(v1.ScheduleBookingSpec)
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.Spec.Schedule, "scheduleBooking.spec.schedule", "", "Resource name of the schedule to use..\nMust point to a specific revision of a schedule. Trying to use\na schedule without providing a revision will fail.")
	// TODO: list: VehicleSelections message
	// TODO: list: DriverSelections message
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule = new(v1.RecurrenceRule)
	// TODO: enum Frequency
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.Interval, "scheduleBooking.recurrenceRule.interval", 10, "The interval between each frequency iteration.\nIf not set, this value will default to '1'.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime = new(timestamppb.Timestamp)
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int64Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime.Seconds, "scheduleBooking.recurrenceRule.startTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime.Nanos, "scheduleBooking.recurrenceRule.startTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.Count, "scheduleBooking.recurrenceRule.count", 10, "The maximum number of instances to create from the rule.\nIf not set, there is no limit on number of instances.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime = new(timestamppb.Timestamp)
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int64Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime.Seconds, "scheduleBooking.recurrenceRule.endTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime.Nanos, "scheduleBooking.recurrenceRule.endTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: enum WeekStart
	// TODO: list: Weekdays enum

	einride_planning_v1_ScheduleBookingService_CreateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_CreateScheduleBooking_Request.ScheduleBookingId, "scheduleBookingId", "", "The ID to use for the schedule booking, which will become the final component of\nthe schedule booking's resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\nlowercase letters, numbers and hyphens.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_GetScheduleBooking)

	einride_planning_v1_ScheduleBookingService_GetScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_GetScheduleBooking_Request.Name, "name", "", "Resource name of the schedule booking to get.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings)

	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings_Request.Parent, "parent", "", "Resource name of the transport installation shared by all schedule bookings.\nTo get schedules across parents, specify a wildcard as the last segment.\nIf this is not a wildcard, the parent of all of the schedule bookings specified in `names`\nmust match this field.")

	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings.Flags().StringSliceVar(&einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookings_Request.Names, "names", []string{}, "Resource names of the schedule bookings to retrieve.\nA maximum of 1000 schedule bookings can be retrieved in a batch.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking)

	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking = new(v1.ScheduleBooking)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Name, "scheduleBooking.name", "", "The resource name.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Etag, "scheduleBooking.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Spec = new(v1.ScheduleBookingSpec)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.Spec.Schedule, "scheduleBooking.spec.schedule", "", "Resource name of the schedule to use..\nMust point to a specific revision of a schedule. Trying to use\na schedule without providing a revision will fail.")
	// TODO: list: VehicleSelections message
	// TODO: list: DriverSelections message
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule = new(v1.RecurrenceRule)
	// TODO: enum Frequency
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.Interval, "scheduleBooking.recurrenceRule.interval", 10, "The interval between each frequency iteration.\nIf not set, this value will default to '1'.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime = new(timestamppb.Timestamp)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int64Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime.Seconds, "scheduleBooking.recurrenceRule.startTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.StartTime.Nanos, "scheduleBooking.recurrenceRule.startTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.Count, "scheduleBooking.recurrenceRule.count", 10, "The maximum number of instances to create from the rule.\nIf not set, there is no limit on number of instances.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime = new(timestamppb.Timestamp)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int64Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime.Seconds, "scheduleBooking.recurrenceRule.endTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.ScheduleBooking.RecurrenceRule.EndTime.Nanos, "scheduleBooking.recurrenceRule.endTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")
	// TODO: enum WeekStart
	// TODO: list: Weekdays enum

	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking.Flags().StringSliceVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBooking_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance)

	einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_GetScheduleBookingInstance_Request.Name, "name", "", "Resource name of the schedule booking instance to get.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances)

	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances_Request.Parent, "parent", "", "Resource name of the schedule booking shared by all instances.\nTo get schedules across parents, specify a wildcard as the last segment.\nIf this is not a wildcard, the parent of all of the schedules specified in `names`\nmust match this field.")

	einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances.Flags().StringSliceVar(&einride_planning_v1_ScheduleBookingService_BatchGetScheduleBookingInstances_Request.Names, "names", []string{}, "Resource names of the schedule instances to retrieve.\nA maximum of 1000 schedule instances can be retrieved in a batch.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances)

	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request.Parent, "parent", "", "Resource name of the schedule booking, which owns the schedule instances to list.\nTo list schedule booking instances across schedule bookings, specify\na wildcard as the last segment.")

	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request.Filter, "filter", "", "The filter to applied to schedules booking instances listed.\n\nSupported filters:\nField              Operators   Note\n---\nstart_time         =,<,>\ntasks_start_time   =,<,>\ntasks_end_time     =,<,>\nstate              =\nspec.schedule      =           The RHS can be a schedule resource name with or without a revision.\n                               If no revision is specified, any schedule booking instance referring to the\n                               schedule is returned.\nis_exception       =")

	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request.PageSize, "pageSize", 10, "The maximum number of instances to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 instances will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_ListScheduleBookingInstances_Request.PageToken, "pageToken", "", "A page token, received from a previous `ListScheduleBookingInstances` call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided to `ListScheduleBookingInstances` must match\nthe call that provided the page token.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance)

	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance = new(v1.ScheduleBookingInstance)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.Name, "scheduleBookingInstance.name", "", "The resource name.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.Etag, "scheduleBookingInstance.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.Spec = new(v1.ScheduleBookingSpec)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.Spec.Schedule, "scheduleBookingInstance.spec.schedule", "", "Resource name of the schedule to use..\nMust point to a specific revision of a schedule. Trying to use\na schedule without providing a revision will fail.")
	// TODO: list: VehicleSelections message
	// TODO: list: DriverSelections message
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.StartTime = new(timestamppb.Timestamp)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().Int64Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.StartTime.Seconds, "scheduleBookingInstance.startTime.seconds", 10, "Represents seconds of UTC time since Unix epoch\n1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to\n9999-12-31T23:59:59Z inclusive.")
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().Int32Var(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.ScheduleBookingInstance.StartTime.Nanos, "scheduleBookingInstance.startTime.nanos", 10, "Non-negative fractions of a second at nanosecond resolution. Negative\nsecond values with fractions must still have non-negative nanos values\nthat count forward in time. Must be from 0 to 999,999,999\ninclusive.")

	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance.Flags().StringSliceVar(&einride_planning_v1_ScheduleBookingService_UpdateScheduleBookingInstance_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance)

	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance_Request.Name, "name", "", "Resource name of the schedule booking instance to reconcile.")

	einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance.Flags().BoolVar(&einride_planning_v1_ScheduleBookingService_ReconcileScheduleBookingInstance_Request.Force, "force", false, "If set to true, reconciliation will be attempted even\nif some safety checks reports false.\nFor example, reconciliation will be attempted even if\nsome of the tasks it would remove are `STARTED` or `COMPLETED`.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance)

	einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_DeleteScheduleBookingInstance_Request.Name, "name", "", "Resource name of the schedule to delete.")
	einride_planning_v1_ScheduleBookingService.AddCommand(einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance)

	einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance.Flags().StringVar(&einride_planning_v1_ScheduleBookingService_UndeleteScheduleBookingInstance_Request.Name, "name", "", "Resource name of the schedule to delete.")
}
