package shipperv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

// einride.shipper.v1beta1.LaneService.
var (
	einride_shipper_v1beta1_LaneServiceClient v1beta1.LaneServiceClient
	einride_shipper_v1beta1_LaneService       = &cobra.Command{
		Use:   "shipper.v1beta1.LaneService",
		Short: "Lane resource service.",
		Long:  "Lane resource service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_v1beta1_LaneServiceClient = v1beta1.NewLaneServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.CreateLane.
var (
	einride_shipper_v1beta1_LaneService_CreateLane_Request v1beta1.CreateLaneRequest
	einride_shipper_v1beta1_LaneService_CreateLane         = &cobra.Command{
		Use: "CreateLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.CreateLane(cmd.Context(), &einride_shipper_v1beta1_LaneService_CreateLane_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.GetLane.
var (
	einride_shipper_v1beta1_LaneService_GetLane_Request v1beta1.GetLaneRequest
	einride_shipper_v1beta1_LaneService_GetLane         = &cobra.Command{
		Use: "GetLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.GetLane(cmd.Context(), &einride_shipper_v1beta1_LaneService_GetLane_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.BatchGetLanes.
var (
	einride_shipper_v1beta1_LaneService_BatchGetLanes_Request v1beta1.BatchGetLanesRequest
	einride_shipper_v1beta1_LaneService_BatchGetLanes         = &cobra.Command{
		Use: "BatchGetLanes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.BatchGetLanes(cmd.Context(), &einride_shipper_v1beta1_LaneService_BatchGetLanes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.SearchLanes.
var (
	einride_shipper_v1beta1_LaneService_SearchLanes_Request v1beta1.SearchLanesRequest
	einride_shipper_v1beta1_LaneService_SearchLanes         = &cobra.Command{
		Use: "SearchLanes",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.SearchLanes(cmd.Context(), &einride_shipper_v1beta1_LaneService_SearchLanes_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.UpdateLane.
var (
	einride_shipper_v1beta1_LaneService_UpdateLane_Request v1beta1.UpdateLaneRequest
	einride_shipper_v1beta1_LaneService_UpdateLane         = &cobra.Command{
		Use: "UpdateLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.UpdateLane(cmd.Context(), &einride_shipper_v1beta1_LaneService_UpdateLane_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.v1beta1.LaneService.ReferenceLane.
var (
	einride_shipper_v1beta1_LaneService_ReferenceLane_Request v1beta1.ReferenceLaneRequest
	einride_shipper_v1beta1_LaneService_ReferenceLane         = &cobra.Command{
		Use: "ReferenceLane",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_v1beta1_LaneServiceClient.ReferenceLane(cmd.Context(), &einride_shipper_v1beta1_LaneService_ReferenceLane_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddLaneServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_v1beta1_LaneService)
}

func init() {
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_CreateLane)

	einride_shipper_v1beta1_LaneService_CreateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_CreateLane_Request.Parent, "parent", "", "Resource name of the parent shipper where this lane will be created.")

	einride_shipper_v1beta1_LaneService_CreateLane_Request.Lane = new(v1beta1.Lane)
	einride_shipper_v1beta1_LaneService_CreateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_CreateLane_Request.Lane.Name, "lane.name", "", "The resource name of the lane.")
	einride_shipper_v1beta1_LaneService_CreateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_CreateLane_Request.Lane.IntegrationFile, "lane.integrationFile", "", "Resource name of the integration file the lane was ingested from, if the lane was ingested.")
	einride_shipper_v1beta1_LaneService_CreateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_CreateLane_Request.Lane.OriginSite, "lane.originSite", "", "Resource name of the site from which the lane originates.")
	einride_shipper_v1beta1_LaneService_CreateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_CreateLane_Request.Lane.DestinationSite, "lane.destinationSite", "", "Resource name of the site which the lane is destined for.")
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_GetLane)

	einride_shipper_v1beta1_LaneService_GetLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_GetLane_Request.Name, "name", "", "Resource name of the lane to retrieve.")
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_BatchGetLanes)

	einride_shipper_v1beta1_LaneService_BatchGetLanes.Flags().StringVar(&einride_shipper_v1beta1_LaneService_BatchGetLanes_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all lanes being retrieved.\nThe parent of all of the lanes specified in `names`\nmust match this field.")

	einride_shipper_v1beta1_LaneService_BatchGetLanes.Flags().StringSliceVar(&einride_shipper_v1beta1_LaneService_BatchGetLanes_Request.Names, "names", []string{}, "Resource names of the lanes to retrieve.\nA maximum of 1000 lanes can be retrieved in a batch.")
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_SearchLanes)

	einride_shipper_v1beta1_LaneService_SearchLanes.Flags().StringVar(&einride_shipper_v1beta1_LaneService_SearchLanes_Request.Parent, "parent", "", "Resource name of the parent shipper shared by all lanes being searched for.")

	// TODO: list: SiteFilters message

	einride_shipper_v1beta1_LaneService_SearchLanes.Flags().Int32Var(&einride_shipper_v1beta1_LaneService_SearchLanes_Request.PageSize, "pageSize", 10, "The maximum number of results to return. The service may return fewer\nresults than this value.\n\nIf unspecified, at most 50 results will be returned.\nThe maximum value is 1000; values above 1000 will be coerced to 1000.")

	einride_shipper_v1beta1_LaneService_SearchLanes.Flags().StringVar(&einride_shipper_v1beta1_LaneService_SearchLanes_Request.PageToken, "pageToken", "", "A page token, received from a previous call. Provide this to retrieve the\nsubsequent page.\n\nWhen paginating, all other parameters provided must match the call that\nprovided the page token.")
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_UpdateLane)

	einride_shipper_v1beta1_LaneService_UpdateLane_Request.Lane = new(v1beta1.Lane)
	einride_shipper_v1beta1_LaneService_UpdateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_UpdateLane_Request.Lane.Name, "lane.name", "", "The resource name of the lane.")
	einride_shipper_v1beta1_LaneService_UpdateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_UpdateLane_Request.Lane.IntegrationFile, "lane.integrationFile", "", "Resource name of the integration file the lane was ingested from, if the lane was ingested.")
	einride_shipper_v1beta1_LaneService_UpdateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_UpdateLane_Request.Lane.OriginSite, "lane.originSite", "", "Resource name of the site from which the lane originates.")
	einride_shipper_v1beta1_LaneService_UpdateLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_UpdateLane_Request.Lane.DestinationSite, "lane.destinationSite", "", "Resource name of the site which the lane is destined for.")

	einride_shipper_v1beta1_LaneService_UpdateLane_Request.UpdateMask = new(fieldmaskpb.FieldMask)
	einride_shipper_v1beta1_LaneService_UpdateLane.Flags().StringSliceVar(&einride_shipper_v1beta1_LaneService_UpdateLane_Request.UpdateMask.Paths, "updateMask.paths", []string{}, "The set of field mask paths.")
	einride_shipper_v1beta1_LaneService.AddCommand(einride_shipper_v1beta1_LaneService_ReferenceLane)

	einride_shipper_v1beta1_LaneService_ReferenceLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_ReferenceLane_Request.Parent, "parent", "", "Resource name of the shipper to look up the lane for.")

	einride_shipper_v1beta1_LaneService_ReferenceLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_ReferenceLane_Request.OriginSiteReferenceId, "originSiteReferenceId", "", "The reference ID of origin site for lane to retrieve.")

	einride_shipper_v1beta1_LaneService_ReferenceLane.Flags().StringVar(&einride_shipper_v1beta1_LaneService_ReferenceLane_Request.DestinationSiteReferenceId, "destinationSiteReferenceId", "", "The reference ID of desination site for lane to retrieve.")
}
