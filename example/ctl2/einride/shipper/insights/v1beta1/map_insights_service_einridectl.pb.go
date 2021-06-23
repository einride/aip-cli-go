package shipperinsightsv1beta1

import (
	fmt "fmt"
	ctl "github.com/einride/ctl"
	v1beta1 "github.com/einride/proto/gen/go/einride/shipper/insights/v1beta1"
	cobra "github.com/spf13/cobra"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// einride.shipper.insights.v1beta1.MapInsightsService.
var (
	einride_shipper_insights_v1beta1_MapInsightsServiceClient v1beta1.MapInsightsServiceClient
	einride_shipper_insights_v1beta1_MapInsightsService       = &cobra.Command{
		Use:   "einride.shipper.insights.v1beta1.MapInsightsService",
		Short: "Map insights service.",
		Long:  "Map insights service.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := ctl.ParseDialConfig(cmd.Flags())
			if err != nil {
				return err
			}
			conn, err := ctl.Dial(cmd.Context(), config)
			if err != nil {
				return err
			}
			einride_shipper_insights_v1beta1_MapInsightsServiceClient = v1beta1.NewMapInsightsServiceClient(conn)
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetShipperMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights_Request v1beta1.BatchGetShipperMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights         = &cobra.Command{
		Use: "BatchGetShipperMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_insights_v1beta1_MapInsightsServiceClient.BatchGetShipperMapInsights(cmd.Context(), &einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetSiteMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights_Request v1beta1.BatchGetSiteMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights         = &cobra.Command{
		Use: "BatchGetSiteMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_insights_v1beta1_MapInsightsServiceClient.BatchGetSiteMapInsights(cmd.Context(), &einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

// einride.shipper.insights.v1beta1.MapInsightsService.BatchGetLaneMapInsights.
var (
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights_Request v1beta1.BatchGetLaneMapInsightsRequest
	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights         = &cobra.Command{
		Use: "BatchGetLaneMapInsights",
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := einride_shipper_insights_v1beta1_MapInsightsServiceClient.BatchGetLaneMapInsights(cmd.Context(), &einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights_Request)
			if err != nil {
				return err
			}
			fmt.Println(protojson.Format(response))
			return nil
		},
	}
)

func AddMapInsightsServiceCommand(parent *cobra.Command) {
	parent.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService)
}

func init() {
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights)

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights.Flags().StringVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights_Request.Parent, "parent", "", "Resource name of the parent shipper.\nUse `\"shippers/-\"` or leave unset to batch get for multiple shippers.")

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights.Flags().StringSliceVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetShipperMapInsights_Request.Names, "names", []string{}, "Resource names of the shipper map insights to get.\nA maximum of 1000 insights can be retrieved in a batch.")
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights)

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights.Flags().StringVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights_Request.Parent, "parent", "", "Resource name of the shipper.")

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights.Flags().StringSliceVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetSiteMapInsights_Request.Names, "names", []string{}, "Resource names of the site map insights to get.\nA maximum of 1000 insights can be retrieved in a batch.")
	einride_shipper_insights_v1beta1_MapInsightsService.AddCommand(einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights)

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights.Flags().StringVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights_Request.Parent, "parent", "", "Resource name of the shipper.")

	einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights.Flags().StringSliceVar(&einride_shipper_insights_v1beta1_MapInsightsService_BatchGetLaneMapInsights_Request.Names, "names", []string{}, "Resource names of the lane map insights to get.\nA maximum of 1000 insights can be retrieved in a batch.")
}
