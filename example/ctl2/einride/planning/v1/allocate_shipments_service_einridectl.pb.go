package planningv1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1 "github.com/einride/proto/gen/go/einride/planning/v1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.planning.v1.AllocateShipmentsService.
var (
	einride_planning_v1_AllocateShipmentsServiceClient v1.AllocateShipmentsServiceClient
	einride_planning_v1_AllocateShipmentsService       = &cobra.Command{
		Use:   "einride.planning.v1.AllocateShipmentsService",
		Short: "A service for allocating shipments to pickup tasks.",
		Long:  "A service for allocating shipments to pickup tasks.",
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.CreateAllocateShipmentsJob(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.GetAllocateShipmentsJob(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.UpdateAllocateShipmentsJob(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.ListAllocateShipmentsJobs(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.DeleteAllocateShipmentsJob(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.RunAllocateShipmentsJob(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.BatchRunAllocateShipmentsJobs(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
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
			response, err := einride_planning_v1_AllocateShipmentsServiceClient.ListAllocateShipmentsJobExecutions(cmd.Context(), &einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddAllocateShipmentsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_planning_v1_AllocateShipmentsService)
}

func init() {
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob)

	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.Parent, "parent", "", "Resource name of the transport installation where the job should be created.")

	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob = new(v1.AllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Name, "allocateShipmentsJob.name", "", "The resource name of the job.\nFor example: `\"transportInstallations/transportInstallation/allocateShipmentsJobs/allocateShipmentsJob\"`.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Etag, "allocateShipmentsJob.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().BoolVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Automatic, "allocateShipmentsJob.automatic", false, "Flag indicating whether the job is `Run` automatically.\nIf set, every hour (not configurable) `Run` will be called on the job.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Shipper, "allocateShipmentsJob.shipper", "", "Resource name of the shipper which shipments should be allocated for.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringSliceVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Lanes, "allocateShipmentsJob.lanes", []string{}, "Resource names of lanes which shipments should be allocated on.\nIf this list is empty, all the shippers shipments will be considered.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.CarrierReferenceId, "allocateShipmentsJob.carrierReferenceId", "", "The carrier reference ID which shipments should have to be allocated.\nIf not set, no filtering on carrier reference ID is done.")
	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJob.Carrier, "allocateShipmentsJob.carrier", "", "Resource name of the carrier who owns the tasks that shipments should be\nallocated to.")

	einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_CreateAllocateShipmentsJob_Request.AllocateShipmentsJobId, "allocateShipmentsJobId", "", "The ID to use for the resource, which will become the final component of\nthe resource name.\n\nIf an ID is not provided, an ID will be automatically generated.\n\nThis value should be 4-63 characters, and valid characters are\nlowercase letters, numbers and hyphens.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob)

	einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_GetAllocateShipmentsJob_Request.Name, "name", "", "Resource name of the job to get.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob)

	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob = new(v1.AllocateShipmentsJob)
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Name, "allocateShipmentsJob.name", "", "The resource name of the job.\nFor example: `\"transportInstallations/transportInstallation/allocateShipmentsJobs/allocateShipmentsJob\"`.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Etag, "allocateShipmentsJob.etag", "", "This checksum is computed by the server based on the value of other\nfields, and may be sent on update and delete requests to ensure the client\nhas an up-to-date value before proceeding.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().BoolVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Automatic, "allocateShipmentsJob.automatic", false, "Flag indicating whether the job is `Run` automatically.\nIf set, every hour (not configurable) `Run` will be called on the job.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Shipper, "allocateShipmentsJob.shipper", "", "Resource name of the shipper which shipments should be allocated for.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringSliceVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Lanes, "allocateShipmentsJob.lanes", []string{}, "Resource names of lanes which shipments should be allocated on.\nIf this list is empty, all the shippers shipments will be considered.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.CarrierReferenceId, "allocateShipmentsJob.carrierReferenceId", "", "The carrier reference ID which shipments should have to be allocated.\nIf not set, no filtering on carrier reference ID is done.")
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.AllocateShipmentsJob.Carrier, "allocateShipmentsJob.carrier", "", "Resource name of the carrier who owns the tasks that shipments should be\nallocated to.")

	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob.Flags().StringSliceVar(&einride_planning_v1_AllocateShipmentsService_UpdateAllocateShipmentsJob_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs)

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs_Request.Parent, "parent", "", "Resource name of the transport installation, which owns the allocate shipments jobs to list.\nWildcard parent is not supported.")

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs.Flags().Int32Var(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs_Request.PageSize, "pageSize", 10, "The maximum number of instances to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 instances will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobs_Request.PageToken, "pageToken", "", "A page token, received from a previous `ListAllocateShipmentsJobs` call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided to `ListAllocateShipmentsJobs` must match\nthe call that provided the page token.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob)

	einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_DeleteAllocateShipmentsJob_Request.Name, "name", "", "Resource name of the job to delete.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob)

	einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_RunAllocateShipmentsJob_Request.Name, "name", "", "Resource name of the job to run.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs)

	einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_BatchRunAllocateShipmentsJobs_Request.Parent, "parent", "", "Resource name of the transport installation to run all jobs for.\nTo run all jobs, regardless of transport installation, specify a wildcard as the\nlast segment.")
	einride_planning_v1_AllocateShipmentsService.AddCommand(einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions)

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions_Request.Parent, "parent", "", "Resource name of the allocate shipments jobs, which owns the executions to list.\nTo list executions across jobs, specify a wildcard as the last segment.")

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions.Flags().Int32Var(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions_Request.PageSize, "pageSize", 10, "The maximum number of executions to return. The service may return fewer than\nthis value.\nIf unspecified, at most 50 executions will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions.Flags().StringVar(&einride_planning_v1_AllocateShipmentsService_ListAllocateShipmentsJobExecutions_Request.PageToken, "pageToken", "", "A page token, received from a previous `ListAllocateShipmentsJobExecutions` call.\nProvide this to retrieve the subsequent page.\n\nWhen paginating, all other parameters provided to `ListAllocateShipmentsJobExecutions` must match\nthe call that provided the page token.")
}
