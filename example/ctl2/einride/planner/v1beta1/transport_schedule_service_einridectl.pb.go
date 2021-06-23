package plannerv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/planner/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.planner.v1beta1.TransportScheduleService.
var (
	einride_planner_v1beta1_TransportScheduleServiceClient v1beta1.TransportScheduleServiceClient
	einride_planner_v1beta1_TransportScheduleService       = &cobra.Command{
		Use:   "einride.planner.v1beta1.TransportScheduleService",
		Short: "Transport schedule service.",
		Long:  "Transport schedule service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_planner_v1beta1_TransportScheduleServiceClient = v1beta1.NewTransportScheduleServiceClient(conn)
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.CreateTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request v1beta1.CreateTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule         = &cobra.Command{
		Use: "CreateTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.CreateTransportSchedule(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.GetTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule_Request v1beta1.GetTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule         = &cobra.Command{
		Use: "GetTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.GetTransportSchedule(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.UpdateTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request v1beta1.UpdateTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule         = &cobra.Command{
		Use: "UpdateTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.UpdateTransportSchedule(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.ListTransportSchedules.
var (
	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request v1beta1.ListTransportSchedulesRequest
	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules         = &cobra.Command{
		Use: "ListTransportSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.ListTransportSchedules(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.BatchGetTransportSchedules.
var (
	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules_Request v1beta1.BatchGetTransportSchedulesRequest
	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules         = &cobra.Command{
		Use: "BatchGetTransportSchedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.BatchGetTransportSchedules(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.DeleteTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule_Request v1beta1.DeleteTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule         = &cobra.Command{
		Use: "DeleteTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.DeleteTransportSchedule(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.planner.v1beta1.TransportScheduleService.UndeleteTransportSchedule.
var (
	einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule_Request v1beta1.UndeleteTransportScheduleRequest
	einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule         = &cobra.Command{
		Use: "UndeleteTransportSchedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_planner_v1beta1_TransportScheduleServiceClient.UndeleteTransportSchedule(cmd.Context(), &einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddTransportScheduleServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planner_v1beta1_TransportScheduleService)
}

func init() {
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule)

	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.Parent, "parent", "", "Resource name of the parent transport installation where this transport schedule will be created.\nPattern: `transportInstallations/{transport_installation}`.")

	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule = new(v1beta1.TransportSchedule)
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule.Name, "transportSchedule.name", "", "The resource name of the transport schedule.")
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule.Etag, "transportSchedule.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule.DisplayName, "transportSchedule.displayName", "", "Display name of the schedule.")
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule.Description, "transportSchedule.description", "", "A description of the schedule.\n\"E.g. 10 truckloads to Trensums, 4 truckloads to Seafrigo\".")
	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportSchedule.TimeZone, "transportSchedule.timeZone", "", "Time zone of the schedule.")
	// TODO: list: VehicleSchedules message
	// TODO: list: DriverSchedules message

	einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_CreateTransportSchedule_Request.TransportScheduleId, "transportScheduleId", "", "The ID to use for the transport schedule.\nWill become the final component of the transport schedule's resource name.\nIf an ID is not provided, a unique ID will be selected by the service.\nThe ID should be 3-63 characters and valid characters are /[a-zA-Z0-9]/.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule)

	einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_GetTransportSchedule_Request.Name, "name", "", "Resource name of the transport installation to retrieve.\nPattern: `transportInstallations/{transport_installation}/schedules/{schedule}`.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule)

	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule = new(v1beta1.TransportSchedule)
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule.Name, "transportSchedule.name", "", "The resource name of the transport schedule.")
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule.Etag, "transportSchedule.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule.DisplayName, "transportSchedule.displayName", "", "Display name of the schedule.")
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule.Description, "transportSchedule.description", "", "A description of the schedule.\n\"E.g. 10 truckloads to Trensums, 4 truckloads to Seafrigo\".")
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.TransportSchedule.TimeZone, "transportSchedule.timeZone", "", "Time zone of the schedule.")
	// TODO: list: VehicleSchedules message
	// TODO: list: DriverSchedules message

	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule.Flags().StringSliceVar(&einride_planner_v1beta1_TransportScheduleService_UpdateTransportSchedule_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules)

	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request.Parent, "parent", "", "Resource name of the parent transport installation.")

	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules.Flags().Int32Var(&einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_ListTransportSchedules_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules)

	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules_Request.Parent, "parent", "", "Resource name of the parent transport installation shared by all transport schedules being retrieved.\nIf set, the parent of all transport schedules specified in `names` must match this field.\nPattern: `transportInstallations/{transport_installation}`.")

	einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules.Flags().StringSliceVar(&einride_planner_v1beta1_TransportScheduleService_BatchGetTransportSchedules_Request.Names, "names", []string{}, "Resource names of the transport schedules to retrieve.\nA maximum of 1000 transport schedules can be retrieved in a batch.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule)

	einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_DeleteTransportSchedule_Request.Name, "name", "", "Resource name of the transport schedule to delete.")
	einride_planner_v1beta1_TransportScheduleService.AddCommand(einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule)

	einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule.Flags().StringVar(&einride_planner_v1beta1_TransportScheduleService_UndeleteTransportSchedule_Request.Name, "name", "", "Resource name of the transport schedule to undelete.")
}
